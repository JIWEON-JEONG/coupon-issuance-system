package model

import (
	"time"
)

// Campaign struct to describe Campaign object.
type Campaign struct {
	ID               uint      `db:"id" json:"id"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
	AvailableCoupons int       `db:"available_coupons" json:"available_coupons" validate:"required,min=0"`
	StartDateTime    time.Time `db:"start_date_time" json:"start_date_time" validate:"required"`
}

func (Campaign) TableName() string {
	return "campaign"
}
