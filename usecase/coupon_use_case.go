package usecase

import (
	"campaign-coupon-system/domain"
	"campaign-coupon-system/model"
)

type couponUseCase struct {
	repository domain.CouponRepository
}

func (c *couponUseCase) GetByCampaignId(campaignId int) ([]model.Coupon, error) {
	coupons, err := c.repository.FindByCampaignId(campaignId)
	return coupons, err
}

func (c *couponUseCase) CreateCoupons(campaignId int, amount int) error {
	//
	//campaign, err := c.repository.Insert(campaign)
	return nil
}

func NewCouponUseCase(repository domain.CouponRepository) domain.CouponUseCase {
	return &couponUseCase{repository: repository}
}
