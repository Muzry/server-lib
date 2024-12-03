package server

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func ClientIPHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := getClientIP(c)
		c.Set("client_ip", ip)
		c.Next()
	}
}

func getClientIP(c *gin.Context) string {
	ip := c.Request.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = c.Request.Header.Get("X-Real-IP")
	}
	if ip == "" {
		ip = c.ClientIP()
	}
	ip = strings.Split(ip, ",")[0]
	return strings.TrimSpace(ip)
}
