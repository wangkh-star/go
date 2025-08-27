package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "wangkh",
		})
	})
	router.Run()
}
