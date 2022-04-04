package game_model

import (
	"github.com/oasis/game_boat/biz/const_def"
	"github.com/oasis/game_boat/biz/model/common_model"
)

type TagInfo struct {
	common_model.ID
	TagId   uint   `json:"tag_id"`
	TagName string `json:"tag_name,omitempty"`
}

func (s TagInfo) TableName() string {
	return const_def.TagInfoTableName
}
