package model

import "time"

type Products struct {
	id         int
	name       string
	price      float64
	created_at time.Time
}

// Getter
func (product *Products) GetId() *int {
	return &product.id
}

func (product *Products) GetName() *string {
	return &product.name
}

func (product *Products) GetPrice() *float64 {
	return &product.price
}

// func (product *Products) GetCreatedAt() *time.Time {
// 	return &product.created_at
// }
