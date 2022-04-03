package handler_model

import (
	"github.com/oasis/game_boat/biz/model/user_model"
)

type NewUserInfo struct {
	Name         string `json:"name,omitempty"`
	Password     string `json:"password,omitempty"`
	AvatarUrl    string `json:"avatar_url,omitempty"`
	WallImageUrl string `json:"wall_image_url,omitempty"`
}

type UpdateUserInfoRequest struct {
	CommonRequest
	ChangeUserInfo NewUserInfo `json:"change_user_info"`
}

type UpdateUserInfoResponse struct {
	Result              string `json:"result,omitempty"`
	user_model.UserInfo `json:"user_info"`
}
