package router

import (
  "github.com/gin-gonic/gin"
  "go_sample/controller"
)

func Init() {
  // routerの初期化
  router := gin.Default()

  // routerの設定
  router.GET("/", controller.Index)

	// APIの設定
	apiV1 := router.Group("/api/v1")
	apiV1.GET("/config", controller.ConfigIndex)

  // routerを起動
  router.Run(":3001")
}