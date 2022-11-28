package template

import (
	"fmt"
	"transaction_database/helper"
)

func (template *transactionTemplate) ListVoucher() {
	helper.ClearScreen()
	template.ShowVoucher()
	helper.BackHandler()
	Menu(template.db)
}

func (template *transactionTemplate) ShowVoucher() {
	fmt.Println("====================================================")
	fmt.Println("ID || Code \t\t|| Persen \t\t  ||")
	fmt.Println("====================================================")
	vouchers, err := template.transactionHandler.GetVouchers()
	if err != nil {
		panic(err)
	}
	if len(vouchers) == 0 {
		fmt.Println("Data Kosong")
	} else {
		for _, v := range vouchers {
			fmt.Printf("%v  || %v\t        || %v \t\t\t  ||\n", *v.GetId(), *v.GetCode(), *v.GetValue())
		}
	}
	fmt.Println("====================================================")
}
