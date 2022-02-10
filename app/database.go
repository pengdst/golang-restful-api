package app

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"pengdst/golang-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	var (
		user       = os.Getenv("DB_USERNAME")
		password   = os.Getenv("DB_PASSWORD")
		dbHost     = os.Getenv("DB_HOST")
		dbPort     = os.Getenv("DB_PORT")
		dbDatabase = os.Getenv("DB_NAME")
		sslMode    = os.Getenv("DB_SSL_MODE")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, user, password, dbDatabase, sslMode)

	db, err := sql.Open("postgres", psqlInfo)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
