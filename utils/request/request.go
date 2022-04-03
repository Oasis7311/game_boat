package request

import (
	"github.com/gin-gonic/gin"
)

func BindJsonRequestBody(ctx *gin.Context, req interface{}) (interface{}, error) {
	err := ctx.BindJSON(req)
	return req, err
}
