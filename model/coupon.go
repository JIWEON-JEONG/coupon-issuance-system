package model

import (
	"time"
)

// Coupon struct to describe Coupon object.
type Coupon struct {
	ID            uint      `db:"id" json:"id"`
	Code          string    `db:"code" json:"code" validate:"required,lte=10"`
	CampaignID    int       `db:"campaign_id" json:"campaignId"`
	AvailableFrom time.Time `db:"available_from" json:"availableFrom"`
	CreatedAt     time.Time `db:"created_at" json:"createdAt"`
}

type IssuedCoupon struct {
	ID            uint      `db:"id" json:"id"`
	Code          string    `db:"code" json:"code" validate:"required,lte=10"`
	UserID        int       `db:"user_id" json:"userId"`
	CampaignID    int       `db:"campaign_id" json:"campaignId"`
	AvailableFrom time.Time `db:"available_from" json:"availableFrom"`
	CreatedAt     time.Time `db:"created_at" json:"createdAt"`
}

type CouponDto struct {
	CampaignID    int       `db:"campaign_id" json:"campaignId"`
	Code          string    `db:"code" json:"code" validate:"required,lte=10"`
	AvailableFrom time.Time `db:"available_from" json:"availableFrom"`
}

func (IssuedCoupon) TableName() string {
	return "issued_coupon"
}

func (Coupon) TableName() string {
	return "coupon"
}
