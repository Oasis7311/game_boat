package comment_model

import (
	"github.com/oasis/game_boat/biz/model/common_model"
	"github.com/oasis/game_boat/biz/model/user_model"
)

type Comment struct {
	common_model.ID
	GroupId      *uint   `json:"group_id,omitempty"`
	RootId       *uint   `json:"root_id,omitempty"`
	CommentLevel *uint   `json:"comment_level,omitempty"`
	Msg          *string `json:"msg,omitempty"`
	UserId       *uint   `json:"user_id,omitempty"`
	ReplyUserId  *uint   `json:"reply_user_id,omitempty"`
	Status       *int    `json:"status,omitempty"`
	Pid          *uint   `json:"pid,omitempty"`
	CreateTime   *uint   `json:"create_time,omitempty"`
	ModifyTime   *uint   `json:"modify_time,omitempty"`
}

func (c *Comment) TableName() string {
	return "comment"
}

type CommentDetail struct {
	Comment         *Comment             `json:"comment"`
	UserInfo        *user_model.UserInfo `json:"user_info"`
	ReplyToUserInfo *user_model.UserInfo `json:"reply_to_user_info,omitempty"`
}
