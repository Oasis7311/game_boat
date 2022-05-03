package recommend_dal

import (
	"github.com/pkg/errors"

	"github.com/oasis/game_boat/biz/model/recommend_model"
	"github.com/oasis/game_boat/global"
)

func GetAllRecommendData(userId uint) ([]*recommend_model.RecommendData, error) {
	res := make([]*recommend_model.RecommendData, 0)
	err := global.App.DB.Where("is_deleted = 0").Find(&res).Error
	return res, errors.Wrap(err, "获取用户推荐源数据失败")
}
