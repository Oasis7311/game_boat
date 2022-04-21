package action_dal

import (
	"github.com/pkg/errors"

	"github.com/oasis/game_boat/biz/model/action_model"
	"github.com/oasis/game_boat/global"
)

func LikeContent(UserId uint, ContentId int64) error {
	err := global.App.DB.Debug().Create(&action_model.Action{
		UserId:  UserId,
		GroupId: ContentId,
	}).Error
	return errors.Wrap(err, "like content fail")
}

func GetLikeCount(ContentId int64) (int64, error) {
	count := new(int64)
	res := global.App.DB.Debug().Where("content_id = ?", ContentId).Count(count)
	return *count, errors.Wrap(res.Error, "get count fail")
}
