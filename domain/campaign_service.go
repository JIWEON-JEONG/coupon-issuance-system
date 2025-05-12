package domain

import "campaign-coupon-system/model"

type CampaignService interface {
	CreateCampaign(campaign model.Campaign) (model.Campaign, error)
}

type campaignService struct {
	repository CampaignRepository
}

func (c *campaignService) CreateCampaign(campaign model.Campaign) (model.Campaign, error) {
	campaign, err := c.repository.Save(campaign)
	return campaign, err
}

func NewCampaignService(repository CampaignRepository) CampaignService {
	return &campaignService{repository: repository}
}
