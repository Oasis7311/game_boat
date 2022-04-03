package middle_ware

import (
	"github.com/gin-gonic/gin"

	"github.com/oasis/game_boat/biz/const_def"
	utils2 "github.com/oasis/game_boat/utils"
)

func CheckRequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		rId := c.Request.Header.Get(const_def.XRequestId)

		if rId == "" {
			rId = utils2.GenXid()
		}

		c.Request.Header.Set(const_def.XRequestId, rId)
		c.Header(const_def.XRequestId, rId)
	}
}
