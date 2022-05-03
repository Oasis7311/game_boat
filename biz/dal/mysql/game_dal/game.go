package game_dal

import (
	"github.com/pkg/errors"

	"github.com/oasis/game_boat/biz/model/game_model"
	"github.com/oasis/game_boat/global"
)

func GetGamesDetail(gameId []uint) (map[uint]*game_model.GameInfo, error) {
	gameList := make([]*game_model.GameInfo, 0)
	gameMap := make(map[uint]*game_model.GameInfo)

	err := global.App.DB.Debug().Where("id in (?) ", gameId).Find(&gameList).Error
	if err != nil {
		return nil, errors.Wrap(err, "GetGamesDetail fail")
	}

	for _, info := range gameList {
		if _, ok := gameMap[info.ID.ID]; !ok {
			gameMap[info.ID.ID] = new(game_model.GameInfo)
		}
		*gameMap[info.ID.ID] = *info
	}

	return gameMap, nil
}

func GetAllGamesIdList() ([]uint, error) {
	gameList := make([]*game_model.GameInfo, 0)

	err := global.App.DB.Debug().Select("id").Find(&gameList).Error

	if err != nil {
		return nil, errors.Wrap(err, "GetAllGamesIdList fail")
	}

	ids := make(map[uint]struct{}, 0)
	for _, info := range gameList {
		ids[info.ID.ID] = struct{}{}
	}
	res := make([]uint, 0)
	for id := range ids {
		res = append(res, id)
	}
	return res, nil
}

func GetGamesByTag(tagIds []uint) (map[uint][]*game_model.GameInfo, error) {
	res := make([]*game_model.GameInfo, 0)
	err := global.App.DB.Debug().Where("tag_id in (?)", tagIds).Find(&res).Error
	if err != nil {
		return nil, errors.Wrap(err, "根据标签获取游戏信息失败")
	}
	mat := make(map[uint][]*game_model.GameInfo)
	for _, re := range res {
		mat[re.TagId] = append(mat[re.TagId], re)
	}
	return mat, nil
}

func SearchGameByName(gameName string) ([]*game_model.GameInfo, error) {
	gameList := make([]*game_model.GameInfo, 0)

	err := global.App.DB.Debug().Model(&game_model.GameInfo{}).Where("name like ? ", "%"+gameName+"%").Find(&gameList).Error

	if err != nil {
		return nil, errors.Wrap(err, "find games by name fail")
	}

	return gameList, nil
}
