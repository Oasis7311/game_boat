package handler_model

import (
	"github.com/oasis/game_boat/biz/model/user_model"
	"github.com/oasis/game_boat/utils/request"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s LoginRequest) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"Email.required":    "邮箱为空",
		"Password.required": "密码为空",
	}
}

type LoginResponse struct {
	UserInfo    *user_model.UserInfo `json:"user_info"`
	AccessToken string               `json:"access_token"`
	ExpiresIn   int                  `json:"expires_in"`
	TokenType   string               `json:"token_type"`
}

func NewLoginResponse(userInfo *user_model.UserInfo, accessToken, tokenType string, expiresIn int) *LoginResponse {
	return &LoginResponse{
		UserInfo:    userInfo,
		AccessToken: accessToken,
		ExpiresIn:   expiresIn,
		TokenType:   tokenType,
	}
}
