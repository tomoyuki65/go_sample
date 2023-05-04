package controller

import (
  "github.com/gin-gonic/gin"
  "go_sample/model"
  "net/http"
  "strconv"
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

// ユーザーを作成
func UserCreate(c *gin.Context) {
  // リクエストボディをバインド
  var req model.User
  if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }

  if err := model.CreateUser(req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, req)
}

// ユーザー１件をJSON形式で出力
func UserShow(c *gin.Context) {
  // パラメータ「:id」を数値変換
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }

  user, err := model.GetUser(id)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "message": err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, user)
}

// ユーザー情報を更新
func UserUpdate(c *gin.Context) {
  // パラメータ「:id」を数値変換
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }

  // リクエストボディをバインド
  var req model.User
  if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }

  if err := model.UpdateUser(id, req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, req)
}

// ユーザーを削除
func UserDelete(c *gin.Context) {
  // パラメータ「:id」を数値変換
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }

  if err := model.DeleteUser(id); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "message": "user(" + strconv.Itoa(id) + ") is deleted",
  })
}