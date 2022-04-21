package content

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/oasis/game_boat/biz/dal/mysql/content_dal"
	"github.com/oasis/game_boat/biz/model/content_model"
	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/global"
	"github.com/oasis/game_boat/services"
	"github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

func GetFollowContent(ctx *gin.Context) {
	logs := utils.NewLoggerWithXRId(ctx, global.App.Log)
	method := "GetFollowContent"

	req := new(handler_model.GetFollowUserContentRequest)
	err := ctx.Bind(req)
	if err != nil {
		logs.Error(fmt.Sprintf("%v get req fail, err = %v", method, err))
		response.ValidateFail(ctx, err.Error())
		return
	}
	logs.Info(fmt.Sprintf("%v req = %v", method, utils.JsonStrFormatIgnoreErr(req)))

	contentIds := make([]interface{}, 0)
	for _, userId := range req.UserId {
		contents, _ := content_dal.GetContentListByUserId(userId)
		for _, content := range contents {
			contentIds = append(contentIds, content.ContentId)
		}
	}
	newRawIds := utils.FakeShuffleNumSlice(contentIds, time.Now().Unix())
	newIds := make([]int64, 0)
	for _, id := range newRawIds {
		newIds = append(newIds, cast.ToInt64(id))
		if len(newIds) == 10 {
			break
		}
	}

	s := services.ContentDetailHandler{}
	s.ContentMap, err = content_dal.GetContentRaw(newIds)
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	resp := new(handler_model.GetFollowUserContentResponse)
	err = s.GetContentDetail()
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	resp.Content = make([]*content_model.ContentDetail, 0)
	for _, detail := range s.ContentDetailMap {
		resp.Content = append(resp.Content, detail)
	}
	response.Success(ctx, resp)
	logs.Info(fmt.Sprintf("%v success", method))
}
