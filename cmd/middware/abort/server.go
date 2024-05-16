package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.New()

	v1 := r.Group("/v1")

	// 注册中间件
	v1.Use(middleware("one")) // 注意中间件注册的顺序
	v1.Use(middleware("two"))
	v1.Use(middleware("three"))
	{
		v1.GET("/handler", handler)
	}

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}

func middleware(msg string) func(*gin.Context) {
	return func(c *gin.Context) {
		// 请求进入中间件
		fmt.Printf("middleware %s befer\n", msg)

		// 如果请求头中包含 abort: true 则在第三个中间件中断请求
		if c.GetHeader("abort") == "true" && msg == "three" {
			c.Abort()
			c.JSON(http.StatusBadRequest, gin.H{"hasAbortHeader": true})
			return
		}
		// 在此处中断此middleware的执行，将请求的控制权转交给下一个middleware
		c.Next()

		// 请求结束，返回响应前执行
		fmt.Printf("middlewareOne %s after\n", msg)
	}
}

func handler(c *gin.Context) {
	fmt.Println("in handler")

	c.JSON(http.StatusOK, gin.H{"message": "handler success"})
}
