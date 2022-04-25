package user

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/oasis/game_boat/biz/dal/mysql/user_dal"
	"github.com/oasis/game_boat/global"
	"github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

func GetUser(ctx *gin.Context) {
	logs := utils.NewLoggerWithXRId(ctx, global.App.Log)
	method := "GetUser"

	email := ctx.Param("email")

	userInfo, err := user_dal.GetUserInfoByEmail(email)
	if err != nil {
		logs.Error(fmt.Sprintf("%v find user fail, err = %v", method, err))
		response.Success(ctx, nil)
		return
	}
	response.Success(ctx, userInfo)
	return
}
