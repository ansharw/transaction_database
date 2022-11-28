package template

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strings"
	"transaction_database/helper"
	"transaction_database/model"
)

func (template *transactionTemplate) AddTransactionTemplate() {
	helper.ClearScreen()
	// var nameProduct string
	var idProduct, qtyProduct int
	fmt.Println("===================================")
	fmt.Println("=  Form Penjualan Produk Phincon  =")
	template.ShowProduct()

	// template.InputNameOfProduct(&nameProduct)
	template.InputIdOfProduct(&idProduct)
	// quantity of product
	template.InputQtyOfProduct(&qtyProduct)

	// Insert dulu transaction Detail
	// template.transactionHandler.

	// name, err := InputNameOfProduct()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	helper.BackHandler()
	// 	FormPenjualan()
	// }
	// qty, err := InputQtyOfProduct()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	helper.BackHandler()
	// 	FormPenjualan()
	// }
}

func (template *transactionTemplate) InputNameOfProduct(name *string) {
	fmt.Print("Masukkan Nama Produk : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nameOfProduct := scanner.Text()
	strLwr := strings.ToLower(nameOfProduct)
	products, err := template.transactionHandler.GetProducts()
	if err != nil {
		panic(err)
	}

	if !ValidateName(&strLwr) {
		fmt.Println("Nama Produk tidak boleh kosong")
		helper.BackHandler()
		template.InputNameOfProduct(&strLwr)
	} else if ValidateName(&strLwr) {
		for _, v := range products {
			if strings.ToLower(*v.GetName()) == strLwr {
				*name = strLwr
			}
		}
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
