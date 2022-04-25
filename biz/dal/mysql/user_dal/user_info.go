package user_dal

import (
	"github.com/pkg/errors"
	"github.com/spf13/cast"

	"github.com/oasis/game_boat/biz/model/content_model"
	"github.com/oasis/game_boat/biz/model/user_model"
	"github.com/oasis/game_boat/global"
	"github.com/oasis/game_boat/utils"
)

// CreateUserLoginInfo 创建用户账号记录
func CreateUserLoginInfo(user *user_model.UserLoginInfo) (int64, uint, error) {
	res := global.App.DB.Debug().Where("email = ?", user.Email).Select("id").First(&user_model.UserLoginInfo{})

	if res.RowsAffected != 0 {
		return 0, 0, errors.New("邮箱已注册")
	}

	user.Password = utils.BcryptMake([]byte(user.Password)) //加密用户密码

	res = global.App.DB.Debug().Model(&user_model.UserLoginInfo{}).Create(user)

	return res.RowsAffected, user.ID.ID, errors.Wrap(res.Error, "创建用户账号记录失败")
}

// CreateUserInfo 创建用户记录
func CreateUserInfo(user *user_model.UserInfo) (int64, error) {
	res := global.App.DB.Debug().Model(&user_model.UserInfo{}).Create(user)

	return res.RowsAffected, errors.Wrap(res.Error, "创建用户记录失败")
}

// GetUserLoginInfoByEmail 根据用户邮箱获取用户登陆信息
func GetUserLoginInfoByEmail(email string) (*user_model.UserLoginInfo, error) {
	user := &user_model.UserLoginInfo{}
	res := global.App.DB.Debug().Where("email = ?", email).First(user)
	return user, errors.Wrap(res.Error, "查找用户登陆记录失败")
}

// GetUserLoginInfoByLoginId 根据LoginId获取用户登陆信息
func GetUserLoginInfoByLoginId(id uint) (*user_model.UserLoginInfo, error) {
	user := &user_model.UserLoginInfo{}
	res := global.App.DB.Debug().Where("id = ?", id).First(user)
	return user, errors.Wrap(res.Error, "查找用户登陆信息失败")
}

// GetUserInfo 根据用户Id获取用户信息
func GetUserInfo(id uint) (*user_model.UserInfo, error) {
	user := &user_model.UserInfo{}
	res := global.App.DB.Debug().Where("id = ?", id).First(user)
	return user, errors.Wrap(res.Error, "获取用户信息失败")
}

// GetUserInfoMap 获取多个用户信息
func GetUserInfoMap(id []uint) (map[uint]*user_model.UserInfo, error) {
	userList := make([]*user_model.UserInfo, 0)
	err := global.App.DB.Debug().Where("id in (?)", id).Find(&userList).Error
	if err != nil {
		return map[uint]*user_model.UserInfo{}, errors.Wrap(err, "获取用户信息失败")
	}

	res := make(map[uint]*user_model.UserInfo)
	for _, info := range userList {
		res[cast.ToUint(info.ID.ID)] = &user_model.UserInfo{
			ID:           info.ID,
			UserBaseInfo: info.UserBaseInfo,
		}
	}
	return res, nil
}

// UpdateUserInfo 更新用户信息
func UpdateUserInfo(userInfo *user_model.UserInfo) error {
	err := global.App.DB.Debug().Where("id = ?", userInfo.ID.ID).Updates(userInfo).Error
	return errors.Wrap(err, "更新用户信息失败")
}

// UpdateUserLoginInfo 更新用户登陆信息
func UpdateUserLoginInfo(userLoginInfo *user_model.UserLoginInfo) error {
	return errors.Wrap(global.App.DB.Debug().Where("id = ?", userLoginInfo.ID.ID).Updates(userLoginInfo).Error, "更新用户登陆信息失败")
}

// GetUserInfoByContent 根据内容获取用户信息
func GetUserInfoByContent(contents map[int64]*content_model.Content) (map[int64]*user_model.UserInfo, error) {
	authorIds := make([]uint, 0)
	for _, content := range contents {
		authorIds = append(authorIds, content.AuthorId)
	}

	authorInfoMap, err := GetUserInfoMap(authorIds)
	if err != nil {
		return nil, errors.Wrap(err, "GetUserInfoByContent -> GetUserInfoMap Fail")
	}

	res := make(map[int64]*user_model.UserInfo)
	for _, content := range contents {
		res[content.ContentId] = authorInfoMap[content.AuthorId]
	}

	return res, nil
}

func GetUserInfoByEmail(email string) (*user_model.UserInfo, error) {
	res := new(user_model.UserInfo)
	err := global.App.DB.Debug().Where("email = ? ", email).First(res).Error
	return res, errors.Wrap(err, "find user by email fail")

}
