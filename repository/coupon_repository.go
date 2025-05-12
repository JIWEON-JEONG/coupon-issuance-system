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

func (c *couponRepository) FindByCampaignId(campaignId int) ([]model.Coupon, error) {
	var coupons []model.Coupon
	if err := c.db.Where("campaign_id = ?", campaignId).Find(&coupons).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Coupon{}, nil
		}
		return nil, errors.New("internal server error: failed to find coupons by campaign ID")
	}
	return coupons, nil
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
