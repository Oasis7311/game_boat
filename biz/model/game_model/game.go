package game_model

import (
	"github.com/oasis/game_boat/biz/const_def"
	"github.com/oasis/game_boat/biz/model/common_model"
)

type GameInfo struct {
	common_model.ID
	Name            string  `json:"name,omitempty"`
	TagId           uint    `json:"tag_id,omitempty"`
	Introduction    string  `json:"introduction,omitempty"`
	Icon            string  `json:"icon,omitempty"`
	Image           string  `json:"image,omitempty"`
	Developer       string  `json:"developer,omitempty"`
	Lang            string  `json:"lang,omitempty"`
	AgeLimit        uint    `json:"age_limit,omitempty"`
	Score           float64 `json:"score,omitempty"`
	TopMedia        string  `json:"top_media,omitempty"`
	MediaScoreUrl   string  `json:"media_score_url,omitempty"`
	GameDescription string  `json:"game_description,omitempty"`
}

func (s GameInfo) TableName() string {
	return const_def.GameInfoTableName
}
