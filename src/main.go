package main

import (
  "go_sample/config"
  "go_sample/database"
  "go_sample/router"
)

func main() {
  // configの初期化
  config.Init()

  // DBの初期化
  database.Init()
  defer database.Close()

  // routerの初期化
  router.Init()
}