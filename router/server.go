package router

import "github.com/gin-gonic/gin"

func InitServer() {
	r := gin.Default()
	InitRouter(r)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}