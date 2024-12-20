package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HTTPServer struct {
	srv         *http.Server
	engine      *gin.Engine
	routers     []Router
	version     string
	appName     string
	middlewares []gin.HandlerFunc
}

func NewHTTPServer(appName, version, addr string, routers []Router, middlewares []gin.HandlerFunc) HTTPServer {
	engine := gin.Default()
	engine.Use(ExceptionHandle())
	engine.Use(ClientIPHandle())
	s := HTTPServer{
		engine:      engine,
		version:     version,
		appName:     appName,
		routers:     routers,
		middlewares: middlewares,
	}
	s.registerHealthRouter()
	s.registerRouters()
	s.srv = &http.Server{
		Addr:    addr,
		Handler: s.engine,
	}
	return s
}
func (s *HTTPServer) registerRouters() {
	for _, mw := range s.middlewares {
		s.engine.Use(mw)
	}
	for _, router := range s.routers {
		router.Register(s.engine)
	}
}

func (s *HTTPServer) Run() {
	logrus.Info("listen addr: ", s.srv.Addr)
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("listen: %v", err)
		}
	}()
	s.gracefullyShutDown()
}

func (s *HTTPServer) registerHealthRouter() {
	s.engine.GET("/health", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			map[string]interface{}{
				"app_name":  s.appName,
				"status":    "running",
				"timestamp": time.Now().Format(time.DateTime),
				"version":   s.version,
			})
	})
}

func (s *HTTPServer) gracefullyShutDown() {
	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		logrus.Info("server forced to shutdown:", err)
	}

	logrus.Info("server closed")
}
