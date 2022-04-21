package action

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/global"
	"github.com/oasis/game_boat/services"
	"github.com/oasis/game_boat/utils"
	utils2 "github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

func ContentAction(ctx *gin.Context) {
	logs := utils.NewLoggerWithXRId(ctx, global.App.Log)
	method := "ContentAction"

	req := new(handler_model.ContentActionRequest)
	err := ctx.Bind(req)
	if err != nil {
		logs.Error(fmt.Sprintf("%v get req fail, err = %v", method, err))
		response.ValidateFail(ctx, "get req fail")
		return
	}
	logs.Info(fmt.Sprintf("%v req = %v", method, utils2.JsonStrFormatIgnoreErr(req)))

	s := services.NewActionService()
	err = s.LikeContent(cast.ToUint(ctx.Value("id")), cast.ToInt64(req.ContentId))
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	response.Success(ctx, new(handler_model.ContentActionResponse))
}
