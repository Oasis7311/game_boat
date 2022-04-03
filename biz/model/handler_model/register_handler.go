package handler_model

import (
	"github.com/oasis/game_boat/biz/model/user_model"
	"github.com/oasis/game_boat/utils/request"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r RegisterRequest) GetMessages() request.ValidatorMessages { //获取参数错误信息
	return request.ValidatorMessages{
		//"Name.required": "用户名为空",
		"Email.required":    "邮箱为空",
		"Password.required": "密码为空",
		//"Avatar_url.required": "头像为空",
	}
}

type RegisterResponse struct {
	UserBaseInfo *user_model.UserInfo `json:"user_base_info"`
}
