package handler_model

import (
	"github.com/oasis/game_boat/biz/model/comment_model"
)

type PostCommentRequest struct {
	GroupId      string `json:"group_id"`
	Msg          string `json:"msg,omitempty"`
	UserId       uint   `json:"user_id,omitempty"`
	CommentLevel uint   `json:"comment_level,omitempty"`
	ReplyToUser  *uint  `json:"reply_to_user,omitempty"`
	Pid          *uint  `json:"pid,omitempty"`
	RootId       *uint  `json:"root_id,omitempty"`
}

type PostCommentResponse struct {
	Comment *comment_model.CommentDetail `json:"comment"`
}

type ListCommentSort struct {
	SortField *string `json:"sort_field,omitempty"`
	Desc      *bool   `json:"desc,omitempty"`
}

type ListCommentByUser struct {
	UserId uint `json:"user_id,omitempty"`
	Level  int  `json:"level,omitempty"`
	ListCommentSort
}

type ListCommentByGroup struct {
	GroupId string `json:"group_id,omitempty"`
	ListCommentSort
}

type ListCommentByComment struct {
	CommentId uint `json:"comment_id,omitempty"`
	ListCommentSort
}

type ListCommentRequest struct {
	ByUser    *ListCommentByUser    `json:"by_user,omitempty"`
	ByGroup   *ListCommentByGroup   `json:"by_group,omitempty"`
	ByComment *ListCommentByComment `json:"by_comment,omitempty"`
	PageInfo  *PageInfo             `json:"page_info,omitempty"`
}

type ListCommentResponse struct {
	Comments []*comment_model.CommentDetail `json:"comments"`

	PageInfo *PageInfo `json:"page_info"`
}

type UpdateCommentRequest struct {
	Msg       string `json:"msg,omitempty"`
	CommentId uint   `json:"comment_id,omitempty"`
}

type UpdateCommentResponse struct {
}

type DeleteCommentRequest struct {
	CommentId uint `json:"comment_id"`
}
type DeleteCommentResponse struct {
}
