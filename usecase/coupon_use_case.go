package usecase

import (
	"campaign-coupon-system/domain"
	"campaign-coupon-system/model"
)

type CouponUseCase interface {
	GetByCampaignId(campaignId int) ([]model.Coupon, error)
}

type couponUseCase struct {
	service domain.CouponService
}

func (c *couponUseCase) GetByCampaignId(campaignId int) ([]model.Coupon, error) {
	coupons, err := c.service.GetByCampaignId(campaignId)
	return coupons, err
}

func (c *couponUseCase) CreateCoupons(campaignId int, amount int) error {
	//
	//campaign, err := c.repository.Insert(campaign)
	return nil
}

func NewCouponUseCase(service domain.CouponService) CouponUseCase {
	return &couponUseCase{service: service}
}
