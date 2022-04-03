package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/oasis/game_boat/biz/dal/user_dal"
	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/biz/model/user_model"
	"github.com/oasis/game_boat/global"
	utils2 "github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

type UpdateUserInfoHandler struct {
}

// UpdateUserInfo 更新用户信息
func UpdateUserInfo(ctx *gin.Context) {
	s := new(UpdateUserInfoHandler)
	method := "[UpdateUserInfo] "
	logs := utils2.NewLoggerWithXRId(ctx, global.App.Log)

	req := new(handler_model.UpdateUserInfoRequest)
	err := ctx.BindJSON(req)
	if err != nil {
		logs.Error(method + "get req fail, err = " + err.Error())
		response.BusinessFail(ctx, err.Error())
		return
	}
	req.UserId = cast.ToUint(ctx.Value("id"))

	logs.Info(fmt.Sprintf(method+"req = %v", utils2.JsonStrFormatIgnoreErr(req)))

	userInfo, userLoginInfo, err := s.getUserInfo(req.UserId)
	if err != nil {
		logs.Error(fmt.Sprintf("%v%+v", method, err.Error()))
		response.BusinessFail(ctx, err.Error())
		return
	}

	s.fillUpInfo(&req.ChangeUserInfo, &userInfo.UserBaseInfo, userLoginInfo)

	err = s.updateInfo(userInfo, userLoginInfo)
	if err != nil {
		logs.Error(fmt.Sprintf("%v%+v", method, err))
	}

	resp := new(handler_model.UpdateUserInfoResponse)

	s.fillUpResponse(userInfo, resp)

	response.Success(ctx, resp)
}

//填充用户新信息
func (s UpdateUserInfoHandler) fillUpInfo(newInfo *handler_model.NewUserInfo, userBaseInfo *user_model.UserBaseInfo, userLoginInfo *user_model.UserLoginInfo) {
	if newInfo.Password != "" {
		userLoginInfo.Password = utils2.BcryptMake([]byte(newInfo.Password))
	}
	userBaseInfo.Name = newInfo.Name
	userBaseInfo.WallImageUrl = newInfo.WallImageUrl
	userBaseInfo.AvatarUrl = newInfo.AvatarUrl
}

//获取用户旧信息
func (s UpdateUserInfoHandler) getUserInfo(userId uint) (*user_model.UserInfo, *user_model.UserLoginInfo, error) {
	userInfo, err := user_dal.GetUserInfo(userId)
	if err != nil {
		return nil, nil, err
	}
	userLoginInfo, err := user_dal.GetUserLoginInfoByLoginId(userInfo.LoginInfoId)
	if err != nil {
		return nil, nil, err
	}
	return userInfo, userLoginInfo, nil
}

//更新用户信息
func (s UpdateUserInfoHandler) updateInfo(userInfo *user_model.UserInfo, userLoginInfo *user_model.UserLoginInfo) error {
	err := user_dal.UpdateUserInfo(userInfo)
	if err != nil {
		return err
	}

	err = user_dal.UpdateUserLoginInfo(userLoginInfo)
	if err != nil {
		return err
	}

	return nil
}

//填充回参
func (s UpdateUserInfoHandler) fillUpResponse(userInfo *user_model.UserInfo, resp *handler_model.UpdateUserInfoResponse) {
	resp.Result = "Success"
	resp.UserInfo = *userInfo
}
