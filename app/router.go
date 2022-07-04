package app

import (
	"github.com/gin-gonic/gin"
	"pengdst/golang-restful-api/controller"
	"pengdst/golang-restful-api/exception"
	"pengdst/golang-restful-api/middleware"
)

func NewRouter(categoryController controller.CategoryController) *gin.Engine {
	router := gin.Default()

	router.Use(gin.CustomRecovery(exception.ErrorHandler))

	router.Static("/docs", "./docs")

	api := router.Group("api").Use(middleware.AuthMiddleware)
	api.GET("/categories", categoryController.GetAll)
	api.POST("/categories", categoryController.Create)
	api.GET("/categories/:categoryId", categoryController.FindById)
	api.PUT("/categories/:categoryId", categoryController.Update)
	api.DELETE("/categories/:categoryId", categoryController.Delete)

	return router
}
