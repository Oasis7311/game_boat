package handler_model

import (
	"github.com/oasis/game_boat/biz/model/game_model"
)

type SearchRequest struct {
	Name string `json:"name"`
}

type SearchResponse struct {
	Games []*game_model.GameInfo `json:"games"`
}
