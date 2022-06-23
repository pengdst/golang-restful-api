package app

import (
	"database/sql"
	"fmt"
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
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, dbHost, dbPort, dbDatabase)
	db, err := sql.Open("mysql", dsn)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
