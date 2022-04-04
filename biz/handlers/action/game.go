package action

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/oasis/game_boat/biz/const_def"
	user_dal2 "github.com/oasis/game_boat/biz/dal/mysql/user_dal"
	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/global"
	utils2 "github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

func GameAction(ctx *gin.Context) {
	logs := utils2.NewLoggerWithXRId(ctx, global.App.Log)
	method := "[GameAction]"

	userId := cast.ToUint(ctx.Value("id"))
	req := &handler_model.GameActionRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		logs.Error(fmt.Sprintf("%v get req fail, err = %v", method, err))
		response.ValidateFail(ctx, "get req fail")
		return
	}
	logs.Info(fmt.Sprintf("%v req = %v", method, utils2.JsonStrFormatIgnoreErr(req)))

	key, timeKey := "is_collected", "collected_time"
	value, timeValue := false, time.Now()
	if req.Action > 0 {
		value = true
		if req.Action == const_def.ActionEnumReserveGame {
			key, timeKey = "is_reserved", "reserved_time"
		}
	} else if req.Action == const_def.ActionEnumCancelReserveGame {
		key, timeKey = "is_reserved", "reserved_time"
	}

	err = user_dal2.UpdateUserGameRelation(userId, req.GameId, key, timeKey, value, timeValue)
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	logs.Info(fmt.Sprintf("%v success", method))
	response.Success(ctx, new(handler_model.GameActionResponse))
	return
}
