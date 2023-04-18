package database

import (
  "fmt"
  "log"
  "go_sample/config"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "github.com/golang-migrate/migrate/v4"
  _ "github.com/golang-migrate/migrate/v4/database/mysql"
  _ "github.com/golang-migrate/migrate/v4/source/file"
)

var db *gorm.DB
var m *migrate.Migrate

func Init() {
  // config設定を取得
  cfg := config.GetConfig()

  // DBの接続先設定
  dsn :=  fmt.Sprintf(
            "%s:%s@tcp(%s:%d)/%s",
            cfg.Db.User,
            cfg.Db.Password,
            cfg.Db.Host,
            cfg.Db.Port,
            cfg.Db.Database,
          )

  dsn_option := fmt.Sprintf(
                  "?charset=%s&parseTime=%t&loc=%s",
                  cfg.Db.Charset,
                  cfg.Db.ParseTime,
                  cfg.Db.Loc,
                )

  dsn_mysql := dsn + dsn_option

  // DBに接続
  var err error
  db, err = gorm.Open(mysql.Open(dsn_mysql), &gorm.Config{})
  if err != nil {
    panic(err)
  }

  // マイグレーション設定
  dsn_m :=  fmt.Sprintf(
              "%s://%s",
              cfg.Db.Type,
              dsn,
            )
  m, err =  migrate.New(
              cfg.Migrate.FilePath,
              dsn_m,
            )
  if err != nil {
    panic(err)
  }

  // マイグレーションの実行
  err = m.Up()
  // エラーメッセージが「no change」の場合もスキップ
  if err != nil && err.Error() != "no change" {
    log.Printf("m.Up() Error Message: %s\n", err)
  }
}

func GetDB() *gorm.DB {
  return db
}

func GetM() *migrate.Migrate {
  return m
}

func Close() {
  getDB, err := db.DB()
  if err != nil {
    panic(err)
  }
  getDB.Close()
}