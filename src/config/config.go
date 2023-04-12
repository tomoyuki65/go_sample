package config

import (
  "fmt"
  "os"
  "github.com/spf13/viper"
  "github.com/fsnotify/fsnotify"
)

// config設定の構造体
type Config struct {
  Env     string
  Tz      string
  Db      Db      `yml:db`
  Migrate Migrate `yml:migrate`
}

type Db struct {
  Type      string  `yml:type`
  Host      string  `yml:host`
  Port      int     `yml:port`
  Charset   string  `yml:charset`
  ParseTime bool    `yml:parseTime`
  Loc       string  `yml:loc`
  Database  string
  User      string
  Password  string
}

type Migrate struct {
  FilePath string `yml:filePath`
}

var cfg *Config

func Init() {
  // viperの初期設定
  viper.SetConfigName("config_" + fmt.Sprintf("%s", os.Getenv("ENV")))
  viper.SetConfigType("yml")
  viper.AddConfigPath("config/")

  // configファイル更新時に再読み込み
  viper.WatchConfig()
  viper.OnConfigChange(func(e fsnotify.Event) {
    fmt.Println("Config file changed:", e.Name)
    viper.Unmarshal(&cfg)
  })

  // configファイルの読み込み
  err := viper.ReadInConfig()
  if err != nil {
    panic(err)
  }

  // 読み込んだデータを変数cfgに設定
  err = viper.Unmarshal(&cfg)
  if err != nil {
    panic(err)
  }

  // 環境変数の値を変数cfgに設定
  cfg.Env         = os.Getenv("ENV")
  cfg.Tz          = os.Getenv("TZ")
  cfg.Db.Database = os.Getenv("MYSQL_DATABASE")
  cfg.Db.User     = os.Getenv("MYSQL_USER")
  cfg.Db.Password = os.Getenv("MYSQL_PASSWORD")
}

func GetConfig() *Config {
  return cfg
}