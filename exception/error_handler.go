package exception

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pengdst/golang-restful-api/model/web"
)

func ErrorHandler(c *gin.Context, error interface{}) {

	if notFoundError(c, error) {
		return
	}

	if validationError(c, error) {
		return
	}

	if authorizationError(c, error) {
		return
	}

	internalServerError(c, error)
}

func authorizationError(c *gin.Context, error interface{}) bool {
	exception, ok := error.(AuthorizationError)
	if ok {
		code := http.StatusUnauthorized
		webResponse := web.WebResponse{
			Code:   code,
			Status: "Bad Request",
			Data:   exception.Error,
		}

		c.AbortWithStatusJSON(code, webResponse)
	}

	return ok
}

func validationError(c *gin.Context, error interface{}) bool {
	exception, ok := error.(ValidationError)
	if ok {
		code := http.StatusBadRequest
		webResponse := web.WebResponse{
			Code:   code,
			Status: "Bad Request",
			Data:   exception.Error,
		}

		c.AbortWithStatusJSON(code, webResponse)
	}

	return ok
}

func notFoundError(c *gin.Context, error interface{}) bool {
	exception, ok := error.(NotFoundError)
	if ok {
		code := http.StatusNotFound
		webResponse := web.WebResponse{
			Code:   code,
			Status: "Not Found",
			Data:   exception.Error,
		}

		c.AbortWithStatusJSON(code, webResponse)
	}

	return ok
}

func internalServerError(c *gin.Context, error interface{}) {
	code := http.StatusInternalServerError
	webResponse := web.WebResponse{
		Code:   code,
		Status: "Internal Server Error",
		Data:   error,
	}

	c.AbortWithStatusJSON(code, webResponse)
}
