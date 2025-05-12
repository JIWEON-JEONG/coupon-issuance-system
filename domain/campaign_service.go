package domain

import (
	"campaign-coupon-system/model"
	"context"
	"gorm.io/gorm"
)

type CampaignService interface {
	CreateCampaign(ctx context.Context, tx *gorm.DB, campaign model.Campaign) (model.Campaign, error)
}

type campaignService struct {
	repository CampaignRepository
}

func (c *campaignService) CreateCampaign(ctx context.Context, tx *gorm.DB, campaign model.Campaign) (model.Campaign, error) {
	campaign, err := c.repository.Save(ctx, tx, campaign)
	return campaign, err
}

func NewCampaignService(repository CampaignRepository) CampaignService {
	return &campaignService{repository: repository}
}
