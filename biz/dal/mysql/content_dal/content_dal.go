package content_dal

import (
	"github.com/pkg/errors"

	"github.com/oasis/game_boat/biz/model/content_model"
	"github.com/oasis/game_boat/global"
)

// GetContentRaw 根据内容Id获取内容详情
func GetContentRaw(contentIds []int64) (map[int64]*content_model.Content, error) {
	contents := make([]*content_model.Content, 0)

	err := global.App.DB.Debug().Where("content_id in (?)", contentIds).Find(&contents).Error

	res := make(map[int64]*content_model.Content)
	for _, content := range contents {
		res[content.ContentId] = content
	}
	return res, errors.Wrap(err, "GetContentDetail Fail")
}

// GetAllContentId 获取所有内容Id
func GetAllContentId() ([]int64, error) {
	contents := make([]*content_model.Content, 0)

	err := global.App.DB.Debug().Select("content_id").Find(&contents).Error

	ids := make([]int64, 0)
	for _, content := range contents {
		ids = append(ids, content.ContentId)
	}

	return ids, errors.Wrap(err, "GetAllContentId Fail")
}

// CreateContent 创建文章
func CreateContent(content *content_model.Content) error {
	err := global.App.DB.Debug().Create(content).Error
	return errors.Wrap(err, "CreateContent fail")
}
