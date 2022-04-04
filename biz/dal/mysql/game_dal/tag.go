package game_dal

import (
	"github.com/pkg/errors"
	"github.com/spf13/cast"

	"github.com/oasis/game_boat/biz/model/game_model"
	"github.com/oasis/game_boat/global"
)

func GetTagList() ([]*game_model.TagInfo, error) {
	res := make([]*game_model.TagInfo, 0)
	err := global.App.DB.Debug().Find(&res).Error
	return res, errors.Wrap(err, "GetTagList fail")
}

func GetGameListWithTag(tagId, page, pageSize uint) ([]*game_model.GameInfo, error) {
	res := make([]*game_model.GameInfo, 0)
	err := global.App.DB.Debug().
		Where("tag_id = ?", tagId).
		Offset(cast.ToInt((page - 1) * pageSize)).Limit(cast.ToInt(pageSize)).
		Find(&res).Error

	return res, errors.Wrap(err, "GetGameListWithTag fail")
}
