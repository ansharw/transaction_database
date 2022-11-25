package model

import "time"

type TransactionDetails struct {
	id             int
	transaction_id int
	product_id     int
	product_name   string
	price          float64
	quantity       int `required:"true"`
	total          float64
	created_at     time.Time
}
