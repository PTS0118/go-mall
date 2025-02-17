package mysql

import (
	"github.com/PTS0118/go-mall/api/biz/model"
)

func CreateUsers(users []*model.User) error {
	return DB.Create(users).Error
}

func FindUserByNameOrEmail(userName, email string) ([]*model.User, error) {
	res := make([]*model.User, 0)
	if err := DB.Where(DB.Or("user_name = ?", userName).
		Or("email = ?", email)).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func CheckUser(account, password string) ([]*model.User, error) {
	res := make([]*model.User, 0)
	if err := DB.Where(DB.Or("user_name = ?", account).
		Or("email = ?", account)).Where("password = ?", password).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
