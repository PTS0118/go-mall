package model

import (
	"context"
	"github.com/PTS0118/go-mall/api/biz/dal/mysql"
	"gorm.io/gorm"
)

type User struct {
	Base
	Username  string `json:"username" column:"username"`
	Email     string `json:"email" column:"email"`
	Password  string `json:"password" column:"password"`
	Telephone string `json:"telephone" column:"telephone"`
	Role      string `json:"role" column:"role"`
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

func CreateUsers(users []User) error {
	return mysql.DB.Create(users).Error
}

func FindUserByNameOrEmail(userName, email string) ([]User, error) {
	res := make([]User, 0)
	if err := mysql.DB.Where(mysql.DB.Or("username = ?", userName).
		Or("email = ?", email)).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func CheckUser(account, password string) ([]User, error) {
	res := make([]User, 0)
	if err := mysql.DB.Where(mysql.DB.Or("username = ?", account).
		Or("email = ?", account)).Where("password = ?", password).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
