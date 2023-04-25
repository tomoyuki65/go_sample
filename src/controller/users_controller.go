package controller

import (
  "github.com/gin-gonic/gin"
  "go_sample/model"
  "net/http"
)

// ユーザー一覧をJSON形式で出力
func UsersIndex(c *gin.Context) {
  users, err := model.GetUsersAll()
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, users)
}