package main

import (
  "fmt"
  "go_sample/config"
  "go_sample/database"
)

func main() {
  // Configの初期化
  config.Init()

  // DBの初期化
  database.Init()
  defer database.Close()

  // mの取得
  m := database.GetM()

  // マイグレーションのDownコマンド実行
  err := m.Down()
  if err != nil {
    panic(err)
  }

  fmt.Println("Execution m.Down() !!")
}