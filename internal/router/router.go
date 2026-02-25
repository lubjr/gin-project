package router

import (
	"hello/internal/handler"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.GET("/", handler.Healthy)
	r.GET("/hello", handler.Hello)

	return r
}
