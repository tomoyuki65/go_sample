package main

import (
  "fmt"
  "go_sample/config"
  "go_sample/database"
  "go_sample/seed"
)

func main() {
  // Configの初期化
  config.Init()

  // DBの初期化
  database.Init()
  defer database.Close()

  // DBの取得
  db := database.GetDB()

  // seedの実行
  err := seed.UserSeeds(db)
  if err != nil {
    panic(err)
  }

  fmt.Println("Execution user_seeds.go !!")
}