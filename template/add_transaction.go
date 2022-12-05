package template

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"transaction_database/helper"
	"transaction_database/model"
)

func (template *transactionTemplate) AddTransactionTemplate() {
	helper.ClearScreen()
	// var nameProduct string
	var idProduct, qtyProduct int
	var custName, email, phone string
	var discount string
	var pay float64
	var trx model.Transaction
	fmt.Println("===================================")
	fmt.Println("=  Form Penjualan Produk Phincon  =")
	fmt.Print("Masukkan Nama Customer : ")
	fmt.Scanln(&custName)
	fmt.Print("Masukkan Email Anda : ")
	fmt.Scanln(&email)
	fmt.Print("Masukkan Phone Anda : ")
	fmt.Scanln(&phone)

loop:
	for {
		template.ShowProduct()
		// template.InputNameOfProduct(&nameProduct)
		template.InputIdOfProduct(&idProduct)
		// quantity of product
		template.InputQtyOfProduct(&qtyProduct)

		product, err := template.transactionHandler.GetProduct(idProduct)
		if err != nil {
			panic(err)
		}

		if *product.GetId() == idProduct {
			trxD := template.transactionHandler.GenerateProduct(idProduct, *product.GetName(), *product.GetPrice(), qtyProduct)
			*trx.GetTransactionDetails() = append(*trx.GetTransactionDetails(), trxD)
		}

		fmt.Println("Input produk kembali? (y/n)")
		var option string
		fmt.Scanln(&option)
		switch option {
		case "n":
			break loop
		case "y":
			continue
		default:
			break loop
		}
	}

	if len(*trx.GetTransactionDetails()) > 0 {
		template.ShowVoucher()
		fmt.Print("Masukkan Code Voucher : ")
		fmt.Scanln(&discount)
		fmt.Print("Masukkan Uang Anda : ")
		fmt.Scanln(&pay)

		trx, _, err := template.transactionHandler.AddTransaction(&trx, custName, email, phone, discount, pay)
		if err != nil {
			panic(err)
		}

		fmt.Println("")
		fmt.Println("Total Belanja: ", *trx.GetTotal())
		fmt.Println("Total Diskon: ", *trx.GetDiscount())
		fmt.Println("Total Bayar: ", *trx.GetTotal()-*trx.GetDiscount())
		fmt.Println("jumlah uang anda:", *trx.GetPay())
		helper.BackHandler()
		Menu(template.db)
	}
}

func (template *transactionTemplate) InputCustName(custName *string) {
	fmt.Print("Masukkan Nama Customer : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	customerName := scanner.Text()

	if !ValidateCustName(&customerName) {
		fmt.Println("Nama Customer tidak boleh kosong")
		helper.BackHandler()
		template.InputCustName(&customerName)
	}
}

func (template *transactionTemplate) InputIdOfProduct(prodId *int) {
	var input int
	fmt.Print("Masukkan Id Produk : ")
	fmt.Scanln(&input)
	products, err := template.transactionHandler.GetProducts()
	if err != nil {
		panic(err)
	}

	if !ValidateIdProduct(&input) {
		fmt.Println("Id Produk tidak boleh kosong")
		helper.BackHandler()
		template.InputIdOfProduct(&input)
	} else if ValidateIdProduct(&input) {
		for _, v := range products {
			if *v.GetId() == input {
				*prodId = input
			}
		}
	}
}

func (template *transactionTemplate) InputQtyOfProduct(qty *int) {
	var input int
	fmt.Print("Masukkan Jumlah Produk : ")
	fmt.Scanln(&input)

	if !ValidateQty(&input) {
		fmt.Println("Jumlah Produk tidak boleh kosong")
		helper.BackHandler()
		template.InputQtyOfProduct(&input)
	}
	*qty = input
}

func ValidateName(name *string) bool {
	var c model.TransactionDetails
	typeOf := reflect.TypeOf(c)
	if typeOf.Field(3).Tag.Get("required") == "true" {
		if *name == "" {
			return false
		}
	}
	return true
}

func ValidateQty(qty *int) bool {
	var c model.TransactionDetails
	typeOf := reflect.TypeOf(c)
	if typeOf.Field(5).Tag.Get("type") == "number" {
		if *qty == 0 {
			return false
		} else if float64(*qty) == math.NaN() {
			return false
		}
	}
	return true
}

func ValidateIdProduct(prodId *int) bool {
	var c model.TransactionDetails
	typeOf := reflect.TypeOf(c)
	if typeOf.Field(2).Tag.Get("required") == "true" {
		if *prodId == 0 {
			return false
		} else if float64(*prodId) == math.NaN() {
			return false
		}
	}
	return true
}

func ValidateCustName(custName *string) bool {
	var c model.Transaction
	typeOf := reflect.TypeOf(c)
	if typeOf.Field(2).Tag.Get("required") == "true" {
		if *custName == "" {
			return false
		}
	}
	return true
}
