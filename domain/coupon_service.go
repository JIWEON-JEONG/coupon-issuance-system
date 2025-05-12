package domain

import "campaign-coupon-system/model"

type CouponService interface {
	GetByCampaignId(campaignId int) ([]model.Coupon, error)
}

type couponService struct {
	repository CouponRepository
}

func (c *couponService) GetByCampaignId(campaignId int) ([]model.Coupon, error) {
	coupons, err := c.repository.FindByCampaignId(campaignId)
	return coupons, err
}

func NewCouponService(repository CouponRepository) CouponService {
	return &couponService{repository: repository}
}
