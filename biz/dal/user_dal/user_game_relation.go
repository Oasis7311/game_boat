package user_dal

import (
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cast"

	"github.com/oasis/game_boat/biz/model/user_model"
	"github.com/oasis/game_boat/global"
)

// GetUserReservedGameList 获取用户预约游戏列表
func GetUserReservedGameList(userId uint) ([]*user_model.UserGameRelation, error) {
	res := make([]*user_model.UserGameRelation, 0)
	err := global.App.DB.Debug().Where("user_id = ? and is_reserved = 1", userId).Order("updated_at DESC ").Find(&res).Error
	return res, errors.Wrap(err, "GetUserReservedGameList fail")
}

// GetUserCollectedGameList 获取用户收藏游戏列表
func GetUserCollectedGameList(userId uint) ([]*user_model.UserGameRelation, error) {
	res := make([]*user_model.UserGameRelation, 0)
	err := global.App.DB.Debug().Where("user_id = ? and is_collected = 1", userId).Order("updated_at DESC ").Find(&res).Error
	return res, errors.Wrap(err, "GetUserCollectedGameList fail")
}

// GetUserRelatedGameCount 获取用户收藏、预约游戏数量
func GetUserRelatedGameCount(userId uint) (uint, uint, error) {
	collectCount := int64(0)
	reserveCount := int64(0)
	err := global.App.DB.Debug().Model(&user_model.UserGameRelation{}).Where("user_id = ? and is_collected = 1", userId).Count(&collectCount).Error
	if err != nil {
		return 0, 0, errors.Wrap(err, "GetUserRelatedGameCount get collect count fail")
	}
	err = global.App.DB.Debug().Model(&user_model.UserGameRelation{}).Where("user_id = ? and is_reserved = 1", userId).Count(&reserveCount).Error
	return cast.ToUint(collectCount), cast.ToUint(reserveCount), errors.Wrap(err, "GetUserRelatedGameCount get reserve count fail")
}

// UpdateUserGameRelation 更新用户游戏状态
func UpdateUserGameRelation(userId, gameId uint, key, timeKey string, value bool, timeValue time.Time) error {

	res := global.App.DB.Debug().Model(&user_model.UserGameRelation{}).Where("user_id = ? and game_id = ?", userId, gameId).Updates(
		map[string]interface{}{
			key:     value,
			timeKey: timeValue,
		})
	if res.RowsAffected == 0 {
		res = global.App.DB.Debug().Model(&user_model.UserGameRelation{}).Create(map[string]interface{}{
			"user_id": userId,
			"game_id": gameId,
			key:       value,
			timeKey:   timeValue,
		})
	}
	return errors.Wrap(res.Error, "UpdateUserGameRelation fail")
}
