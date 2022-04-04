package handler_model

import (
	"github.com/oasis/game_boat/biz/model/game_model"
)

type GetUserGameRequest struct {
}

type UserGame struct {
	GameInfo      *game_model.GameInfo `json:"game_info"`
	CollectedTime uint                 `json:"collected_time,omitempty"`
	ReserveTime   uint                 `json:"reserve_time,omitempty"`
}

type GetUserGameResponse struct {
	UserCollectedGame   []*UserGame   `json:"user_collected_game"`
	UserReservedGame    []*UserGame   `json:"user_reserved_game"`
	UserCollectdGameMap map[uint]bool `json:"user_collectd_game_map"`
}

type GetUserGameCountResponse struct {
	CollectedCount uint `json:"collected_count"`
	ReservedCount  uint `json:"reserved_count"`
}
