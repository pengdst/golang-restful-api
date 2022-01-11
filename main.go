package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"pengdst/golang-restful-api/app"
	"pengdst/golang-restful-api/controller"
	"pengdst/golang-restful-api/exception"
	"pengdst/golang-restful-api/helper"
	"pengdst/golang-restful-api/repository"
	"pengdst/golang-restful-api/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	categortRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categortRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/categories", categoryController.GetAll)
	router.POST("/categories", categoryController.Create)
	router.GET("/categories/:categoryId", categoryController.FindById)
	router.PUT("/categories/:categoryId", categoryController.Update)
	router.DELETE("/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
