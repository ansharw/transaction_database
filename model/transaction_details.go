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

// Getter
func (transactionDetails *TransactionDetails) GetId() *int {
	return &transactionDetails.id
}

func (transactionDetails *TransactionDetails) GetTrxId() *int {
	return &transactionDetails.transaction_id
}

func (transactionDetails *TransactionDetails) GetProdId() *int {
	return &transactionDetails.product_id
}

func (transactionDetails *TransactionDetails) GetProdName() *string {
	return &transactionDetails.product_name
}

func (transactionDetails *TransactionDetails) GetPrice() *float64 {
	return &transactionDetails.price
}

func (transactionDetails *TransactionDetails) GetQty() *int {
	return &transactionDetails.quantity
}

func (transactionDetails *TransactionDetails) GetTotal() *float64 {
	return &transactionDetails.total
}

func (transactionDetails *TransactionDetails) GetCreatedAt() *time.Time {
	return &transactionDetails.created_at
}

// Setter
func (transactionDetails *TransactionDetails) SetId(id *int) {
	transactionDetails.id = *id
}

func (transactionDetails *TransactionDetails) SetTrxId(trxId *int) {
	transactionDetails.transaction_id = *trxId
}

func (transactionDetails *TransactionDetails) SetProdId(prodId *int) {
	transactionDetails.product_id = *prodId
}

func (transactionDetails *TransactionDetails) SetProdName(prodName *string) {
	transactionDetails.product_name = *prodName
}

func (transactionDetails *TransactionDetails) SetPrice(price *float64) {
	transactionDetails.price = *price
}

func (transactionDetails *TransactionDetails) SetQty(qty *int) {
	transactionDetails.quantity = *qty
}

func (transactionDetails *TransactionDetails) SetTotal(total *float64) {
	transactionDetails.total = *total
}

func (transactionDetails *TransactionDetails) SetCreatedAt(createdAt *time.Time) {
	transactionDetails.created_at = *createdAt
}
