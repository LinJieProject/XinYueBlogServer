package main

import (
	"github.com/gin-gonic/gin"
	"xybs/controller"
	"xybs/db"
	"xybs/middleware"

	_ "xybs/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title 心阅博客
// @version 0.0.1
// @description 给心阅博客网站前端提供API
// @BasePath /api/v1/
func main() {
	// 连接数据库
	err := db.InitDB()
	if err != nil {
		panic(err)
	}

	// 创建一个默认的路由引擎
	r:=gin.Default()

	// 使用中间件
	r.Use(middleware.CorsMiddleware())

	// 创建路由组
	v1 := r.Group("/api/v1")
	{
		// 浏览所有文章的基本信息
		v1.GET("/Article",controller.ArticleListHandler)

		// 用户登录
		v1.POST("/Login",controller.LoginHandler)

		// 用户注册
		v1.POST("/Register",controller.RegisterHandler)

		// 浏览一篇文章
		v1.GET("/Article/:id",controller.ArticleDetailHandler)

		// 发布一篇文章
		v1.POST("/PublishArticle",controller.PublishArticleHandler)

		// 查询评论
		v1.GET("/Comment/:ArticleID",controller.QueryCommentHandler)

		// 发布一条评论
		v1.POST("/PublishComment",controller.PublishCommentHandler)
	}

	// 文档界面访问URL
	// http://127.0.0.1:9090/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 启动http服务，设置在:9090启动，默认在127.0.0.1:8080启动服务
	r.Run(":9090")
}
