package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCompanyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"key": "List Company",
	})
}