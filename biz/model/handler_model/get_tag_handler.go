package handler_model

import (
	"github.com/oasis/game_boat/biz/model/game_model"
)

type GetTagListResponse struct {
	TagList []*game_model.TagInfo `json:"tag_list"`
}

type GetGameInTagRequest struct {
	TagId    uint `json:"tag_id,omitempty"`
	Page     uint `json:"page,omitempty"`
	PageSize uint `json:"page_size,omitempty"`
}

type GetGameInTagResponse struct {
	GameList           []*game_model.GameInfo `json:"game_list"`
	UserCollectGameMap map[uint]bool          `json:"user_collect_game_map"`
	UserReserveGameMap map[uint]bool          `json:"user_reserve_game_map"`
}
