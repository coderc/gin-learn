package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.New()

	// 创建一个路由组
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "v1 pong"})
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "v2 pong"})
		})
	}

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
