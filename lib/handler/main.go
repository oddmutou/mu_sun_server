package handler

import (
    "fmt"
		"github.com/gin-gonic/gin"
    "strconv"
		"net/http"
    "../user"
)

func GetRank (ctx *gin.Context) {
    user_id, _ := strconv.Atoi(ctx.Param("user_id"))
    user := user.Get(user_id)
		fmt.Printf("%+v", user)

		ctx.JSON(http.StatusOK, gin.H{
				"status":  "get",
				"user": user,
		})
}


func RegisterScore (ctx *gin.Context) {
    user_id, _ := strconv.Atoi(ctx.Param("user_id"))
    user := user.Get(user_id)
		fmt.Printf("%+v", user)

		ctx.JSON(http.StatusOK, gin.H{
				"status":  "posted",
				"hoge": "huga",
		})
}
