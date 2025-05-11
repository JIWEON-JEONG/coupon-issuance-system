package usecase

import (
	"campaign-coupon-system/domain"
	"campaign-coupon-system/model"
)

type campaignUseCase struct {
	repository domain.CampaignRepository
}

func (c *campaignUseCase) CreateCampaign(campaign model.Campaign) (model.Campaign, error) {
	campaign, err := c.repository.CreateCampaign(campaign)
	return campaign, err
}

func NewCampaignUseCase(repository domain.CampaignRepository) domain.CampaignUseCase {
	return &campaignUseCase{repository: repository}
}
