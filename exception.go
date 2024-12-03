package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Exception struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Request string `json:"request"`
}

func (e *Exception) Error() string {
	return e.Msg
}

func newException(code int, msg string) *Exception {
	return &Exception{
		Code: code,
		Msg:  msg,
	}
}

func ExceptionHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if err := c.Errors.Last(); err != nil {
			c.Errors = c.Errors[:len(c.Errors)-1]
			var e *Exception
			var except *Exception
			if errors.As(err.Err, &except) {
				e = except
				e.Request = c.Request.Method + " " + c.Request.URL.String()
				c.JSON(http.StatusOK, e)
				return
			}
			c.JSON(http.StatusOK, err)
			return
		}
	}
}
