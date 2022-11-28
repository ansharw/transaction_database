package template

import (
	"fmt"
	"transaction_database/helper"
)

func (template *transactionTemplate) ListProduct() {
	helper.ClearScreen()
	template.ShowProduct()
	helper.BackHandler()
	Menu(template.db)
}

func (template *transactionTemplate) ShowProduct() {
	fmt.Println("====================================================")
	fmt.Println("ID || Nama \t\t|| Price \t\t  ||")
	fmt.Println("====================================================")

	products, err := template.transactionHandler.GetProducts()
	if err != nil {
		panic(err)
	}

	if len(products) == 0 {
		fmt.Println("Data Kosong")
	} else {
		for _, v := range products {
			fmt.Printf("%v  || %v \t|| %v \t\t  ||\n", *v.GetId(), *v.GetName(), *v.GetPrice())
		}
	}
	fmt.Println("====================================================")
}
