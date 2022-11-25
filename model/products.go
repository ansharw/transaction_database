package model

import "time"

type Products struct {
	id         int
	name       string
	price      float64
	created_at time.Time
}
