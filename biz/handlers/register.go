package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/biz/model/user_model"
	"github.com/oasis/game_boat/global"
	"github.com/oasis/game_boat/services/user"
	"github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/request"
	"github.com/oasis/game_boat/utils/response"
)

type registerHandler struct {
	req *handler_model.RegisterRequest
	rsp *handler_model.RegisterResponse
}

func Register(ctx *gin.Context) {
	logs := utils.NewLoggerWithXRId(ctx, global.App.Log)
	handler := new(registerHandler)

	if err := handler.getRequestBody(ctx); err != nil {
		logs.Error("[Register] Get Request Body Fail, Error = " + err.Error())
		response.ValidateFail(ctx, request.GetErrorMsg(handler.req, err))
		return
	}

	logs.Info("[Register] ReqBody = " + utils.MarshalIgnoreErr(handler.req))

	userLoginInfo := &user_model.UserLoginInfo{
		Email:    handler.req.Email,
		Password: handler.req.Password,
	}
	userBaseInfo := &user_model.UserBaseInfo{
		Name:         "polis-" + utils.GenXid(),
		AvatarUrl:    "https://lf3-static.bytednsdoc.com/obj/eden-cn/lluhfyeh7uhbfpzbps/avatar_default.jpeg",
		WallImageUrl: "https://assets.tapimg.com/web-app/static/img/default_bg.05387bbb.png",
		Email:        userLoginInfo.Email,
	}

	err, userInfo := user.UserService.Register(userLoginInfo, userBaseInfo)
	if err != nil {
		logs.Error("[Register] Create User Fail, Error = " + err.Error())
		response.BusinessFail(ctx, err.Error())
		return
	}
	logs.Info("[Register] userInfo = " + utils.MarshalIgnoreErr(userInfo))

	handler.rsp = new(handler_model.RegisterResponse)
	handler.rsp.UserBaseInfo = userInfo

	response.Success(ctx, handler.rsp)
}

func (r *registerHandler) getRequestBody(c *gin.Context) error {
	r.req = new(handler_model.RegisterRequest)

	return c.ShouldBindJSON(r.req)
}
