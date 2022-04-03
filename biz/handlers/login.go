package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/oasis/game_boat/biz/const_def"
	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/global"
	"github.com/oasis/game_boat/services"
	"github.com/oasis/game_boat/services/user"
	utils2 "github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/request"
	"github.com/oasis/game_boat/utils/response"
)

type LoginHandler struct {
	req *handler_model.LoginRequest
	rsp *handler_model.LoginResponse
}

func Login(ctx *gin.Context) {
	loginHandler := new(LoginHandler)
	method := "Login"

	logs := utils2.NewLoggerWithXRId(ctx, global.App.Log)

	if err := loginHandler.getRequestBody(ctx); err != nil {
		logs.Error(utils2.NewErrorMessage(method, "getRequestBody Fail", err))
		response.ValidateFail(ctx, request.GetErrorMsg(loginHandler.req, err))
		return
	}
	logs.Info(fmt.Sprintf("[Login] reqBody = %v", utils2.JsonStrFormatIgnoreErr(loginHandler.req)))

	userinfo, err := user.UserService.Login(loginHandler.req)
	if err != nil {
		logs.Error(fmt.Sprintf("[Login] LoginFail, Error = %+v", err))
		response.BusinessFail(ctx, err.Error())
		return
	}
	if userinfo == nil {
		logs.Error(utils2.NewErrorMessage(method, "获取用户信息失败", errors.New("get user info fail")))
		response.BusinessFail(ctx, "获取用户信息失败")
		return
	}

	token, err, _ := services.JwtService.CreateToken(services.AppGuardName, userinfo)
	if err != nil {
		logs.Error(fmt.Sprintf("[Login] CreateToken Fail, Error = %+v", err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	loginHandler.rsp = handler_model.NewLoginResponse(userinfo, token.AccessToken, token.TokenType, token.ExpiresIn)

	ctx.Header(const_def.XPolisToken, token.AccessToken)

	logs.Info(fmt.Sprintf("[Login] rsp = %v", utils2.JsonStrFormatIgnoreErr(loginHandler.rsp)))

	response.Success(ctx, loginHandler.rsp)
}

func (s *LoginHandler) getRequestBody(ctx *gin.Context) error {
	s.req = new(handler_model.LoginRequest)

	return ctx.ShouldBindJSON(s.req)
}
