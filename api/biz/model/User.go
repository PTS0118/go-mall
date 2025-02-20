package model

import (
	"context"
	"github.com/PTS0118/go-mall/api/biz/dal/mysql"
	"github.com/PTS0118/go-mall/app/user/biz/model"
	"gorm.io/gorm"
)

type User struct {
	Base
	UserName string `json:"username" column:"username"`
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

func CreateUsers(users []*model.User) error {
	return mysql.DB.Create(users).Error
}

func FindUserByNameOrEmail(userName, email string) ([]*model.User, error) {
	res := make([]*model.User, 0)
	if err := mysql.DB.Where(mysql.DB.Or("user_name = ?", userName).
		Or("email = ?", email)).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func CheckUser(account, password string) ([]*model.User, error) {
	res := make([]*model.User, 0)
	if err := mysql.DB.Where(mysql.DB.Or("user_name = ?", account).
		Or("email = ?", account)).Where("password = ?", password).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
