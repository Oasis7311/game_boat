package handler_model

import (
	"github.com/oasis/game_boat/biz/model/content_model"
)

type GetUserMomentRequest struct {
	UserId  uint `json:"id"`
	Review  bool `json:"review"`
	Content bool `json:"content"`
}

type GetUserMomentResponse struct {
	Contents []*content_model.ContentDetail `json:"contents"`
}
