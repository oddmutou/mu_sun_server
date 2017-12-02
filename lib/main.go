package main

import (
	"./handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/user/:user_id/rank", handler.GetRank)
	router.POST("/register", handler.Register)

	router.Run(":8080")
}
