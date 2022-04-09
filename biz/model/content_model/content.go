package content_model

import (
	"github.com/oasis/game_boat/biz/const_def"
	"github.com/oasis/game_boat/biz/model/common_model"
	"github.com/oasis/game_boat/biz/model/user_model"
)

type Content struct {
	ContentId     int64  `json:"content_id" gorm:"content_id"`
	AuthorId      uint   `json:"author_id" gorm:"author_id"`
	Title         string `json:"title" gorm:"title"`
	SubTitle      string `json:"sub_title" gorm:"sub_title"`
	CoverImageUrl string `json:"cover_image_url" gorm:"cover_image_url"`
	Buffer        []byte `json:"buffer" gorm:"buffer"`
}

func (c Content) TableName() string {
	return const_def.ContentTableName
}

type ContentResponse struct {
	ContentId     string                `json:"content_id"`
	AuthorInfo    *user_model.UserInfo  `json:"author_info"`
	Title         string                `json:"title,omitempty"`
	SubTitle      string                `json:"sub_title,omitempty"`
	CoverImageUrl string                `json:"cover_image_url,omitempty"`
	ImageList     []*common_model.Image `json:"image_list,omitempty"`
	Summary       string                `json:"summary,omitempty"`
	Text          string                `json:"text,omitempty"`
	PublishTime   int64                 `json:"publish_time,omitempty"`
	LikeCount     uint                  `json:"like_count,omitempty"`
}

type ContentBuffer struct {
	ContentImage []*common_model.Image `json:"content_image"`
	Summary      string                `json:"summary"`
	PublishTime  int64                 `json:"publish_time"`
	Text         string                `json:"text"`
}
