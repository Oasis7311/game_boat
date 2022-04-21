package services

import (
	"github.com/oasis/game_boat/biz/dal/mysql/action_dal"
)

type ActionService struct {
}

func NewActionService() *ActionService {
	return new(ActionService)
}

func (a *ActionService) LikeContent(UserId uint, ContentId int64) error {
	return action_dal.LikeContent(UserId, ContentId)
}
