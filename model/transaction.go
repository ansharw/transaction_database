package model

import "time"

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

func (transaction *Transaction) GetCreatedAt() *time.Time {
	return &transaction.created_at
}

func (transaction *Transaction) GetTransactionDetails() *[]TransactionDetails {
	return &transaction.transaction_details
}

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

func (transaction *Transaction) SetTransactionDetails(transactionDetail []TransactionDetails) {
	transaction.transaction_details = transactionDetail
}
