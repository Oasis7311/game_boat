package user_model

import (
	"time"

	"github.com/oasis/game_boat/biz/const_def"
	"github.com/oasis/game_boat/biz/model/common_model"
)

type UserGameRelation struct {
	common_model.ID
	common_model.Timestamps
	UserId        uint      `json:"user_id"`
	GameId        uint      `json:"game_id"`
	IsCollected   bool      `json:"is_collected"`
	IsReserved    bool      `json:"is_reserved"`
	CollectedTime time.Time `json:"collected_time"`
	ReservedTime  time.Time `json:"reserved_time"`
}

func (u UserGameRelation) TableName() string {
	return const_def.UserGameRelationTableName
}
