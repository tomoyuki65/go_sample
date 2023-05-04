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
  apiV1.GET("/users", controller.UsersIndex)

  // ユーザーのCRUD用APIを追加
  apiV1.POST("/users", controller.UserCreate)
  apiV1.GET("/users/:id", controller.UserShow)
  apiV1.PATCH("/users/:id", controller.UserUpdate)
  apiV1.DELETE("/users/:id", controller.UserDelete)

  // routerを起動
  router.Run(":3001")
}