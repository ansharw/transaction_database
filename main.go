package main

import (
	"transaction_database/database"
	"transaction_database/template"
)

func main() {
	db := database.GetConnection()
	template.Menu(db)
}
