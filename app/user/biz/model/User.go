package model

import (
	"context"
	"github.com/PTS0118/go-mall/app/user/biz/dal/mysql"
	"gorm.io/gorm"
)

type User struct {
	Base
	UserName  string `json:"username" column:"username"`
	Email     string `json:"email" column:"email"`
	Password  string `json:"password" column:"password"`
	Telephone string `json:"telephone" column:"telephone"`
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

func FindUserByNameOrEmail(userName, email *string, id int32) (res *User, err error) {
	// 初始化一个空的查询条件
	query := mysql.DB.Model(&User{})
	// 动态添加查询条件
	if userName != nil && *userName != "" {
		query = query.Or("username = ?", *userName)
	}
	if email != nil && *email != "" {
		query = query.Or("email = ?", *email)
	}
	if id != 0 {
		query = query.Or("id = ?", id)
	}

	// 执行查询
	if err = query.First(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
