package app

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"pengdst/golang-restful-api/controller"
	"pengdst/golang-restful-api/exception"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.ServeFiles("/docs/*filepath", http.Dir("./docs"))
	router.GET("/categories", categoryController.GetAll)
	router.POST("/categories", categoryController.Create)
	router.GET("/categories/:categoryId", categoryController.FindById)
	router.PUT("/categories/:categoryId", categoryController.Update)
	router.DELETE("/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
