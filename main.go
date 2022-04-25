package main

import (
	"github.com/gin-gonic/gin"

	"github.com/oasis/game_boat/biz/handlers"
	handlerAction "github.com/oasis/game_boat/biz/handlers/action"
	"github.com/oasis/game_boat/biz/handlers/deletehandler"
	"github.com/oasis/game_boat/biz/handlers/gethandler"
	"github.com/oasis/game_boat/biz/handlers/gethandler/comment"
	handlerContent "github.com/oasis/game_boat/biz/handlers/gethandler/content"
	handlerGetGame "github.com/oasis/game_boat/biz/handlers/gethandler/game"
	handlerGetUser "github.com/oasis/game_boat/biz/handlers/gethandler/user"
	"github.com/oasis/game_boat/biz/handlers/posthandler"
	"github.com/oasis/game_boat/biz/handlers/updatehandler"
	handlerUpdateUser "github.com/oasis/game_boat/biz/handlers/updatehandler/user"
	"github.com/oasis/game_boat/global"
	"github.com/oasis/game_boat/initializer"
	"github.com/oasis/game_boat/middle_ware"
	"github.com/oasis/game_boat/services"
	"github.com/oasis/game_boat/storage"
)

func init() {
	// 初始化配置
	initializer.InitializeConfig()

	// 初始化日志设置
	global.App.Log = initializer.InitializeLog()
	global.App.Log.Info("log init success!")

	// 初始化数据库
	global.App.DB = initializer.InitializeDB()
	global.App.Log.Info("db init success!")

	// 初始化redis
	global.App.Redis = initializer.InitializeRedis()
	global.App.Log.Info("redis init success!")

	//初始化三个列表
	storage.InitStorage()
}

func main() {

	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	r := gin.Default()

	register(r)

	r.Run(":" + global.App.Config.App.Port)

}

func register(r *gin.Engine) {
	r.Use(middle_ware.CustomRecovery(), //错误恢复堆栈现场
		middle_ware.Cors(),           //跨域设置
		middle_ware.CheckRequestId()) //填充请求Id

	r.GET("/ping", handlers.Ping)

	{ //获取数据相关接口
		rGet := r.Group("/get")
		rGet.POST("/comment/list", comment.GetCommentList)
		{
			user := rGet.Group("/user")                                   //用户
			user.POST("/followers", handlerGetUser.GetUserFollowerList)   //粉丝列表
			user.POST("/follows", handlerGetUser.GetUserFollowList)       //关注列表
			user.POST("/relation_count", handlerGetUser.GetRelationCount) //关系数量
			user.GET("/game", handlerGetUser.GetUserGame)                 //游戏
			user.GET("/game_count", handlerGetUser.GetUserGameCount)      //收藏预约游戏数量
			user.POST("/moment", handlerGetUser.GetUserMoment)            //获取动态列表
			user.GET("/:email", handlerGetUser.GetUser)
		}
		{
			game := rGet.Group("/game")                       //游戏
			game.GET("/tag_list", handlerGetGame.GetTagList)  //标签列表
			game.POST("/in_tag", handlerGetGame.GetGameInTag) //标签下游戏列表
			game.GET("/three_list", handlerGetGame.GetGameThreeList)
		}
		{
			article := rGet.Group("/content")                        //文章
			article.GET("/detail", handlerContent.GetContentDetail)  //详情
			article.POST("/follow", handlerContent.GetFollowContent) //关注列表
		}
		{
			rGet.POST("/community", gethandler.ListCommunity) //社区页
			rGet.POST("/main_page", gethandler.ListMainPage)  //首页
		}
	}

	{ //数据写入相关接口
		{
			rPost := r.Group("/post") //发布
			rPost.Use(middle_ware.JWTAuth(services.AppGuardName))
			rPost.POST("/content", posthandler.PostContent) //文章
			rPost.POST("/comment", posthandler.PostComment) //评论
		}
		{
			rDelete := r.Group("/delete") //删除
			rDelete.Use(middle_ware.JWTAuth(services.AppGuardName))
			rDelete.POST("/comment", deletehandler.DeleteComment) //评论 or 回复
		}
		{
			rAction := r.Group("/action") //交互行为
			rAction.Use(middle_ware.JWTAuth(services.AppGuardName))
			rAction.POST("/game", handlerAction.GameAction)       //关注、取消关注游戏
			rAction.POST("/content", handlerAction.ContentAction) //点赞、取消赞文章
		}
		{
			rUpdate := r.Group("/update") //更新
			rUpdate.Use(middle_ware.JWTAuth(services.AppGuardName))
			{
				rUpdateUser := rUpdate.Group("/user")                         //用户
				rUpdateUser.POST("/relation", handlerUpdateUser.PeopleRelate) //更新用户关系
				rUpdateUser.POST("/info", handlerUpdateUser.UpdateUserInfo)   //信息
			}
			rUpdate.POST("comment", updatehandler.CommentUpdate) //更新评论内容
		}
	}

	{
		r.POST("/register", handlers.Register)
		r.POST("/login", handlers.Login)
		r.POST("/logout", middle_ware.JWTAuth(services.AppGuardName), middle_ware.UserIdAuth(), handlers.Logout)
		r.GET("/token_check", middle_ware.JWTAuth(services.AppGuardName), handlers.CheckToken)
		r.POST("/search", handlers.Search)
	}
}
