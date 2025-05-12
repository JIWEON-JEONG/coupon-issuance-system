package model

import (
	"time"
)

// Coupon struct to describe Coupon object.
type Coupon struct {
	ID         uint      `db:"id" json:"id"`
	Code       string    `db:"code" json:"code" validate:"required,lte=10"`
	CampaignID uint      `db:"campaign_id" json:"campaign_id"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}

func (Coupon) TableName() string {
	return "coupon"
}
