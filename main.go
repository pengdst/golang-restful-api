package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"net/http"
	"pengdst/golang-restful-api/app"
	"pengdst/golang-restful-api/controller"
	"pengdst/golang-restful-api/helper"
	"pengdst/golang-restful-api/middleware"
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

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
