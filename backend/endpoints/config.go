package endpoints

import "database/sql"

type Handlers struct {
	DB     *sql.DB
	Pepper string
	links  []ChangePassData
}

type ChangePassData struct {
	link string
	id   int
}
