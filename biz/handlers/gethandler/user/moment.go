package user

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/oasis/game_boat/biz/dal/mysql/content_dal"
	"github.com/oasis/game_boat/biz/model/content_model"
	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/global"
	"github.com/oasis/game_boat/services"
	utils2 "github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

func GetUserMoment(ctx *gin.Context) {
	logs := utils2.NewLoggerWithXRId(ctx, global.App.Log)
	method := "GetUserMoment"
	req := new(handler_model.GetUserMomentRequest)

	err := ctx.Bind(req)
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}
	logs.Info(fmt.Sprintf("%v req = %v", method, utils2.JsonStrFormatIgnoreErr(req)))

	resp := new(handler_model.GetUserMomentResponse)

	if req.Content {
		resp.Contents = make([]*content_model.ContentDetail, 0)
		contents, err := content_dal.GetContentListByUserId(req.UserId)
		if err != nil {
			logs.Error(fmt.Sprintf("%v %+v", method, err))
		} else {
			s := services.GetContentDetailHandler()
			for _, content := range contents {
				s.ContentMap[content.ContentId] = content
			}
			err = s.GetContentDetail()
			if err != nil {
				logs.Error(fmt.Sprintf("%v %+v", method, err))
			} else {
				for i, content := range contents {
					resp.Contents = append(resp.Contents, s.ContentDetailMap[content.ContentId])
					resp.Contents[i].Text = ""
				}
			}
		}
	}
	if req.Review {

	}
	response.Success(ctx, resp)
	logs.Info(fmt.Sprintf("%v success, userId = %v, contentlen = %v", method, req.UserId, len(resp.Contents)))
}
