package handler

import (
	"../user"
	"../pass"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetRank(ctx *gin.Context) {
	user_id, _ := strconv.Atoi(ctx.Param("user_id"))
	user.GetRank(user_id)
	selected_user := user.Get(user_id)

	ctx.JSON(http.StatusOK, gin.H{
		"rank": user.GetRank(user_id),
		"user": selected_user,
	})
}

func Register(ctx *gin.Context) {
	var input_user user.User
	ctx.BindJSON(&input_user)
	created_user := user.Insert(input_user.Name, input_user.Score)

	ctx.JSON(http.StatusOK, gin.H{
		"rank": user.GetRank(created_user.Id),
		"user": created_user,
	})
}

func LastPass(ctx *gin.Context) {
	user_id := ctx.Param("user_id")

	lastPass := pass.GetByUserId(user_id)
	ctx.JSON(http.StatusOK, gin.H{
	        "user_id": lastPass.UserId,
	        "point_id": lastPass.PointId,
	})
}

func Pass(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	point_id := ctx.Param("point_id")

	pass.Insert(user_id, point_id)

	ctx.JSON(http.StatusOK, gin.H{
		"user_id":  user_id,
		"point_id": point_id,
	})
}

func DeletePass(ctx *gin.Context) {
	user_id := ctx.Param("user_id")

	deletedPass := pass.Delete(user_id)
	ctx.JSON(http.StatusOK, gin.H{
	        "user_id": deletedPass.UserId,
	        "point_id": deletedPass.PointId,
	})
}

func GetRanks(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.Param("limit"))

	users := user.GetRankUsers(limit)
	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
