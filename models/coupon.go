package models

import "time"

type Coupon struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Brand     string    `json:"brand"`
	Value     float32   `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
	Expiry    time.Time `json:"expiry"`
}
