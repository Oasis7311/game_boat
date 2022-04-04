package handler_model

import (
	"github.com/oasis/game_boat/biz/const_def"
	"github.com/oasis/game_boat/biz/model/user_model"
)

type GetUserFollowerListRequest struct {
	Id             uint `json:"id"`
	PageSize       int  `json:"page_size"`
	LastFollowerId uint `json:"last_follower_id"`
	LastFollowTime uint `json:"last_follow_time"`
}

type GetUserFollowerListResponse struct {
	Followers      []*user_model.UserInfo `json:"followers"`
	LastFollowTime uint                   `json:"last_follow_time"`
}

type GetUserFollowListRequest struct {
	Id                 uint `json:"id"`
	PageSize           int  `json:"page_size"`
	LastFollowedUserId uint `json:"last_followed_user_id"`
	LastFollowTime     uint `json:"last_follow_time"`
}

type GetUserFollowListResponse struct {
	FollowPeople   []*user_model.UserInfo `json:"follow_people"`
	LastFollowTime uint                   `json:"last_follow_time"`
}

type PeopleRelateRequest struct {
	CommonRequest
	BUserId  uint                   `json:"b_user_id"`
	Relation const_def.RelationEnum `json:"relation"`
}

type PeopleRelateResponse struct {
}

type GetRelationCountRequest struct {
	AUserId uint `json:"a_user_id"`
}

type GetRelationCountResponse struct {
	FollowCount   uint `json:"follow_count"`
	FollowerCount uint `json:"follower_count"`
}
