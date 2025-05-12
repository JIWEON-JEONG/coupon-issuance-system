package repository

import (
	"campaign-coupon-system/domain"
	"campaign-coupon-system/model"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type campaignRepository struct {
	db    *gorm.DB
	cache *redis.Client
}

func (c *campaignRepository) Save(ctx context.Context, tx *gorm.DB, campaign model.Campaign) (model.Campaign, error) {
	if err := tx.WithContext(ctx).Create(&campaign).Error; err != nil {
		return model.Campaign{}, errors.New("internal server error: cannot create campaign") // 에러와 빈 캠페인 객체 반환
	}
	return campaign, nil
}

func NewCampaignRepository(db *gorm.DB, cache *redis.Client) domain.CampaignRepository {
	return &campaignRepository{
		db:    db,
		cache: cache,
	}
}
