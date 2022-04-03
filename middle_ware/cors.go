package middle_ware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 跨域Header设置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"*"}
	config.AllowCredentials = true
	config.ExposeHeaders = []string{"New-Token", "New-Expires-In", "Content-Disposition"}

	return cors.New(config)
}
