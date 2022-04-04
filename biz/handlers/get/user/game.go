package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	game_dal2 "github.com/oasis/game_boat/biz/dal/mysql/game_dal"
	user_dal2 "github.com/oasis/game_boat/biz/dal/mysql/user_dal"
	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/global"
	utils2 "github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

type GetUserGameHandler struct{}

func GetUserGame(ctx *gin.Context) {
	logs := utils2.NewLoggerWithXRId(ctx, global.App.Log)
	userId := cast.ToUint(ctx.Query("id"))
	commonLog := "[GetUserGame] userId = " + cast.ToString(userId)

	logs.Info(fmt.Sprintf("%v", commonLog))

	collectionGameList, err := user_dal2.GetUserCollectedGameList(userId)
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", commonLog, err))
		response.BusinessFail(ctx, err.Error())
		return
	}
	reservedGameList, err := user_dal2.GetUserReservedGameList(userId)
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", commonLog, err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	gameIds := make([]uint, 0)
	for _, relation := range collectionGameList {
		gameIds = append(gameIds, relation.GameId)
	}
	for _, relation := range reservedGameList {
		gameIds = append(gameIds, relation.GameId)
	}

	gameDetail, err := game_dal2.GetGamesDetail(gameIds)
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", commonLog, err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	resp := handler_model.GetUserGameResponse{
		UserCollectedGame:   make([]*handler_model.UserGame, len(collectionGameList)),
		UserReservedGame:    make([]*handler_model.UserGame, len(reservedGameList)),
		UserCollectdGameMap: make(map[uint]bool),
	}
	for i, relation := range collectionGameList {
		resp.UserCollectedGame[i] = &handler_model.UserGame{
			GameInfo:      gameDetail[relation.GameId],
			CollectedTime: cast.ToUint(relation.CollectedTime.Unix()),
		}
		resp.UserCollectdGameMap[relation.GameId] = true
	}
	for i, relation := range reservedGameList {
		resp.UserReservedGame[i] = &handler_model.UserGame{
			GameInfo:    gameDetail[relation.GameId],
			ReserveTime: cast.ToUint(relation.ReservedTime.Unix()),
		}
	}

	logs.Info(fmt.Sprintf("%v resp = %v", commonLog, utils2.JsonStrFormatIgnoreErr(resp)))
	response.Success(ctx, resp)
}

// GetUserGameCount 获取用户收藏、预约的数量
func GetUserGameCount(ctx *gin.Context) {
	logs := utils2.NewLoggerWithXRId(ctx, global.App.Log)
	userId := cast.ToUint(ctx.Query("id"))
	commonLog := "[GetUserGameCount] userId = " + cast.ToString(userId)

	logs.Info(fmt.Sprintf("%v", commonLog))

	resp := new(handler_model.GetUserGameCountResponse)
	var err error
	resp.CollectedCount, resp.ReservedCount, err = user_dal2.GetUserRelatedGameCount(userId)
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", commonLog, err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	logs.Info(fmt.Sprintf("%v resp = %v", commonLog, utils2.JsonStrFormatIgnoreErr(resp)))
	response.Success(ctx, resp)

}
