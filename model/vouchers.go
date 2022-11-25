package model

import "time"

type Vouchers struct {
	id         int
	name       string
	price      float64
	created_at time.Time
}
