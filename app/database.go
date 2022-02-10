package app

import (
	"database/sql"
	"fmt"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pengdst/golang-restful-api/helper"
	"pengdst/golang-restful-api/model/domain"
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

func NewGormDB() *gorm.DB {
	var (
		user       = "root"
		password   = "root"
		dbHost     = "localhost"
		dbPort     = "3306"
		dbDatabase = "golang-restful-api"
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, dbHost, dbPort, dbDatabase)
	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	err = gormDb.AutoMigrate(&domain.Category{})
	helper.PanicIfError(err)

	db, err := gormDb.DB()
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return gormDb
}
