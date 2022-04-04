package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	user_dal2 "github.com/oasis/game_boat/biz/dal/mysql/user_dal"
	"github.com/oasis/game_boat/services"
	"github.com/oasis/game_boat/utils/response"
)

func Logout(c *gin.Context) {
	err := services.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token))
	if err != nil {
		response.BusinessFail(c, "登出失败")
		return
	}

	id := cast.ToUint(c.Query("id"))
	userInfo, err := user_dal2.GetUserInfo(id)

	response.Success(c, gin.H{
		"message":   "登出成功",
		"user_info": userInfo,
	})
}
