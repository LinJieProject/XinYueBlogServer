package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建一个默认的路由引擎
	r:=gin.Default()

	// 启动http服务，设置在:9090启动，默认在127.0.0.1:8080启动服务
	r.Run("127.0.0.1:9090")
}
