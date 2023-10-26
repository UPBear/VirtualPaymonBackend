package main

import (
	"ChatGptTest/general"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	GPTIns *general.ChatGPTVirtualCharacter
)

func chatgptTestHandler(c *gin.Context) {
	contextValue := c.DefaultQuery("context", "default context")
	resContext := GPTIns.ChatGPT(contextValue)

	c.JSON(200, gin.H{
		"context": resContext,
	})
}

func init() {
	GPTIns = general.NewChatGPTVirtualCharacter("")
	go GPTIns.StartGPT()
}

func main() {
	r := gin.Default()

	// 添加 CORS 中间件
	// 下面的配置允许所有源进行跨域请求
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	// 你可以根据需要进行更详细的配置，如设置特定的源、HTTP 方法、HTTP 头等
	r.Use(cors.New(config))
	r.GET("/api/chatgpt-test", chatgptTestHandler)

	r.Run(":8000")
}
