package router

import (
	v1 "gin-base/api/v1/controller"
	"gin-base/tool"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable" // 改包可完美解决问题
)

func InitRouter() *gin.Engine {

	// 启用gin的日志输出带颜色
	gin.ForceConsoleColor()
	// 替换默认Writer（关键步骤）
	gin.DefaultWriter = colorable.NewColorableStdout()
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(tool.GetConfig().Server.ServerMode)

	new(v1.HelloController).Router(r)
	new(v1.UserController).Router(r)

	// apiv1 := r.Group("/api/v1")
	// {

	// 	auth := apiv1.Group("/hello")
	// 	{
	// 		auth.POST("/login", v1.HelloController)
	// 		// auth.GET("/refresh_token", v1.RefreshToken)
	// 	}

	// 	// user := apiv1.Group("/user")
	// 	// user.Use(jwt.JWT())
	// 	{
	// 		// user.GET("/info", v1.UserInfo)

	// 	}
	// 	// albums := apiv1.Group("/album")
	// 	// albums.Use(jwt.JWT())
	// 	{
	// 		// albums.GET("/list", v1.GetAlbums)
	// 		// albums.POST("/create", v1.AddAlbum)
	// 		// albums.POST("/edit", v1.EditAlbum)
	// 		// albums.GET("/get", v1.GetAlbum)
	// 	}

	// 	// article := apiv1.Group("/article")
	// 	// article.Use(jwt.JWT())
	// 	{
	// 		// article.POST("/create", v1.AddArticle)s

	// 	}
	// }

	return r
}
