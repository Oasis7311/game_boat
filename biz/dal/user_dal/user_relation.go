package user_dal

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cast"

	"github.com/oasis/game_boat/biz/const_def"
	"github.com/oasis/game_boat/biz/model/user_model"
	"github.com/oasis/game_boat/global"
)

// GetAUserRelationList 获取A用户发起的关系列表
func GetAUserRelationList(aUserId uint, lastBUserId uint, lastUpdateTime uint, relationEnum const_def.RelationEnum, limit int) ([]*user_model.UserRelation, error) {
	userRelationList := make([]*user_model.UserRelation, 0)

	res := global.App.DB.Debug().
		Where("a_user_id = ? and relation = ?", aUserId, relationEnum).
		Where("updated_at > ? and b_user_id > ?", lastUpdateTime, lastBUserId).
		Limit(limit).Find(&userRelationList)
	return userRelationList, errors.Wrap(res.Error, "GetAUserRelationList fail")
}

// GetBUserRelationList B用户被发起的关系列表
func GetBUserRelationList(bUserId uint, lastAUserId uint, lastUpdateTime uint, relationEnum const_def.RelationEnum, limit int) ([]*user_model.UserRelation, error) {
	userRelationList := make([]*user_model.UserRelation, 0)

	res := global.App.DB.Debug().
		Where("b_user_id = ? and relation = ?", bUserId, relationEnum).
		Where("updated_at > ? and a_user_id > ?", lastUpdateTime, lastAUserId).
		Limit(limit).Find(&userRelationList)
	return userRelationList, errors.Wrap(res.Error, "GetBUserRelationList fail")
}

// GetAUserRelationCount 获取A用户relation数量
func GetAUserRelationCount(aUserId uint, relation const_def.RelationEnum) (uint, error) {
	count := int64(0)
	err := global.App.DB.Debug().Model(&user_model.UserRelation{}).Where("a_user_id = ? and relation = ?", aUserId, relation).Count(&count).Error
	return cast.ToUint(count), errors.Wrap(err, "GetAUserRelationCount fail")
}

func GetBUserRelationCount(bUserId uint, relation const_def.RelationEnum) (uint, error) {
	count := int64(0)
	err := global.App.DB.Debug().Model(&user_model.UserRelation{}).Where("b_user_id = ? and relation = ?", bUserId, relation).Count(&count).Error
	return cast.ToUint(count), errors.Wrap(err, "GetBUserRelationCount fail")
}

func GetABUserRelationCount(aUserId uint, bUserId uint, relation const_def.RelationEnum) (int64, error) {
	count := int64(0)
	err := global.App.DB.Debug().Model(&user_model.UserRelation{}).Where("a_user_id = ? and b_user_id = ? and relation = ?", aUserId, bUserId, relation).Count(&count).Error
	return count, errors.Wrap(err, "GetABUserRelationCount fail")
}

// CreateAUserRelation 创建用户关系记录
func CreateAUserRelation(aUserId, bUserid uint, relation const_def.RelationEnum) error {
	userRelation := &user_model.UserRelation{
		AUserId:  aUserId,
		BUserId:  bUserid,
		Relation: relation,
	}

	count, err := GetABUserRelationCount(aUserId, bUserid, relation)
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.Wrap(errors.New(fmt.Sprintf("%v %v %v count != 0", aUserId, const_def.RelationEnumEngMap[relation], bUserid)), "")
	}
	return errors.Wrap(global.App.DB.Debug().Create(userRelation).Error, const_def.RelationEnumStrMap[relation]+" 失败")
}

// DeleteAUserRelation 删除用户关系记录
func DeleteAUserRelation(aUserId, bUserId uint, relation const_def.RelationEnum) error {
	return errors.Wrap(global.App.DB.Debug().
		Where("a_user_id = ? and b_user_id = ? and relation = ?", aUserId, bUserId, relation).
		Delete(&user_model.UserRelation{}).Error, const_def.RelationEnumStrMap[relation]+" 失败")
}
