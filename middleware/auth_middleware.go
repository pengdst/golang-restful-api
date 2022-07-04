package middleware

import (
	"github.com/gin-gonic/gin"
	"pengdst/golang-restful-api/exception"
)

func AuthMiddleware(c *gin.Context) {
	auth := c.GetHeader("X-API-Key")
	if auth != "RAHASIA" {
		panic(exception.NewValidationError("wrong api key"))
	}
}
