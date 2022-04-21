package services

import (
	"github.com/gin-gonic/gin"

	"github.com/oasis/game_boat/biz/dal/mysql/comment_dal"
	"github.com/oasis/game_boat/biz/dal/mysql/user_dal"
	"github.com/oasis/game_boat/biz/model/comment_model"
	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/biz/model/user_model"
	"github.com/oasis/game_boat/utils"
)

type CommentService struct {
	ctx *gin.Context
}

func NewCommentService(ctx *gin.Context) *CommentService {
	return &CommentService{ctx: ctx}
}

func (s *CommentService) PostComment(req *handler_model.PostCommentRequest) (commentDetail *comment_model.CommentDetail, err error) {
	comment := &comment_model.Comment{
		GroupId:      utils.PtrStr(req.GroupId),
		CommentLevel: utils.PtrUint(req.CommentLevel),
		Msg:          utils.PtrStr(req.Msg),
		UserId:       utils.PtrUint(req.UserId),
		ReplyUserId:  req.ReplyToUser,
		Pid:          req.Pid,
		RootId:       req.RootId,
	}
	err = comment_dal.CreateComment(comment)
	if err != nil {
		return nil, err
	}

	commentDetail = &comment_model.CommentDetail{
		Comment:         comment,
		UserInfo:        &user_model.UserInfo{},
		ReplyToUserInfo: nil,
	}

	commentDetail.UserInfo, _ = user_dal.GetUserInfo(*comment.UserId)
	if req.ReplyToUser != nil {
		commentDetail.ReplyToUserInfo, _ = user_dal.GetUserInfo(*req.ReplyToUser)
	}
	return commentDetail, err
}

func (s *CommentService) ListComment(req *handler_model.ListCommentRequest) (commentDetail []*comment_model.CommentDetail, count int64, err error) {
	dto := comment_model.NewListCommentDto()
	dto.Limit = req.PageInfo.PageSize
	dto.Offset = (req.PageInfo.Page - 1) * req.PageInfo.PageSize

	if req.ByComment != nil {
		dto.RootId = utils.PtrUint(req.ByComment.CommentId)
		s.getSort(req.ByComment.SortField, req.ByComment.Desc, dto)
	}
	if req.ByGroup != nil {
		dto.GroupId = utils.PtrStr(req.ByGroup.GroupId)
		s.getSort(req.ByGroup.SortField, req.ByGroup.Desc, dto)
	}
	if req.ByUser != nil {
		dto.UserId = utils.PtrUint(req.ByUser.UserId)
		s.getSort(req.ByUser.SortField, req.ByUser.Desc, dto)
	}

	comments, userId, replyToUserId, count, err := comment_dal.ListComment(dto)
	if err != nil {
		return nil, 0, err
	}

	userInfo, _ := user_dal.GetUserInfoMap(userId)
	replyToUserInfo, _ := user_dal.GetUserInfoMap(replyToUserId)

	for _, comment := range comments {
		commentDetail = append(commentDetail, &comment_model.CommentDetail{
			Comment:         comment,
			UserInfo:        userInfo[*comment.UserId],
			ReplyToUserInfo: replyToUserInfo[*comment.ReplyUserId],
		})
	}
	return
}

func (s *CommentService) UpdateCommentMsg(req *handler_model.UpdateCommentRequest) error {
	return comment_dal.UpdateCommentMsg(req.CommentId, req.Msg)
}

func (s *CommentService) DeleteComment(req *handler_model.DeleteCommentRequest) error {
	return comment_dal.DeleteComment(req.CommentId)
}

func (s *CommentService) getSort(field *string, desc *bool, dto *comment_model.ListCommentDto) {
	if field != nil {
		dto.SortField = *field
	}
	if desc != nil {
		dto.Desc = *desc
	}
}
