package user_model

import (
	"github.com/oasis/game_boat/biz/const_def"
	"github.com/oasis/game_boat/biz/model/common_model"
)

type UserRelation struct {
	common_model.ID

	AUserId  common_model.ID        `json:"a_user_id"`
	BUserId  common_model.ID        `json:"b_user_id"`
	Relation const_def.RelationEnum `json:"relation"`

	common_model.Timestamps
	common_model.SoftDeletes
}

func (s UserRelation) TableName() string {
	return const_def.UserRelationTableName
}
