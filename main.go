package main

import (
	"service1/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/getdataa", handler.GetData)
	r.Run("localhost:8081")
}
