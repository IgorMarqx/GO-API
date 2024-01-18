package main

import (
	"apiGo/cmd/database"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	database.ConnectDB()
	g := gin.Default()

	g.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status":     "success",
			"message":    "pong",
			"timestamps": time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	g.Run(":3000")
}
