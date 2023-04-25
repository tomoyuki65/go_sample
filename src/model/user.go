package model

import (
  "time"
  "go_sample/database"
  "errors"
)

// ユーザー情報の構造体
// usersテーブルのカラム名でjson出力するため、各項目に「`json:""`」を設定します。
type User struct {
  Id        int        `json:"id"`
  Uid       string     `json:"uid"`
  Name      string     `json:"name"`
  Email     string     `json:"email"`
  CreatedAt time.Time  `json:"created_at"`
  UpdatedAt time.Time  `json:"updated_at"`
  DeletedAt *time.Time `json:"deleted_at"`
}

// DBからusersテーブルの全レコードを取得
func GetUsersAll() (*[]User, error) {
  db := database.GetDB()

  var users *[]User
  result := db.Find(&users)
  if result.Error != nil {
    return nil, result.Error
  } else if result.RowsAffected == 0 {
    return nil, errors.New("users not registerd")
  }

  return users, nil
}