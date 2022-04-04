package user_model

import (
	"strconv"

	"github.com/oasis/game_boat/biz/const_def"
	"github.com/oasis/game_boat/biz/model/common_model"
)

type UserInfo struct {
	common_model.ID
	UserBaseInfo
	LoginInfoId uint `json:"login_info_id"`
	common_model.Timestamps
	common_model.SoftDeletes
}

func (s *UserInfo) GetUid() string {
	return strconv.Itoa(int(s.ID.ID))
}

func (s *UserInfo) TableName() string {
	return const_def.UserInfoTableName
}

type UserLoginInfo struct {
	common_model.ID
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	common_model.Timestamps
	common_model.SoftDeletes
}

func (s *UserLoginInfo) TableName() string {
	return const_def.LoginInfoTableName
}

type UserBaseInfo struct {
	Name         string `json:"name,omitempty"`
	AvatarUrl    string `json:"avatar_url,omitempty"`
	WallImageUrl string `json:"wall_image_url,omitempty"`
	LikeCount    int    `json:"like_count"`
	Email        string `json:"email"`
}
