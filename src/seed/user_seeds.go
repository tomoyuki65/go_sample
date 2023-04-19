package seed

import (
  "time"
  "go_sample/model"
  "gorm.io/gorm"
)

func UserSeeds(db *gorm.DB) error {
  // １件目のデータ
  user1 :=  model.User{
              Uid:       "uid:go_sample_1",
              Name:      "Sample_User_1",
              Email:     "sample1@example.com",
              CreatedAt: time.Now(),
              UpdatedAt: time.Now(),
            }
  if err := db.Create(&user1).Error; err != nil {
    return err
  }

  // ２件目のデータ
  user2 :=  model.User{
              Uid:       "uid:go_sample_2",
              Name:      "Sample_User_2",
              Email:     "sample2@example.com",
              CreatedAt: time.Now(),
              UpdatedAt: time.Now(),
            }
  if err := db.Create(&user2).Error; err != nil {
    return err
  }

  return nil
}