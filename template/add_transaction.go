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
	var tempSliceMap []map[string]interface{}
	fmt.Println("===================================")
	fmt.Println("=  Form Penjualan Produk Phincon  =")
	fmt.Print("Masukkan Nama Customer : ")
	fmt.Scanln(&custName)
	fmt.Print("Masukkan Email Anda : ")
	fmt.Scanln(&email)
	fmt.Print("Masukkan Phone Anda : ")
	fmt.Scanln(&phone)

outer:
	for {
		template.ShowProduct()
		// template.InputNameOfProduct(&nameProduct)
		template.InputIdOfProduct(&idProduct)
		// quantity of product
		template.InputQtyOfProduct(&qtyProduct)

		tempMap := map[string]interface{}{
			"idProduct":  idProduct,
			"qtyProduct": qtyProduct,
		}
		tempSliceMap = append(tempSliceMap, tempMap)
		fmt.Println("Input data kembali? (y/n)")
		var option string
		fmt.Scanln(&option)
		switch option {
		case "n":
			break outer
		case "y":
			continue
		default:
			break outer
		}
	}
	fmt.Println(tempSliceMap)

	template.ShowVoucher()
	fmt.Print("Masukkan Code Voucher : ")
	fmt.Scanln(&discount)
	fmt.Print("Masukkan Uang Anda : ")
	fmt.Scanln(&pay)

	// fmt.Println(idProduct, qtyProduct, custName, email, phone, discount, pay)

	// _, _, err := template.transactionHandler.AddTransaction(idProduct, qtyProduct, custName, email, phone, discount, pay)

	_, _, err := template.transactionHandler.AddTransaction(tempSliceMap, custName, email, phone, discount, pay)
	// fmt.Println(trx)
	// fmt.Println(trxD)
	// fmt.Println(err)
	if err != nil {
		panic(err)
	}

	fmt.Println("")
	fmt.Println("Data berhasil di input.")
	helper.BackHandler()
	Menu(template.db)
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
