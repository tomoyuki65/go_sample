package controller

import (
  "github.com/gin-gonic/gin"
  "go_sample/config"
)

func ConfigIndex(c *gin.Context) {
  // config設定を取得
  cfg := config.GetConfig()

  c.JSON(200, gin.H{
    "ENV": cfg.Env,
    "Tz": cfg.Tz,
    "DB.Type": cfg.Db.Type,
    "DB.Host": cfg.Db.Host,
    "DB.Port": cfg.Db.Port,
    "DB.Charset": cfg.Db.Charset,
    "DB.ParseTime": cfg.Db.ParseTime,
    "DB.Loc": cfg.Db.Loc,
    "DB.Database": cfg.Db.Database,
    "DB.User": cfg.Db.User,
  })
}