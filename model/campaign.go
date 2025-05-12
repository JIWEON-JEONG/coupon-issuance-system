package model

import (
	"time"
)

// Campaign struct to describe Campaign object.
type Campaign struct {
	ID               uint      `db:"id" json:"id"`
	CreatedAt        time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt        time.Time `db:"updated_at" json:"updatedAt"`
	AvailableCoupons int       `db:"available_coupons" json:"availableCoupons" validate:"required,min=0"`
	StartDateTime    time.Time `db:"start_date_time" json:"startDateTime" validate:"required"`
}

type CampaignInfoDto struct {
	CampaignID    uint      `json:"campaignId"`
	IssuedCodes   []string  `json:"issuedCodes"`
	StartDateTime time.Time `json:"startDateTime"`
}

func (Campaign) TableName() string {
	return "campaign"
}
