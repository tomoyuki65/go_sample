package controller

import (
  "github.com/gin-gonic/gin"
  "go_sample/model"
  "net/http"
)

// ユーザー一覧をJSON形式で出力
func UsersIndex(c *gin.Context) {
  users := model.GetUsersAll()

  c.JSON(http.StatusOK, users)
}