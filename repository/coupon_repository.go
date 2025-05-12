package repository

import (
	"campaign-coupon-system/domain"
	"campaign-coupon-system/model"
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"time"
)

type couponRepository struct {
	db    *gorm.DB
	cache *redis.Client
}

func (c *couponRepository) FindCouponDtoByCampaignIdOrNil(campaignId int) ([]model.CouponDto, error) {
	var couponDtos []model.CouponDto
	ctx := context.Background()

	redisSetKey := fmt.Sprintf("campaign:%d:available_codes", campaignId)
	hashKey := fmt.Sprintf("campaign:%d:coupon_data", campaignId)

	sPopResult := c.cache.SPop(ctx, redisSetKey)
	couponCode, err := sPopResult.Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to pop coupon code: %w", err)
	}

	hGetResult := c.cache.HGet(ctx, hashKey, couponCode)
	availableFromStr, err := hGetResult.Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, fmt.Errorf("available_from not found for coupon: %s", couponCode)
		}
		return nil, fmt.Errorf("failed to get available_from: %w", err)
	}

	availableFrom, err := time.Parse(time.RFC3339, availableFromStr)

	couponDto := model.CouponDto{
		CampaignID:    campaignId,
		Code:          couponCode,
		AvailableFrom: availableFrom,
	}
	couponDtos = append(couponDtos, couponDto)

	return couponDtos, nil
}

func (c *couponRepository) InsertIssuedCoupon(coupon model.IssuedCoupon) error {
	result := c.db.Create(&coupon)
	if result.Error != nil {
		return fmt.Errorf("failed to insert issued coupon: %w", result.Error)
	}
	return nil
}

func (c *couponRepository) Insert(ctx context.Context, tx *gorm.DB, coupons []model.Coupon) error {
	if len(coupons) == 0 {
		return nil
	}

	dbBatchSize := 10000
	if err := tx.WithContext(ctx).CreateInBatches(coupons, dbBatchSize).Error; err != nil {
		return fmt.Errorf("failed to insert coupons into database: %w", err)
	}

	pipe := c.cache.Pipeline()
	setKeys := make(map[string]struct{})

	for _, coupon := range coupons {
		redisSetKey := fmt.Sprintf("campaign:%d:available_codes", coupon.CampaignID)
		pipe.SAdd(ctx, redisSetKey, coupon.Code)

		hashKey := fmt.Sprintf("campaign:%d:coupon_data", coupon.CampaignID)
		pipe.HSet(ctx, hashKey, coupon.Code, coupon.AvailableFrom)

		availableFrom := coupon.AvailableFrom
		deletedAt := availableFrom.Add(24 * time.Hour)
		ttl := time.Until(deletedAt.UTC())

		if _, exists := setKeys[redisSetKey]; !exists {
			pipe.Expire(ctx, redisSetKey, ttl)
			setKeys[redisSetKey] = struct{}{}
		}
	}

	if _, err := pipe.Exec(ctx); err != nil {
		return fmt.Errorf("failed to insert available coupon codes into redis cache: %w", err)
	}
	return nil
}

func NewCouponRepository(db *gorm.DB, cache *redis.Client) domain.CouponRepository {
	return &couponRepository{
		db:    db,
		cache: cache,
	}
}
