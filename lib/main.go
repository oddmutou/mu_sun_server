package main

import (
    "github.com/gin-gonic/gin"
    "./handler"
)

func main() {
    router := gin.Default()

		router.GET("/user/:user_id/rank", handler.GetRank)
		router.POST("/user/:user_id/score", handler.RegisterScore)

    router.Run(":8080")
}
