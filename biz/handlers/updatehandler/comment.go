package updatehandler

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/global"
	"github.com/oasis/game_boat/services"
	"github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

func CommentUpdate(ctx *gin.Context) {
	logs := utils.NewLoggerWithXRId(ctx, global.App.Log)
	method := "CommentUpdate"

	req := new(handler_model.UpdateCommentRequest)
	err := ctx.Bind(req)
	if err != nil {
		logs.Error(fmt.Sprintf("%v get req fail, err = %v", method, err))
		response.ValidateFail(ctx, err.Error())
		return
	}
	logs.Info(fmt.Sprintf("%v req = %v", method, req))

	resp := new(handler_model.UpdateCommentResponse)
	s := services.NewCommentService(ctx)

	err = s.UpdateCommentMsg(req)
	if err != nil {
		logs.Error(fmt.Sprintf("%v update comment fail, err = %+v", err))
		response.BusinessFail(ctx, err.Error())
		return
	}
	logs.Info(fmt.Sprintf("%v success", method))
	response.Success(ctx, resp)
	return
}
