package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/oasis/game_boat/services"
	"github.com/oasis/game_boat/utils/response"
)

func Logout(c *gin.Context) {
	err := services.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token))
	if err != nil {
		response.BusinessFail(c, "登出失败")
		return
	}

	response.Success(c, gin.H{
		"message": "登出成功",
	})
}
