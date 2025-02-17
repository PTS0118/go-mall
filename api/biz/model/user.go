package model

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	Base
	UserName string `json:"user_name" column:"user_name"`
	Email    string `json:"email" column:"email"`
	Password string `json:"password" column:"password"`
}

func (u User) TableName() string {
	return "user"
}

func GetByEmail(db *gorm.DB, ctx context.Context, email string) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where(&User{Email: email}).First(&user).Error
	println(user)
	return
}

func Create(db *gorm.DB, ctx context.Context, user *User) error {
	return db.WithContext(ctx).Create(user).Error
}
