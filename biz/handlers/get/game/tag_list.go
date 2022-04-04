package game

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/oasis/game_boat/biz/dal/game_dal"
	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/global"
	utils2 "github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

type GetTagListHandler struct{}

func GetTagList(ctx *gin.Context) {
	logs := utils2.NewLoggerWithXRId(ctx, global.App.Log)
	method := "[GetTagList]"

	logs.Info(fmt.Sprintf("%v req = \"\"", method))

	tagList, err := game_dal.GetTagList()
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	resp := &handler_model.GetTagListResponse{}
	resp.TagList = tagList

	response.Success(ctx, resp)
}
