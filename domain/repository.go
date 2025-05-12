package domain

import "campaign-coupon-system/model"

type CampaignRepository interface {
	Save(campaign model.Campaign) (model.Campaign, error)
}

type CouponRepository interface {
	FindByCampaignId(campaignId int) ([]model.Coupon, error)
	Insert(coupons []model.Coupon) error
}
