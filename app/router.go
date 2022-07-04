package app

import (
	"github.com/gin-gonic/gin"
	"pengdst/golang-restful-api/controller"
	"pengdst/golang-restful-api/exception"
)

func NewRouter(categoryController controller.CategoryController) *gin.Engine {
	router := gin.Default()

	router.Use(gin.CustomRecovery(exception.ErrorHandler))

	router.Static("/docs", "./docs")
	router.GET("/categories", categoryController.GetAll)
	router.POST("/categories", categoryController.Create)
	router.GET("/categories/:categoryId", categoryController.FindById)
	router.PUT("/categories/:categoryId", categoryController.Update)
	router.DELETE("/categories/:categoryId", categoryController.Delete)

	return router
}
