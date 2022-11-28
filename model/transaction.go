package model

import "time"

// backup
// type Transaction struct {
// 	id                  int
// 	number              string
// 	customer_name       string `required:"true"`
// 	email               string
// 	phone               string
// 	date                time.Time
// 	quantity            int `required:"true"`
// 	discount            float64
// 	total               float64
// 	pay                 float64
// 	created_at          time.Time
// 	transaction_details []TransactionDetails
// }

// eksperimen
type Transaction struct {
	id                  int
	number              string
	customer_name       string `required:"true"`
	email               string
	phone               string
	date                time.Time // ini nanti time.Now
	quantity            int       `required:"true"` // ini nanti di isi
	discount            float64
	total               float64
	pay                 float64
	created_at          time.Time
	transaction_details []TransactionDetails
}

// type TransactionDetail struct {
// 	id         int
// 	price      float64
// 	quantity   int `required:"true"`
// 	total      float64
// 	created_at time.Time
// }

// Getter
func (transaction *Transaction) GetId() *int {
	return &transaction.id
}

func (transaction *Transaction) GetNumber() *string {
	return &transaction.number
}

func (transaction *Transaction) GetCustomerName() *string {
	return &transaction.customer_name
}

func (transaction *Transaction) GetEmail() *string {
	return &transaction.email
}

func (transaction *Transaction) GetPhone() *string {
	return &transaction.number
}

func (transaction *Transaction) GetDate() *time.Time {
	return &transaction.date
}

func (transaction *Transaction) GetQty() *int {
	return &transaction.quantity
}

func (transaction *Transaction) GetDiscount() *float64 {
	return &transaction.discount
}

func (transaction *Transaction) GetTotal() *float64 {
	return &transaction.total
}

func (transaction *Transaction) GetPay() *float64 {
	return &transaction.pay
}

// func (transaction *Transaction) GetCreatedAt() *time.Time {
// 	return &transaction.created_at
// }

func (transaction *Transaction) GetTransactionDetails() []TransactionDetails {
	return transaction.transaction_details
}

// func (transactionDetail *TransactionDetail) GetTransactionDetail() (*int, *float64, *float64) {
// 	return &transactionDetail.id, &transactionDetail.price, &transactionDetail.total
// }

// Setter
func (transaction *Transaction) SetId(id *int) {
	transaction.id = *id
}

func (transaction *Transaction) SetNumber(number *string) {
	transaction.number = *number
}

func (transaction *Transaction) SetCustomerName(custName *string) {
	transaction.customer_name = *custName
}

func (transaction *Transaction) SetEmail(email *string) {
	transaction.email = *email
}

func (transaction *Transaction) SetPhone(phone *string) {
	transaction.phone = *phone
}

func (transaction *Transaction) SetDate(date *time.Time) {
	transaction.date = *date
}

func (transaction *Transaction) SetQty(qty *int) {
	transaction.quantity = *qty
}

func (transaction *Transaction) SetDiscount(discount *float64) {
	transaction.discount = *discount
}

func (transaction *Transaction) SetTotal(total *float64) {
	transaction.total = *total
}

func (transaction *Transaction) SetPay(pay *float64) {
	transaction.pay = *pay
}

// func (transaction *Transaction) SetCreatedAt(createdAt *time.Time) {
// 	transaction.created_at = *createdAt
// }

func (transaction *Transaction) SetTransactionDetails(transactionDetail []TransactionDetails) {
	transaction.transaction_details = transactionDetail
}

// func (transactionDetail *TransactionDetail) SetTrx(id *int, price *float64, quantity *int, total *float64) {
// 	transactionDetail.id = *id
// 	transactionDetail.price = *price
// 	transactionDetail.quantity = *quantity
// 	transactionDetail.total = *total
// }
