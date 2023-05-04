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

// ユーザーを作成
func CreateUser(r User) error {
  db := database.GetDB()

  // トランザクションの開始
  tx := db.Begin()

  user := User{
    Uid:       r.Uid,
    Name:      r.Name,
    Email:     r.Email,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }

  if err := tx.Create(&user).Error; err != nil {
    // エラーの場合はロールバック
    tx.Rollback()
    return err
  }
  // コミットしてトランザクションを終了
  tx.Commit()

  return nil
}

// 対象のユーザーを１件取得
func GetUser(id int) (*User, error) {
  db := database.GetDB()

  var user *User
  result := db.First(&user, id)
  if result.Error != nil {
    return nil, result.Error
  }

  return user, nil
}

// ユーザー情報を更新
func UpdateUser(id int, r User) error {
  db := database.GetDB()

  // トランザクションの開始
  tx := db.Begin()

  // nil値項目の更新なし、updated_atの更新あり
  err := tx.Updates(&User{
              Id:    id,
              Name:  r.Name,
              Email: r.Email,
            }).Error
  if err != nil {
    // エラーの場合はロールバック
    tx.Rollback()
    return err
  }
  // コミットしてトランザクションを終了
  tx.Commit()

  return nil
}

// ユーザー削除
func DeleteUser(id int) error {
  db := database.GetDB()

  // トランザクションの開始
  tx := db.Begin()

  // 対象ユーザーを物理削除
  if err := tx.Delete(&User{}, id).Error; err != nil {
    // エラーの場合はロールバック
    tx.Rollback()
    return err
  }
  // コミットしてトランザクションを終了
  tx.Commit()

  return nil
}