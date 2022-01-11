package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"pengdst/golang-restful-api/app"
	"pengdst/golang-restful-api/controller"
	"pengdst/golang-restful-api/helper"
	"pengdst/golang-restful-api/middleware"
	"pengdst/golang-restful-api/repository"
	"pengdst/golang-restful-api/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	categortRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categortRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
