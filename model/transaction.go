package model

import "time"

type Transaction struct {
	id            int
	number        string
	customer_name string `required:"true"`
	email         string
	phone         string
	date          time.Time
	quantity      uint16 `required:"true"`
	discount      float64
	total         float64
	pay           float64
	created_at    time.Time
}
