package controller

import "time"

type CreateCampaignDto struct {
	AvailableCoupons int       `json:"availableCoupons" validate:"required"`
	StartDateTime    time.Time `json:"startDateTime" validate:"required"`
}
