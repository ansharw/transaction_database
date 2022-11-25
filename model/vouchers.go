package model

import "time"

type Vouchers struct {
	id         int
	name       string
	price      float64
	created_at time.Time
}

// Getter
func (vouchers *Vouchers) GetId() *int {
	return &vouchers.id
}

func (vouchers *Vouchers) GetName() *string {
	return &vouchers.name
}

func (vouchers *Vouchers) GetPrice() *float64 {
	return &vouchers.price
}

func (vouchers *Vouchers) GetCreatedAt() *time.Time {
	return &vouchers.created_at
}
