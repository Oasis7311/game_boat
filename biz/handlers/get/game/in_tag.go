package game

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/oasis/game_boat/biz/dal/game_dal"
	"github.com/oasis/game_boat/biz/dal/user_dal"
	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/global"
	utils2 "github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

type GetGameInTagHandler struct{}

func GetGameInTag(ctx *gin.Context) {
	logs := utils2.NewLoggerWithXRId(ctx, global.App.Log)
	method := "[GetGameInTag]"

	req := new(handler_model.GetGameInTagRequest)
	err := ctx.Bind(req)
	if err != nil {
		logs.Error(fmt.Sprintf("%v get req fail, err = %v", err))
		response.ValidateFail(ctx, err.Error())
		return
	}
	logs.Info(fmt.Sprintf("%v req = %v", method, utils2.JsonStrFormatIgnoreErr(req)))

	GameList, err := game_dal.GetGameListWithTag(req.TagId, req.Page, req.PageSize)
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.ValidateFail(ctx, err.Error())
		return
	}

	userId := cast.ToUint(ctx.Query("id"))
	logs.Info(fmt.Sprintf("%v get game list success, userId = %v", method, userId))

	resp := new(handler_model.GetGameInTagResponse)
	resp.GameList = GameList
	resp.UserCollectGameMap = make(map[uint]bool)
	resp.UserReserveGameMap = make(map[uint]bool)

	if userId != 0 {
		collectedGameList, _ := user_dal.GetUserCollectedGameList(userId)
		for _, relation := range collectedGameList {
			resp.UserCollectGameMap[relation.GameId] = true
		}

		reservedGameList, _ := user_dal.GetUserReservedGameList(userId)
		for _, relation := range reservedGameList {
			resp.UserReserveGameMap[relation.GameId] = true
		}
	}

	logs.Info(fmt.Sprintf("%v resp = %v", method, utils2.JsonStrFormatIgnoreErr(resp)))
	response.Success(ctx, resp)
}
