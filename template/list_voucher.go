package template

import (
	"fmt"
	"transaction_database/helper"
)

func (template *transactionTemplate) ListVoucher() {
	helper.ClearScreen()
	fmt.Println("====================================================")
	fmt.Println("ID || Code \t\t|| Persen \t\t  ||")
	fmt.Println("====================================================")
	products, err := template.transactionHandler.GetVouchers()
	if err != nil {
		panic(err)
	}
	if len(products) == 0 {
		fmt.Println("Data Kosong")
	} else {
		for _, v := range products {
			fmt.Printf("%v  || %v\t        || %v \t\t\t  ||\n", *v.GetId(), *v.GetCode(), *v.GetValue())
		}
	}
	fmt.Println("====================================================")
	helper.BackHandler()
	Menu(template.db)
}
