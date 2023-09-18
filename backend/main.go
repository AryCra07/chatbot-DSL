package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	ginServer := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	ginServer.GET("/hello", func(context *gin.Context) {
		// c.JSON：返回JSON格式的数据
		context.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	err := ginServer.Run(":8082")
	if err != nil {
		return
	}
}
