package gethandler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/oasis/game_boat/biz/dal/mysql/content_dal"
	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/global"
	"github.com/oasis/game_boat/services"
	utils2 "github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

func ListCommunity(ctx *gin.Context) {
	logs := utils2.NewLoggerWithXRId(ctx, global.App.Log)
	method := "[ListMainPage]"
	req := new(handler_model.GetCommunityRequest)
	s := services.GetContentDetailHandler()

	err := ctx.Bind(req)
	if err != nil {
		logs.Error(fmt.Sprintf("%v get req fail, err = %v", method, err))
		response.ValidateFail(ctx, err.Error())
		return
	}
	logs.Info(fmt.Sprintf("%v req = %v", method, utils2.JsonStrFormatIgnoreErr(req)))

	resp := new(handler_model.GetCommunityResponse)

	contentIds, err := content_dal.GetAllContentId()
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}
	logs.Info(fmt.Sprintf("%v GetAllContentId success, lend = %v", method, len(contentIds)))

	oldSlice := make([]interface{}, 0)
	for _, id := range contentIds {
		oldSlice = append(oldSlice, id)
	}
	newSlice := utils2.FakeShuffleNumSlice(oldSlice, req.Timestamp)

	needContentIds := make([]int64, 0)
	for i := (req.Page-1)*req.PageSize + 1; i <= req.Page*req.PageSize; i++ {
		if int(i) >= len(contentIds) {
			break
		}
		needContentIds = append(needContentIds, cast.ToInt64(newSlice[i]))
	}

	s.ContentMap, err = content_dal.GetContentRaw(needContentIds)
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}
	logs.Info(fmt.Sprintf("%v get raw content success", method))

	err = s.GetContentDetail()
	if err != nil {
		logs.Error(fmt.Sprintf("%v %v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}
	for i, _ := range s.ContentDetailMap {
		s.ContentDetailMap[i].Text = ""
	}

	for _, id := range needContentIds {
		resp.ContentList = append(resp.ContentList, s.ContentDetailMap[id])
	}
	resp.Timestamp = req.Timestamp

	response.Success(ctx, resp)
	logs.Info(fmt.Sprintf("%v success, contentIdList = %v, page = %v, pageSize = %v", method, needContentIds, req.Page, req.PageSize))

}
