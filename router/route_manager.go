package router

import (
	"github.com.br/thenopholo/uid_eveo/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/company", handler.ShowCompanyHandler)
		v1.POST("/company", handler.CreateCompmanyHandler)
		v1.PUT("/company", handler.UpdateCompanyHandler)
		v1.DELETE("/company", handler.DeleteCompanyHandler)
		v1.GET("/all_company", handler.ListCompanyHandler)
	}

}