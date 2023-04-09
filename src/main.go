package main

import (
	// Ginをインポート
	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

	// JSON形式で「"message": "Hello World"」を出力するAPI
	router.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "Hello World",
    })
	})

	router.Run(":3001")
}