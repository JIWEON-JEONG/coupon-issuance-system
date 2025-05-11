package domain

import "campaign-coupon-system/model"

type CouponRepository interface {
	FindByCampaignId(campaignId int) ([]model.Coupon, error)
	Insert(coupons []model.Coupon) error
}

type CouponUseCase interface {
	GetByCampaignId(campaignId int) ([]model.Coupon, error)
	CreateCoupons(campaignId int, amount int) error
}
