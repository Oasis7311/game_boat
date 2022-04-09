package handler_model

import (
	"github.com/oasis/game_boat/biz/model/game_model"
)

type GetMainPageRequest struct {
	Page      int64 `json:"page"`
	PageSize  int64 `json:"page_size"`
	Timestamp int64 `json:"timestamp"`
}

type GetMainPageResponse struct {
	GameList  []*game_model.GameInfo `json:"game_list"`
	Timestamp int64                  `json:"timestamp"`
}
