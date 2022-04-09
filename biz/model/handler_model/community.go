package handler_model

import (
	"github.com/oasis/game_boat/biz/model/content_model"
)

type GetCommunityRequest struct {
	Page      int64 `json:"page"`
	PageSize  int64 `json:"page_size"`
	Timestamp int64 `json:"timestamp"`
}

type GetCommunityResponse struct {
	ContentList []*content_model.ContentDetail `json:"content_list"`
	Timestamp   int64                          `json:"timestamp"`
}
