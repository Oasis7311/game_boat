package user

import (
	"fmt"
	"strconv"

	user_dal2 "github.com/oasis/game_boat/biz/dal/mysql/user_dal"
	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/biz/model/user_model"
	"github.com/oasis/game_boat/global"
	"github.com/oasis/game_boat/utils"

	"github.com/pkg/errors"
)

type userService struct {
}

var UserService = new(userService)

func (s *userService) Register(userLoginInfo *user_model.UserLoginInfo, userBaseInfo *user_model.UserBaseInfo) (error, *user_model.UserInfo) {
	rowEffect, userId, err := user_dal2.CreateUserLoginInfo(userLoginInfo)
	if err != nil {
		return err, nil
	}
	if rowEffect == 0 || userId == 0 {
		return errors.New(fmt.Sprintf("创建用户失败, rowEffect = %v, userLoginId = %v", rowEffect, userId)), nil
	}

	userInfo := &user_model.UserInfo{
		UserBaseInfo: *userBaseInfo,
		LoginInfoId:  userId,
	}
	rowEffect, err = user_dal2.CreateUserInfo(userInfo)
	if err != nil {
		return err, nil
	}
	if rowEffect == 0 {
		return errors.New(fmt.Sprintf("创建用户失败, rowEffect = %v", rowEffect)), nil
	}

	return nil, userInfo
}

func (s *userService) Login(request *handler_model.LoginRequest) (*user_model.UserInfo, error) {
	user, err := user_dal2.GetUserLoginInfoByEmail(request.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.Wrap(errors.New("record not found"), "查找用户记录失败")
	}

	if !utils.BcryptMakeCheck([]byte(request.Password), user.Password) {
		return nil, errors.Wrap(errors.New("password [BcryptMakeCheck] fail"), "密码错误")
	}

	return user_dal2.GetUserInfo(user.ID.ID)
}

// GetUserInfo 获取用户信息
func (userService *userService) GetUserInfo(id string) (err error, user *user_model.UserInfo) {
	intId, err := strconv.Atoi(id)
	err = global.App.DB.Where("id = ?", intId).First(user).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}
