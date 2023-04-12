package main

import (
	"go_sample/config"
  "go_sample/router"
)

func main() {
	// configの初期化
  config.Init()

  // routerの初期化
  router.Init()
}