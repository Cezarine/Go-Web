package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase() *sql.DB {
	conection := "sistema:om315@/alura_loja" //"user=sistema dbname=alura_loja password=om315 host=localhost sslmode=disable"
	db, err := sql.Open("mysql", conection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
