package middle_ware

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/cast"

	"github.com/oasis/game_boat/biz/const_def"
	"github.com/oasis/game_boat/global"
	"github.com/oasis/game_boat/services"
	"github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		logs := utils.NewLoggerWithXRId(c, global.App.Log)

		//tokenStr := c.Query(const_def.XPolisToken)
		tokenStr := c.Request.Header.Get(const_def.XPolisToken)
		if tokenStr == "" {
			response.TokenFail(c)
			logs.Error(utils.NewErrorMessage("JWTAuth", "tokenStr is empty", errors.New("empty token str")))
			c.Abort()
			return
		}
		//tokenStr = tokenStr[len(services.TokenType)+1:]

		logs.Info(fmt.Sprintf("tokenStr = %v", tokenStr))

		// Token 解析校验
		token, err := jwt.ParseWithClaims(tokenStr, &services.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Secret), nil
		})
		if err != nil || services.JwtService.IsInBlacklist(tokenStr) {
			response.TokenFail(c)
			c.Abort()
			return
		}

		claims := token.Claims.(*services.CustomClaims)

		if claims.Issuer != GuardName {
			logs.Error(utils.NewErrorMessage("JWTAuth", "发布者校验失败, claims.Issuer = "+claims.Issuer+", GuaedName = "+GuardName, errors.New("发布者校验失败")))
			response.TokenFail(c)
			c.Abort()
			return
		}

		// 校验token是否有限
		if !token.Valid {
			logs.Info(fmt.Sprintf("[JWTAuth] Token无效，id = %v", claims.Id))
			response.TokenFail(c)
			c.Abort()
			return
		}

		// token 续签
		if claims.ExpiresAt-time.Now().Unix() < global.App.Config.Jwt.RefreshGracePeriod {
			lock := global.Lock("refresh_token_lock", global.App.Config.Jwt.JwtBlacklistGracePeriod)
			if lock.Get() {
				err, user := services.JwtService.GetUserInfo(GuardName, claims.Id)
				if err != nil {
					global.App.Log.Error(err.Error())
					lock.Release()
				} else {
					tokenData, _, _ := services.JwtService.CreateToken(GuardName, user)
					c.Header(const_def.XPolisToken, tokenData.AccessToken)
					c.Header(const_def.XTokenExpireAt, time.Now().Add(time.Duration(global.App.Config.Jwt.JwtTtl)*time.Second).Format("2006-01-02 15:04:05"))
					_ = services.JwtService.JoinBlackList(token)
				}
			}
		}

		c.Set("token", token)
		c.Set("id", claims.Id)
		c.Next()
	}
}

func UserIdAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logs := utils.NewLoggerWithXRId(ctx, global.App.Log)
		if cast.ToUint(ctx.Query("id")) != cast.ToUint(ctx.Value("id")) {
			logs.Error(fmt.Sprintf("[UserIdAuth] id not match, queryId = %v, ctxId = %v", ctx.Query("id"), ctx.Value("id")))
			response.ValidateFail(ctx, "id not match")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
