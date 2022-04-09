package handler_model

import (
	"github.com/oasis/game_boat/biz/model/content_model"
)

type GetContentRequest struct {
	ContentId int64 `json:"content_id"`
}

type GetContentResponse struct {
	Content *content_model.ContentDetail `json:"content"`
}

type PostContentRequest struct {
	Content *content_model.ContentDetail `json:"content"`
}

type PostContentResponse struct {
	Content *content_model.ContentDetail `json:"content"`
}
