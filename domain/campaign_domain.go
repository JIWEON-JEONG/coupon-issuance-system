package domain

import "campaign-coupon-system/model"

type CampaignRepository interface {
	CreateCampaign(campaign model.Campaign) (model.Campaign, error)
}

type CampaignUseCase interface {
	CreateCampaign(campaign model.Campaign) (model.Campaign, error)
}
