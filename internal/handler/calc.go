package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CalcRequest struct {
	A int `json:"a" binding:"required"`
	B int `json:"b" binding:"required"`
}

func Calc(c *gin.Context) {
	var req CalcRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := req.A + req.B

	c.JSON(http.StatusOK, gin.H{"result": result})
}