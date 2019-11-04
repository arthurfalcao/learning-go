package db

import (
	"database/sql"

	_ "github.com/lib/pq" // postgres lib
)

func ConnectWithDB() *sql.DB {
	connect := "user=postgres dbname=alura_loja password=123456 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err.Error())
	}

	return db
}
