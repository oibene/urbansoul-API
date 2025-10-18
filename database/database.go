package database

import (
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq" // postgres
)

func ConnectDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres",
		"user=urbanAdm password=urbansoul dbname=urbansouldb sslmode=disable")

	if err != nil {
		log.Println("Erro ao conectar banco!", err)
		return nil
	}
	return db
}
