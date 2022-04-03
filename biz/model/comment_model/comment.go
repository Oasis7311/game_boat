package comment_model

import (
	"github.com/oasis/game_boat/biz/model/common_model"
	"github.com/oasis/game_boat/biz/model/user_model"
)

type Comment struct {
	common_model.ID
	GroupId         int64
	Text            string
	ImageIdList     []int64
	CreateTime      string
	CreateTimeStamp int64
	UserInfo        *user_model.UserBaseInfo
}
