package dao

import (
	"backend/global"
	"backend/model"
)

func Authentication(name string, password string) (model.User, bool, error) {
	var user model.User
	err := global.DB.Where("name = ?", name).First(&user).Error
	if err != nil {
		return user, false, err
	}
	if user.Password == password {
		return user, true, nil
	}
	return user, false, nil
}

func InsertUser(user model.User) error {
	err := global.DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func IsUsernameExist(name string) (bool, error) {
	var user model.User
	result := global.DB.Where("name = ?", name).First(&user)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
