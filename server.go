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
	srv     *http.Server
	engine  *gin.Engine
	router  Router
	version string
	appName string
}

func NewHTTPServer(appName, version, addr string, router Router) HTTPServer {
	engine := gin.Default()
	engine.Use(ExceptionHandle())
	engine.Use(ClientIPHandle())
	s := HTTPServer{
		engine:  engine,
		version: version,
		appName: appName,
		router:  router,
	}
	s.addRouters()
	s.srv = &http.Server{
		Addr:    addr,
		Handler: s.engine,
	}
	return s
}

func (s *HTTPServer) addRouters() {
	s.addHealthRouter()
	s.router.AddToRouter(s.engine)
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

func (s *HTTPServer) addHealthRouter() {
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
