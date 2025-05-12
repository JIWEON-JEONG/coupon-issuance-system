package domain

import (
	"campaign-coupon-system/model"
	"context"
	"gorm.io/gorm"
)

type CampaignRepository interface {
	Save(ctx context.Context, tx *gorm.DB, campaign model.Campaign) (model.Campaign, error)
	FindById(id int) (model.Campaign, error)
}

type CouponRepository interface {
	FindCouponDtoByCampaignIdOrNil(campaignId int) ([]model.CouponDto, error)
	FindIssuedCouponByCampaignId(campaignId int) ([]model.IssuedCoupon, error)
	InsertIssuedCoupon(coupon model.IssuedCoupon) error
	Insert(ctx context.Context, tx *gorm.DB, coupons []model.Coupon) error
}
