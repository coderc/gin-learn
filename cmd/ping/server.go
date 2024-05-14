package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 实例化一个gin.Engine
	r := gin.New()

	// 使用gin.Engine为/ping注册一个GET api
	r.GET("/ping", func(c *gin.Context) {

		//
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 默认监听 8080 端口
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
