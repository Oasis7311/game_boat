package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/oasis/game_boat/biz/const_def"
)

func NewLoggerWithXRId(c *gin.Context, logger *zap.Logger) *zap.Logger {
	newLogger := logger.WithOptions(zap.Fields(zap.String(const_def.XRequestId, c.Request.Header.Get(const_def.XRequestId)))) //添加RequestId
	return newLogger
}

func NewErrorMessage(method string, errorReason string, err error) string {
	return fmt.Sprintf("[%v] %v Error = %v", method, errorReason, err)
}
