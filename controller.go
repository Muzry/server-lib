package server

import "github.com/gin-gonic/gin"

type Controller interface {
	Routers(routers gin.IRouter)
	AuthRouters(routers gin.IRouter)
}

type Router struct {
	middlewares Middlewares
	controllers []Controller
}

type Middlewares struct {
	middleWares     []gin.HandlerFunc
	authMiddlewares []gin.HandlerFunc
}

func NewMiddlewares(middleWares, authMiddlewares []gin.HandlerFunc) Middlewares {
	return Middlewares{
		middleWares:     middleWares,
		authMiddlewares: authMiddlewares,
	}
}

func NewRouter(middlewares Middlewares, controllers []Controller) Router {
	return Router{
		middlewares: middlewares,
		controllers: controllers,
	}
}

func (r *Router) AddToRouter(router gin.IRouter) {
	r.addRouters(router)
	r.addAuthRouters(router)
}

func (r *Router) addRouters(router gin.IRouter) {
	g := router.Group("/api")
	for _, m := range r.middlewares.middleWares {
		g.Use(m)
	}

	for _, c := range r.controllers {
		c.Routers(g)
	}
}

func (r *Router) addAuthRouters(router gin.IRouter) {
	g := router.Group("/api")
	for _, m := range r.middlewares.authMiddlewares {
		g.Use(m)
	}

	for _, c := range r.controllers {
		c.AuthRouters(g)
	}
}
