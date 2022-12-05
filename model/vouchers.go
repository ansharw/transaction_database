package model

import "time"

type Vouchers struct {
	id         int
	code       string
	value      float64
	created_at time.Time
}

// Getter
func (vouchers *Vouchers) GetId() *int {
	return &vouchers.id
}

func (vouchers *Vouchers) GetCode() *string {
	return &vouchers.code
}

func (vouchers *Vouchers) GetValue() *float64 {
	return &vouchers.value
}

func (vouchers *Vouchers) GetCreatedAt() *time.Time {
	return &vouchers.created_at
}
