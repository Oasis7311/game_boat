package handler_model

import (
	"github.com/oasis/game_boat/biz/model/game_model"
)

type GetGameThreeListRequest struct {
}

type GetGameThreeListResponse struct {
	NewGame      []*game_model.GameInfo `json:"new_game,omitempty"`
	HotGame      []*game_model.GameInfo `json:"hot_game,omitempty"`
	RecentUpdate []*game_model.GameInfo `json:"recent_update,omitempty"`
	Page         int64                  `json:"page"`
	PageSize     int64                  `json:"page_size"`
}
