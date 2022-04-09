package posthandler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/oasis/game_boat/biz/dal/mysql/content_dal"
	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/global"
	"github.com/oasis/game_boat/services"
	utils2 "github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

func PostContent(ctx *gin.Context) {
	logs := utils2.NewLoggerWithXRId(ctx, global.App.Log)
	method := "[PostContent]"
	req := new(handler_model.PostContentRequest)
	s := services.GetContentDetailHandler()

	err := ctx.Bind(req)
	if err != nil {
		logs.Error(fmt.Sprintf("%v get req fail, err = %v", method, err))
		response.ValidateFail(ctx, err.Error())
		return
	}

	logs.Info(fmt.Sprintf("%v req = %v", method, utils2.JsonStrFormatIgnoreErr(req)))

	contentId := utils2.GenI64Id()
	s.ContentDetailMap[contentId] = req.Content
	s.GetContentByContentDetail()

	if content, ok := s.ContentMap[contentId]; ok {
		err = content_dal.CreateContent(content)
	} else {
		err = errors.New("Get Content Fail")
	}
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}
	resp := new(handler_model.PostContentResponse)
	s.GetContentDetail()
	if contentDetail, ok := s.ContentDetailMap[contentId]; ok {
		resp.Content = contentDetail
	}
	response.Success(ctx, resp)
	return
}
