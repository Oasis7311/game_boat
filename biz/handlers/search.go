package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/global"
	"github.com/oasis/game_boat/services"
	"github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

func Search(ctx *gin.Context) {
	logs := utils.NewLoggerWithXRId(ctx, global.App.Log)
	method := "Search"

	req := new(handler_model.SearchRequest)
	err := ctx.Bind(req)
	if err != nil {
		logs.Error(fmt.Sprintf("%v get req fail err = %v", method, err))
		response.ValidateFail(ctx, err.Error())
		return
	}
	logs.Info(fmt.Sprintf("%v req = %v", method, utils.JsonStrFormatIgnoreErr(req)))

	s := &services.SearchService{}
	gameList, err := s.SearchGame(req.Name)
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	resp := new(handler_model.SearchResponse)
	resp.Games = gameList
	response.Success(ctx, resp)
	logs.Info(fmt.Sprintf("%v success", method))
}
