package main

import (
	"github.com/gin-gonic/gin"
	"xybs/controller"
	"xybs/db"
	"xybs/middleware"
)

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

	// 浏览所有文章的基本信息
	r.GET("/Article",controller.ArticleListHandler)

	// 用户登录
	r.POST("/Login",controller.LoginHandler)

	// 用户注册
	r.POST("/Register",controller.RegisterHandler)

	// 浏览一篇文章
	r.GET("/Article/:id",controller.ArticleDetailHandler)

	// 启动http服务，设置在:9090启动，默认在127.0.0.1:8080启动服务
	r.Run("127.0.0.1:9090")
}
