package endpoints

import (
	"database/sql"
)

type Handlers struct {
	DB           *sql.DB
	Pepper       string
	links        []ChangePassData
	transactions []TransactionData
}

type ChangePassData struct {
	link string
	id   int
}

type TransactionData struct {
	Code   string
	From   string
	To     string
	Amount int
	Title  string
}
