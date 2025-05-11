package repository

import (
	"campaign-coupon-system/domain"
	"campaign-coupon-system/model"
	"errors"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
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

func (c *couponRepository) Insert(coupons []model.Coupon) error {
	return nil
}

func NewCouponRepository(db *gorm.DB, cache *redis.Client) domain.CouponRepository {
	return &couponRepository{
		db:    db,
		cache: cache,
	}
}
