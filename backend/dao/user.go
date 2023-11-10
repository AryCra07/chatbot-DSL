package dao

import (
	"backend/consts"
	"backend/global"
	"backend/log"
	"backend/model"
	"gorm.io/gorm"
)

type User struct {
	model.User
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if user.State < 0 {
		log.Error(consts.Dao, "State cannot be negative!")
		return err
	}
	return nil
}

func Authentication(name string, password string) (User, bool, error) {
	var user User
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

func IsUsernameExist(name string) bool {
	var user model.User
	result := global.DB.Where("name = ?", name).First(&user)
	if result.Error != nil {
		log.Warning(consts.Dao, "This user does not exist")
		return false
	}
	return true
}

func GetUserInfo(name string) (model.User, bool) {
	var user model.User
	result := global.DB.Where("name = ?", name).First(&user)
	if result.Error != nil {
		log.Error(consts.Dao, "Query error when executing GetUserInfo")
		return user, false
	}
	return user, true
}
