package handler_model

import (
	"github.com/oasis/game_boat/biz/const_def"
)

type GameActionRequest struct {
	GameId uint                 `json:"game_id"`
	Action const_def.ActionEnum `json:"action"`
}

type GameActionResponse struct {
}
