package get

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/oasis/game_boat/biz/dal/mysql/game_dal"
	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/global"
	utils2 "github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

func ListMainPage(ctx *gin.Context) {
	logs := utils2.NewLoggerWithXRId(ctx, global.App.Log)
	method := "[ListMainPage]"
	req := new(handler_model.GetMainPageRequest)

	err := ctx.Bind(req)
	if err != nil {
		logs.Error(fmt.Sprintf("%v get req fail, err = %v", method, err))
		response.ValidateFail(ctx, err.Error())
		return
	}
	logs.Info(fmt.Sprintf("%v req = %v", method, utils2.JsonStrFormatIgnoreErr(req)))

	resp := new(handler_model.GetMainPageResponse)

	gameIds, err := game_dal.GetAllGamesIdList()
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	oldSlice := make([]interface{}, 0)
	for _, id := range gameIds {
		oldSlice = append(oldSlice, id)
	}
	newSlice := utils2.FakeShuffleNumSlice(oldSlice, req.Timestamp)

	needGameIds := make([]uint, 0)
	for i := (req.Page-1)*req.PageSize + 1; i <= req.Page*req.PageSize; i++ {
		if int(i) >= len(gameIds) {
			break
		}
		needGameIds = append(needGameIds, cast.ToUint(newSlice[i]))
	}

	gameInfoMap, err := game_dal.GetGamesDetail(needGameIds)
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	for _, id := range needGameIds {
		resp.GameList = append(resp.GameList, gameInfoMap[id])
	}
	resp.Timestamp = req.Timestamp

	response.Success(ctx, resp)
	logs.Info(fmt.Sprintf("%v success, gameIdList = %v, page = %v, pageSize = %v", method, needGameIds, req.Page, req.PageSize))
}
