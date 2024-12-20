package server

import "github.com/gin-gonic/gin"

type Router struct {
	middlewares []gin.HandlerFunc
	controllers []Controller
}

func NewRouter(controllers []Controller, middlewares []gin.HandlerFunc) Router {
	return Router{
		middlewares: middlewares,
		controllers: controllers,
	}
}

func (r *Router) Register(router gin.IRouter) {
	apiRouter := router.Group("/api")
	for _, middleware := range r.middlewares {
		apiRouter.Use(middleware)
	}
	for _, c := range r.controllers {
		c.Register(apiRouter)
	}
}
