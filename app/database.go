package app

import (
	"database/sql"
	"fmt"
	"pengdst/golang-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	var (
		user       = "root"
		password   = "root"
		dbHost     = "localhost"
		dbPort     = "3306"
		dbDatabase = "golang-restful-api"
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
