package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/cast"

	"github.com/oasis/game_boat/biz/const_def"
	"github.com/oasis/game_boat/biz/dal/user_dal"
	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/biz/model/user_model"
	"github.com/oasis/game_boat/global"
	utils2 "github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

type RelationHandler struct {
	aUserId         uint
	bUserId         uint
	lastUpdatedTime uint
	lastAUserId     uint
	lastBUserId     uint
	relation        const_def.RelationEnum
	limit           int
}

func (s *RelationHandler) fillUp(aUserId, bUserId, lastUpdateTime, lastAUserId, lastBUserId uint, relation const_def.RelationEnum, limit int) error {
	s.aUserId = aUserId
	s.bUserId = bUserId
	s.lastUpdatedTime = lastUpdateTime
	s.lastAUserId = lastAUserId
	s.lastBUserId = lastBUserId
	s.relation = relation
	s.limit = limit
	if s.limit == 0 {
		s.limit = 10
	}
	if s.relation == 0 {
		return errors.Wrap(errors.New("relation = 0"), "relation should not = 0")
	}
	return nil
}

// GetUserFollowerList 获取用户粉丝列表
func GetUserFollowerList(ctx *gin.Context) {
	logs := utils2.NewLoggerWithXRId(ctx, global.App.Log)
	method := "[GetUserFollowerList]"
	req := new(handler_model.GetUserFollowerListRequest)
	s := new(RelationHandler)

	err := ctx.BindJSON(req)
	if err != nil {
		logs.Error(fmt.Sprintf("%v get req fail err = %v", method, err))
		response.ValidateFail(ctx, err.Error())
		return
	}
	logs.Info(fmt.Sprintf("%v req = %v", method, utils2.JsonStrFormatIgnoreErr(req)))

	err = s.fillUp(0, req.Id, req.LastFollowTime, req.LastFollowerId, 0, const_def.RelationEnumFollow, req.PageSize)
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.ValidateFail(ctx, err.Error())
		return
	}

	aUserIdList, lastFollowTime, err := s.getFollowerList()
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	aUserInfoMap, err := user_dal.GetUserInfoMap(aUserIdList)
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	resp := new(handler_model.GetUserFollowerListResponse)
	resp.LastFollowTime = lastFollowTime
	resp.Followers = make([]*user_model.UserInfo, 0)
	for _, id := range aUserIdList {
		resp.Followers = append(resp.Followers, aUserInfoMap[id])
	}
	response.Success(ctx, resp)
}

func (s *RelationHandler) getFollowerList() ([]uint, uint, error) {
	relationList, err := user_dal.GetBUserRelationList(s.bUserId, s.lastAUserId, s.lastUpdatedTime, s.relation, s.limit)
	if err != nil {
		return nil, 0, err
	}

	res := make([]uint, 0)
	lastTime := uint(0)

	for _, relation := range relationList {
		res = append(res, cast.ToUint(relation.AUserId))
		lastTime = cast.ToUint(relation.UpdatedAt.Unix())
	}

	return res, lastTime, nil
}

// GetUserFollowList 获取用户关注列表
func GetUserFollowList(ctx *gin.Context) {
	logs := utils2.NewLoggerWithXRId(ctx, global.App.Log)
	method := "[GetUserFollowList]"
	req := new(handler_model.GetUserFollowListRequest)
	s := new(RelationHandler)

	err := ctx.BindJSON(req)
	if err != nil {
		logs.Error(fmt.Sprintf("%v get req fail err = %v", method, err))
		response.ValidateFail(ctx, err.Error())
		return
	}
	logs.Info(fmt.Sprintf("%v req = %v", method, utils2.JsonStrFormatIgnoreErr(req)))

	err = s.fillUp(req.Id, 0, req.LastFollowTime, 0, req.LastFollowedUserId, const_def.RelationEnumFollow, req.PageSize)
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.ValidateFail(ctx, err.Error())
		return
	}

	bUserIdList, lastFollowTime, err := s.getFollowList()
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	bUserInfoMap, err := user_dal.GetUserInfoMap(bUserIdList)
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	resp := new(handler_model.GetUserFollowListResponse)
	resp.LastFollowTime = lastFollowTime
	resp.FollowPeople = make([]*user_model.UserInfo, 0)
	for _, id := range bUserIdList {
		resp.FollowPeople = append(resp.FollowPeople, bUserInfoMap[id])
	}
	response.Success(ctx, resp)
}

func (s *RelationHandler) getFollowList() ([]uint, uint, error) {
	relationList, err := user_dal.GetAUserRelationList(s.aUserId, s.lastBUserId, s.lastUpdatedTime, s.relation, s.limit)
	if err != nil {
		return nil, 0, err
	}

	res := make([]uint, 0)
	lastTime := uint(0)

	for _, relation := range relationList {
		res = append(res, cast.ToUint(relation.BUserId))
		lastTime = cast.ToUint(relation.UpdatedAt.Unix())
	}

	return res, lastTime, nil
}

func GetRelationCount(ctx *gin.Context) {
	logs := utils2.NewLoggerWithXRId(ctx, global.App.Log)
	method := "[GetRelationCount]"
	req := new(handler_model.GetRelationCountRequest)
	s := new(RelationHandler)

	err := ctx.BindJSON(req)
	if err != nil {
		logs.Error(fmt.Sprintf("%v get req fail err = %v", method, err))
		response.ValidateFail(ctx, err.Error())
		return
	}
	logs.Info(fmt.Sprintf("%v req = %v", method, utils2.JsonStrFormatIgnoreErr(req)))

	s.aUserId = req.AUserId
	s.bUserId = req.AUserId

	count, err := s.getRelationCount()
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	resp := new(handler_model.GetRelationCountResponse)
	resp.FollowCount = count[0]
	resp.FollowerCount = count[1]
	response.Success(ctx, resp)
}

func (s *RelationHandler) getRelationCount() ([]uint, error) { //返回顺序：关注数、粉丝数
	res := make([]uint, 0)
	followCount, err := user_dal.GetAUserRelationCount(s.aUserId, cast.ToInt(const_def.RelationEnumFollow))
	if err != nil {
		return nil, err
	}
	followerCount, err := user_dal.GetBUserRelationCount(s.bUserId, cast.ToInt(const_def.RelationEnumFollow))
	if err != nil {
		return nil, err
	}

	res = append(res, followCount, followerCount)
	return res, nil
}
