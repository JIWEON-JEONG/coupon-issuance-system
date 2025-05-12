package usecase

import (
	"campaign-coupon-system/domain"
	"campaign-coupon-system/model"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type CampaignUseCase interface {
	CreateCampaign(ctx context.Context, campaign model.Campaign) (model.Campaign, error)
}

type campaignUseCase struct {
	db              *gorm.DB
	campaignService domain.CampaignService
	couponService   domain.CouponService
}

func (u *campaignUseCase) CreateCampaign(ctx context.Context, campaign model.Campaign) (model.Campaign, error) {
	var createdCampaign model.Campaign
	err := u.db.Transaction(func(tx *gorm.DB) error {
		var serviceErr error
		createdCampaign, serviceErr = u.campaignService.CreateCampaign(ctx, tx, campaign) // ctx와 tx를 전달
		if serviceErr != nil {
			return fmt.Errorf("failed to create campaign record via service: %w", serviceErr)
		}
		serviceErr = u.couponService.GenerateCoupons(ctx, tx, createdCampaign)

		if serviceErr != nil {
			return fmt.Errorf("failed to generate and insert coupons via service: %w", serviceErr)
		}
		return nil // 트랜잭션 블록 내 작업 성공 표시
	})
	if err != nil {
		return model.Campaign{}, fmt.Errorf("campaign creation transaction failed: %w", err)
	}
	return createdCampaign, nil
}

func NewCampaignUseCase(db *gorm.DB, campaignService domain.CampaignService, couponService domain.CouponService) CampaignUseCase {
	return &campaignUseCase{db: db, campaignService: campaignService, couponService: couponService}
}
