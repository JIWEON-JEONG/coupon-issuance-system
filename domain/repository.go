package domain

import (
	"campaign-coupon-system/model"
	"context"
	"gorm.io/gorm"
)

type CampaignRepository interface {
	Save(ctx context.Context, tx *gorm.DB, campaign model.Campaign) (model.Campaign, error)
}

type CouponRepository interface {
	FindByCampaignId(campaignId int) ([]model.Coupon, error)
	Insert(ctx context.Context, tx *gorm.DB, coupons []model.Coupon) error
}
