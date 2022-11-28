package controller

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"transaction_database/model"
	"transaction_database/repository"
)

// kumpulan fungsi yang nanti dipakai di menu / buat end user
type TransactionHandler interface {
	GetProducts() ([]model.Products, error)
	GetVouchers() ([]model.Vouchers, error)

	GetTransactions() ([]model.Transaction, error)
	// AddTransaction(custName, email, phone string, total, pay float64, trxDetails []model.TransactionDetails) error
	AddTransaction(prodsId int, quantity int, custName, custEmail, custPhone, discount string, pay float64) error
}

// nyimpen
type transactionHandler struct {
	db                           *sql.DB
	productsRepository           repository.ProductsRepository
	transactionRepository        repository.TransactionRepository
	transactionDetailsRepository repository.TransactionDetailsRepository
	vouchersRepository           repository.VouchersRepository
}

func NewTransactionHandler(db *sql.DB, productsRepository repository.ProductsRepository, transactionRepository repository.TransactionRepository, transactionDetailsRepository repository.TransactionDetailsRepository, vouchersRepository repository.VouchersRepository) *transactionHandler {
	return &transactionHandler{db, productsRepository, transactionRepository, transactionDetailsRepository, vouchersRepository}
}

func (handler *transactionHandler) GetProducts() ([]model.Products, error) {
	ctx := context.Background()

	products, err := handler.productsRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (handler *transactionHandler) GetVouchers() ([]model.Vouchers, error) {
	ctx := context.Background()

	vouchers, err := handler.vouchersRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return vouchers, nil
}

// eksperimen
func (handler *transactionHandler) GetTransactions() ([]model.Transaction, error) {
	ctx := context.Background()

	transactions, err := handler.transactionRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	for i, v := range transactions {
		transactionDetails, err := handler.transactionDetailsRepository.GetTrxDetailsByTrxId(ctx, *v.GetId())
		if err != nil {
			fmt.Println("ini dari transaction controller ", v)
			fmt.Println(v.GetId())
			return nil, err
		}
		transactions[i].SetTransactionDetails(transactionDetails)
	}

	return transactions, nil
}

// eksperimen
// ini untuk manggil keseluruhan yaitu transaction detail abis itu transaction
func (handler *transactionHandler) AddTransaction(prodsId int, quantity int, custName, custEmail, custPhone, discount string, pay float64) error {
	ctx := context.Background()

	var transaction model.Transaction
	var transactionDetail model.TransactionDetails
	var transactionDetails []model.TransactionDetails

	transactionDetail.SetProdId(&prodsId)
	transactionDetail.SetQty(&quantity)

	products, err := handler.GetProducts()
	if err != nil {
		panic(err)
	}

	for _, v := range products {
		if v.GetId() == &prodsId {
			transactionDetail.SetProdName(v.GetName())
			transactionDetail.SetPrice(v.GetPrice())
		}
	}

	var price float64 = *transactionDetail.GetPrice()
	var qty float64 = float64(quantity)
	total := qty * price
	transactionDetail.SetTotal(&total)

	trxDetail, err := handler.transactionDetailsRepository.AddTrxDetails(ctx, transactionDetail)
	if err != nil {
		return err
	}
	transactionDetails = append(transactionDetails, trxDetail)

	tx, err := handler.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	timeNow := time.Now()
	transaction.SetDate(&timeNow)

	trxNumber := GenerateTrxNumber()
	transaction.SetNumber(trxNumber)
	transaction.SetEmail(&custEmail)
	transaction.SetPhone(&custPhone)

	vouchers, err := handler.GetVouchers()
	if err != nil {
		panic(err)
	}

	for _, v := range vouchers {
		if *v.GetCode() == discount {
			total := transactionDetail.GetTotal()
			disc := *v.GetValue() / float64(100)
			discounting := *total * disc
			totalFinal := *total - discounting
			transaction.SetDiscount(&totalFinal)
		}
	}
	transaction.SetPay(&pay)

	trxFinal, err := handler.transactionRepository.AddTrx(ctx, tx, transactionDetails, transaction)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	transaction.SetTransactionDetails(trxFinal)

	return nil
}

func GenerateTrxNumber() *string {
	tempTrxRand := rand.Int()
	trxNumber := "TRXPHC-" + strconv.Itoa(tempTrxRand)
	return &trxNumber
}



// eksperimen
// func (handler *transactionHandler) AddTransactionefef(custName, email, phone string, total, pay float64, trxDetails []model.TransactionDetails) error {
// 	// number, date, created_at
// 	ctx := context.Background()

// 	var transaction model.Transaction
// 	var transactionDetails []model.TransactionDetails

// 	transaction.SetCustomerName(&custName)
// 	transaction.SetEmail(&email)
// 	transaction.SetPhone(&phone)
// 	transaction.SetTotal(&total)
// 	transaction.SetPay(&pay)

// 	transaction, err := handler.transactionRepository.AddTrx(ctx, transaction)
// 	if err != nil {
// 		return err
// 	}

// 	tx, err := handler.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	transactionDetails, err = handler.transactionDetailsRepository.AddTrxDetails(ctx, tx, transactionDetails, *transaction.GetId())

// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	tx.Commit()
// 	transaction.SetTransactionDetails(transactionDetails)

// 	return nil
// }
