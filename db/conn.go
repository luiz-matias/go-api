package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "go_db"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

func GetConnection() *sql.DB {
	psqlInformation := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	db, err := sql.Open("postgres", psqlInformation)

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database " + dbname + "!")
	return db
}
