package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()                    //创建一个默认的路由引擎
	r.GET("/ping", func(c *gin.Context) { //创建测试的路由
		//GET:请求方式    /ping:请求的路径
		//当客户端以GET方法请求/ping路径时，会执行后面的匿名函数
		//c.JSON:返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//启动HTTP服务，默认在0.0.0.0:8200启动服务
	r.Run(":8200")
}
