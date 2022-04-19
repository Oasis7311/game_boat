package comment

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/global"
	"github.com/oasis/game_boat/services"
	"github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

func GetCommentList(ctx *gin.Context) {
	logs := utils.NewLoggerWithXRId(ctx, global.App.Log)
	method := "GetCommentList"

	req := new(handler_model.ListCommentRequest)
	err := ctx.Bind(req)
	if err != nil {
		logs.Error(fmt.Sprintf("%v get req fail, err = %v", method, req))
		response.ValidateFail(ctx, err.Error())
		return
	}
	logs.Info(fmt.Sprintf("%v req = %v", method, utils.JsonStrFormatIgnoreErr(req)))

	if req.PageInfo == nil {
		req.PageInfo = &handler_model.PageInfo{
			Page:     1,
			PageSize: 10,
		}
	}
	resp := new(handler_model.ListCommentResponse)
	resp.PageInfo = req.PageInfo
	s := services.NewCommentService(ctx)

	resp.Comments, resp.PageInfo.TotalSize, err = s.ListComment(req)
	if err != nil {
		logs.Info(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}
	resp.PageInfo.TotalPage = cast.ToInt(resp.PageInfo.TotalSize) / req.PageInfo.PageSize

	response.Success(ctx, resp)
	logs.Info(fmt.Sprintf("%v success resp = %v", method, utils.JsonStrFormatIgnoreErr(resp)))
}
