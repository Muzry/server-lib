package server

import "github.com/gin-gonic/gin"

type Controller interface {
	Prefix() string
	Register(router gin.IRouter)
}
