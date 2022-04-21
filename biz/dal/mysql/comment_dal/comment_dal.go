package comment_dal

import (
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"gorm.io/gorm/clause"

	"github.com/oasis/game_boat/biz/model/comment_model"
	"github.com/oasis/game_boat/global"
	"github.com/oasis/game_boat/utils"
)

func CreateComment(comment *comment_model.Comment) error {
	comment.CreateTime = utils.PtrUint(cast.ToUint(time.Now().Unix()))
	comment.ModifyTime = comment.CreateTime
	comment.Status = utils.PtrInt(0)
	err := global.App.DB.Debug().Create(comment).Error
	return errors.Wrap(err, "create comment fail")
}

func DeleteComment(commentId uint) error {
	err := global.App.DB.Debug().Where("id = ?", commentId).
		Updates(map[string]interface{}{
			"status":      1,
			"modify_time": time.Now().Unix(),
		}).Error
	return errors.Wrap(err, "delete comment fail")
}

func UpdateCommentStatus(commentId uint, status int) error {
	err := global.App.DB.Debug().Model(&comment_model.Comment{}).Where("id = ?", commentId).
		Updates(map[string]interface{}{
			"status":      status,
			"modify_time": time.Now().Unix(),
		}).Error
	return errors.Wrap(err, "update comment status fail")
}

func UpdateCommentMsg(commentId uint, msg string) error {
	err := global.App.DB.Debug().Model(&comment_model.Comment{}).Where("id = ?", commentId).
		Updates(map[string]interface{}{
			"msg":         msg,
			"modify_time": time.Now().Unix(),
		}).Error
	return errors.Wrap(err, "update comment msg fail")
}

func GetCommentById(commentId uint) (*comment_model.Comment, error) {
	comment := new(comment_model.Comment)
	err := global.App.DB.Debug().Where("id = ?", commentId).First(comment).Error
	return comment, errors.Wrap(err, "get comment record fail")
}

func GetCommentByUser(userId uint, sortType ...interface{}) ([]*comment_model.Comment, error) {
	comments := make([]*comment_model.Comment, 0)
	if len(sortType) == 0 {
		sortType = make([]interface{}, 2)
		sortType[0] = "create_time"
		sortType[1] = true
	}
	err := global.App.DB.Debug().
		Where("user_id = ?", userId).
		Order(clause.OrderByColumn{Column: clause.Column{Name: sortType[0].(string)}, Desc: sortType[2].(bool)}).
		Find(&comments).Error
	return comments, errors.Wrap(err, "get comment by user fail")
}

func GetSonComment(pid uint) ([]*comment_model.Comment, error) {
	comments := make([]*comment_model.Comment, 0)
	err := global.App.DB.Debug().Where("pid = ?", pid).Find(&comments).Error
	return comments, errors.Wrap(err, "get comment by pid fail")
}

func ListComment(dto *comment_model.ListCommentDto) ([]*comment_model.Comment, []uint, []uint, int64, error) {
	comments := make([]*comment_model.Comment, 0)
	conn := global.App.DB.Debug()
	if dto.RootId != nil {
		conn = conn.Where("root_id = ?", *dto.RootId)
	}
	if dto.GroupId != nil {
		conn = conn.Where("group_id = ?", *dto.GroupId)
	}
	if dto.UserId != nil {
		conn = conn.Where("user_id = ?", *dto.UserId)
	}
	conn = conn.Where("status = ?", dto.Status).Offset(dto.Offset).Limit(dto.Limit)

	err := conn.Find(&comments).Error
	if err != nil {
		return nil, nil, nil, 0, errors.Wrap(err, "find comments fail")
	}

	userId, replyToUserId := make([]uint, 0), make([]uint, 0)
	for _, comment := range comments {
		userId = append(userId, *comment.UserId)
		if *comment.ReplyUserId != 0 {
			replyToUserId = append(replyToUserId, *comment.ReplyUserId)
		}
	}

	count := int64(0)
	conn.Count(&count)

	return comments, userId, replyToUserId, count, errors.Wrap(err, "find comments fail")
}
