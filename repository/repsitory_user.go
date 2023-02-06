package repository

import (
	"gin-demo/model/dao"
)

func FindUserByAccount(account string) (*dao.UserDao, error) {
	var user *dao.UserDao
	err := db.Where("account = ?", account).First(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func FindUserById(id int64) (*dao.UserDao, error) {
	var user *dao.UserDao
	err := db.Where("id = ?", id).First(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func InsertUser(user *dao.UserDao) error {
	if err := db.Select("id", "account", "password", "login_time").Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *dao.UserDao) error {
	if err := db.Model(user).Where("id = ?", user.Id).Updates(map[string]interface{}{"password": user.Password}).Error; err != nil {
		return err
	}
	return nil
}
