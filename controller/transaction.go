package controller

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"
	"transaction_database/model"
	"transaction_database/repository"
)

type TransactionHandler interface {
	GetProducts() ([]model.Products, error)
	GetProduct(id int) (model.Products, error)
	GetVouchers() ([]model.Vouchers, error)
	GetVoucher(code string) (model.Vouchers, error)
	GetTransactions() ([]model.Transaction, error)
	GetTransactionByNumber(trxNumber string) (model.Transaction, error)
	AddTransaction(trx *model.Transaction, custName, custEmail, custPhone string, discount string, pay float64) (model.Transaction, []model.TransactionDetails, error)
	GenerateProduct(idProduct int, prodName string, prodPrice float64, quantity int) model.TransactionDetails
}

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

func (handler *transactionHandler) GetTransactionByNumber(trxNumber string) (model.Transaction, error) {
	ctx := context.Background()

	trx, err := handler.transactionRepository.FindByNumber(ctx, trxNumber)
	if err != nil {
		return trx, err
	}

	if err == sql.ErrNoRows {
		return trx, err
	}

	trxD, err := handler.transactionDetailsRepository.GetTrxDetailsByTrxId(ctx, *trx.GetId())
	if err != nil {
		panic(trxD)
	}
	trx.SetTransactionDetails(trxD)

	return trx, nil
}

func (handler *transactionHandler) GetProducts() ([]model.Products, error) {
	ctx := context.Background()

	products, err := handler.productsRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (handler *transactionHandler) GetProduct(id int) (model.Products, error) {
	ctx := context.Background()
	var prod model.Products

	prod, err := handler.productsRepository.FindProduct(ctx, id)
	if err == sql.ErrNoRows {
		return prod, err
	}

	return prod, nil
}

func (handler *transactionHandler) GetVouchers() ([]model.Vouchers, error) {
	ctx := context.Background()

	vouchers, err := handler.vouchersRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return vouchers, nil
}

func (handler *transactionHandler) GetVoucher(code string) (model.Vouchers, error) {
	ctx := context.Background()
	var vouch model.Vouchers

	vouch, err := handler.vouchersRepository.FindVoucher(ctx, code)
	if err == sql.ErrNoRows {
		return vouch, err
	}

	return vouch, nil
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
			return nil, err
		}
		transactions[i].SetTransactionDetails(transactionDetails)
	}

	return transactions, nil
}

func (handler *transactionHandler) GenerateProduct(idProduct int, prodName string, prodPrice float64, quantity int) model.TransactionDetails {
	var trxD model.TransactionDetails
	trxD.SetProdId(&idProduct)
	trxD.SetProdName(&prodName)
	trxD.SetQty(&quantity)
	trxD.SetPrice(&prodPrice)
	var price float64 = *trxD.GetPrice()
	var qty float64 = float64(*trxD.GetQty())
	total := qty * price
	trxD.SetTotal(&total)
	fmt.Println(trxD)
	return trxD
}

func GenerateTrxNumber() string {
	var trxNumber string = strconv.Itoa(time.Now().Nanosecond())
	return trxNumber[:5]
}

func (handler *transactionHandler) AddTransaction(trx *model.Transaction, custName, custEmail, custPhone string, discount string, pay float64) (model.Transaction, []model.TransactionDetails, error) {
	ctx := context.Background()

	var transaction model.Transaction
	var transactionDetails []model.TransactionDetails

	var totalQuantity int
	var totalFinal float64
	for _, v := range *trx.GetTransactionDetails() {
		var qty int = *v.GetQty()
		totalQuantity += qty
		var total float64 = *v.GetTotal()
		totalFinal += total
		transaction.SetQty(&totalQuantity)
		transaction.SetTotal(&totalFinal)
	}

	var layoutFormat, value string
	var date time.Time
	layoutFormat = "2006-01-02"
	value = time.Now().Format("2006-01-02")
	date, _ = time.Parse(layoutFormat, value)
	transaction.SetDate(&date)

	trxNumber := GenerateTrxNumber()
	transaction.SetNumber(&trxNumber)

	transaction.SetPay(&pay)
	transaction.SetEmail(&custEmail)
	transaction.SetPhone(&custPhone)
	transaction.SetCustomerName(&custName)

	voucher, err := handler.GetVoucher(discount)
	if err != nil {
		panic(err)
	}

	var nol float64 = 0.0
	if totalFinal > 300000 {
		if *voucher.GetCode() == discount {
			total := transaction.GetTotal()
			disc := *voucher.GetValue() / float64(100)
			discounting := *total * disc
			transaction.SetDiscount(&discounting)
		} else {
			transaction.SetDiscount(&nol)
		}
	} else {
		transaction.SetDiscount(&nol)
	}

	trxs, err := handler.transactionRepository.AddTrx(ctx, transaction)
	if err != nil {
		return transaction, transactionDetails, err
	}

	for _, v := range *trx.GetTransactionDetails() {
		_, err := handler.transactionDetailsRepository.AddTrxDetails(ctx, v, *trxs.GetId())
		if err != nil {
			return transaction, transactionDetails, err
		}
		transactionDetails = append(transactionDetails, v)
		transaction.SetTransactionDetails(transactionDetails)
	}
	return transaction, transactionDetails, nil
}
