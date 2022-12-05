package template

import (
	"bufio"
	"fmt"
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
	template.InputCustName(&custName)
	fmt.Print("Masukkan Email Anda : ")
	fmt.Scanln(&email)
	template.InputPhone(&phone)

loop:
	for {
		template.ShowProduct()
		product := template.InputIdOfProduct(&idProduct)
		template.InputQtyOfProduct(&qtyProduct)

		if *product.GetId() == idProduct {
			trxD := template.transactionHandler.GenerateProduct(idProduct, *product.GetName(), *product.GetPrice(), qtyProduct)
			*trx.GetTransactionDetails() = append(*trx.GetTransactionDetails(), trxD)
		}

		fmt.Print("Input produk kembali(y/n)? ")
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

		fmt.Println("ini cek :", trx)

		fmt.Println("")
		fmt.Println("Total Belanja: ", *trx.GetTotal())
		fmt.Println("Total Diskon: ", *trx.GetDiscount())
		fmt.Println("Total Bayar: ", *trx.GetTotal()-*trx.GetDiscount())
		fmt.Println("Jumlah uang anda:", *trx.GetPay())
		fmt.Println("ini cek 2 :", trx)
		tampilanStruk(&trx)
		helper.BackHandler()
		Menu(template.db)
	} else {
		fmt.Println("Transaksi tidak valid")
	}
}

func tampilanStruk(trx *model.Transaction) {
	helper.ClearScreen()
	fmt.Println("Struk Transaksi")
	fmt.Println("Qty\t Item\t Total")
	for _, v := range *trx.GetTransactionDetails() {
		fmt.Printf("%d\t %s\t %.0f\n", *v.GetQty(), *v.GetProdName(), *v.GetTotal())
	}
	fmt.Printf("Total\t \t %v\n", *trx.GetTotal())
	fmt.Printf("Discount\t \t %v\n", *trx.GetDiscount())
	var discounting = *trx.GetTotal() - *trx.GetDiscount()
	fmt.Printf("Total Discount\t \t %v\n", discounting)
	fmt.Printf("Payment\t \t %v\n", *trx.GetPay())
	fmt.Printf("Change\t \t %v\n", *trx.GetPay()-discounting)
	date := trx.GetDate().Format("2006-01-02")
	fmt.Printf("Transaction Date\t %v\n", date)
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

	*custName = customerName
}

func (template *transactionTemplate) InputIdOfProduct(idProduct *int) model.Products {
	fmt.Print("Masukkan Id Produk : ")
	fmt.Scanln(idProduct)

	product, err := template.transactionHandler.GetProduct(*idProduct)
	if err != nil {
		fmt.Println("id product tidak sesuai")
		helper.BackHandler()
		template.InputIdOfProduct(idProduct)
	}

	return product
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

func (template *transactionTemplate) InputPhone(phone *string) {
	fmt.Print("Masukkan Nomer Telepon : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	phoneNumber := scanner.Text()

	if !ValidatePhone(&phoneNumber) {
		fmt.Println("Nomer Telepon Kosong atau kurang dari 10 digit")
		helper.BackHandler()
		template.InputPhone(&phoneNumber)
	}

	*phone = phoneNumber
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

func ValidatePhone(phone *string) bool {
	if *phone == "" {
		return false
	} else if len(*phone) < 10 {
		return false
	}
	return true
}
