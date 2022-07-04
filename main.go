package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
	"pengdst/golang-restful-api/app"
	"pengdst/golang-restful-api/controller"
	"pengdst/golang-restful-api/helper"
	"pengdst/golang-restful-api/repository"
	"pengdst/golang-restful-api/service"
)

func init() {
	godotenv.Load(".env")
}

func main() {

	db := app.NewGormDB()
	validate := validator.New()

	categortRepository := repository.NewGormCategoryRepository(db)
	categoryService := service.NewGormCategoryService(categortRepository, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	err := router.Run(":" + os.Getenv("PORT"))
	helper.PanicIfError(err)
}
