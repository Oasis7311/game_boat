package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	user_dal2 "github.com/oasis/game_boat/biz/dal/mysql/user_dal"
	"github.com/oasis/game_boat/global"
	"github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

func CheckToken(ctx *gin.Context) {
	id := cast.ToUint(ctx.Value("id"))

	logs := utils.NewLoggerWithXRId(ctx, global.App.Log)

	logs.Info(fmt.Sprintf("[CheckToken] id = %v", id))
	userInfo, err := user_dal2.GetUserInfo(id)
	if err != nil {
		logs.Error(fmt.Sprintf("[CheckToken] err = %+v", err))
		response.BusinessFail(ctx, err.Error())
		return
	}
	userLoginInfo, err := user_dal2.GetUserLoginInfoByLoginId(userInfo.LoginInfoId)
	if err != nil {
		logs.Error(fmt.Sprintf("[CheckToken] err = %+v", err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	response.Success(ctx, gin.H{
		"user_info": userInfo,
		"email":     userLoginInfo.Email,
	})
}
