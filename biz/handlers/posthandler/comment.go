package posthandler

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/global"
	"github.com/oasis/game_boat/services"
	"github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

func PostComment(ctx *gin.Context) {
	logs := utils.NewLoggerWithXRId(ctx, global.App.Log)
	method := "PostComment"

	req := new(handler_model.PostCommentRequest)
	err := ctx.Bind(req)
	if err != nil {
		logs.Error(fmt.Sprintf("%v get req fail, err = %v", method, err))
		response.ValidateFail(ctx, "get req fail")
		return
	}
	logs.Info(fmt.Sprintf("%v req = %v", method, utils.JsonStrFormatIgnoreErr(req)))

	s := services.NewCommentService(ctx)
	commentDetail, err := s.PostComment(req)
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	resp := &handler_model.PostCommentResponse{Comment: commentDetail}
	logs.Info(fmt.Sprintf("%v success, resp = %v", method, utils.JsonStrFormatIgnoreErr(resp)))
	response.Success(ctx, resp)
}
