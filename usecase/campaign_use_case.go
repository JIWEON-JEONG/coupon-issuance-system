package usecase

import (
	"campaign-coupon-system/domain"
	"campaign-coupon-system/model"
)

type CampaignUseCase interface {
	CreateCampaign(campaign model.Campaign) (model.Campaign, error)
}

type campaignUseCase struct {
	service domain.CampaignService
}

func (c *campaignUseCase) CreateCampaign(campaign model.Campaign) (model.Campaign, error) {
	campaign, err := c.service.CreateCampaign(campaign)
	return campaign, err
}

func NewCampaignUseCase(service domain.CampaignService) CampaignUseCase {
	return &campaignUseCase{service: service}
}
