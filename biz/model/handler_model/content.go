package handler_model

import (
	"github.com/oasis/game_boat/biz/model/content_model"
)

type GetContentRequest struct {
	ContentId int64 `json:"content_id" query:"content_id`
}

type GetContentResponse struct {
	Content *content_model.ContentResponse `json:"content"`
}
