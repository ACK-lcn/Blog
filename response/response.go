package response

import (
	"net/http"

	"github.com/ACK-lcn/Blog/exception"
	"github.com/gin-gonic/gin"
)

// Success return response data.
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

// Failed return response error.
func Failed(c *gin.Context, err error) {
	defer c.Abort()
	
	var e *exception.ApiException
	if v, ok := err.(*exception.ApiException); ok {
		e = v
	} else {
		e = exception.New(http.StatusInternalServerError, err.Error())
		e.HttpCode = http.StatusInternalServerError
	}

	c.JSON(e.HttpCode, e)
}
