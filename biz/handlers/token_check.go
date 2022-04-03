package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/oasis/game_boat/biz/dal/user_dal"
	"github.com/oasis/game_boat/global"
	"github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

func CheckToken(ctx *gin.Context) {
	id := cast.ToUint(ctx.Value("id"))

	logs := utils.NewLoggerWithXRId(ctx, global.App.Log)

	userInfo, err := user_dal.GetUserInfo(id)
	if err != nil {
		logs.Error(fmt.Sprintf("[CheckToken] get user info fail, err = %v", err))
		response.BusinessFail(ctx, fmt.Sprintf("[CheckToken] get user info fail, err = %v", err))
		return
	}
	userLoginInfo, err := user_dal.GetUserLoginInfoByLoginId(userInfo.LoginInfoId)
	if err != nil {
		logs.Error(fmt.Sprintf("[CheckToken] get user login info fail, err = %v", err))
		response.BusinessFail(ctx, fmt.Sprintf("[CheckToken] get user info fail, err = %v", err))
		return
	}

	response.Success(ctx, gin.H{
		"user_info": userInfo,
		"email":     userLoginInfo.Email,
	})
}
