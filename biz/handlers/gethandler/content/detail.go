package content

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

func GetContentDetail(ctx *gin.Context) {
	logs := utils2.NewLoggerWithXRId(ctx, global.App.Log)
	method := "[GetContentDetail]"
	req := new(handler_model.GetContentRequest)
	c := services.GetContentDetailHandler()

	contentId, ok := ctx.GetQuery("content_id")
	if !ok {
		logs.Error(fmt.Sprintf("%v get query fail", method))
		response.BusinessFail(ctx, "get content_id fail")
		return
	}
	req.ContentId = cast.ToInt64(contentId)
	logs.Info(fmt.Sprintf("%v req = %v", utils2.JsonStrFormatIgnoreErr(req)))

	var err error
	c.ContentMap, err = content_dal.GetContentRaw([]int64{req.ContentId})
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	resp := new(handler_model.GetContentResponse)

	err = c.GetContentDetail()
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}
	for i, _ := range c.ContentDetailMap {
		c.ContentDetailMap[i].Summary = ""
	}

	resp.Content = c.ContentDetailMap[req.ContentId]

	response.Success(ctx, resp)
}
