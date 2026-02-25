package router

import (
	"hello/internal/handler"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Route Not Found"})
	})

	r.HandleMethodNotAllowed = true

	r.NoMethod(func(c *gin.Context) {
		c.JSON(405, gin.H{"error": "Method Not Allowed"})
	})

	r.GET("/", handler.Healthy)
	r.GET("/hello", handler.Hello)
	r.POST("/calc", handler.Calc) 

	return r
}
