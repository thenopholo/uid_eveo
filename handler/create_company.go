package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCompmanyHandler (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"key": "POST",
	})
}