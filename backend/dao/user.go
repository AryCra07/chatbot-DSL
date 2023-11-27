package dao

import (
	"backend/consts"
	"backend/global"
	"backend/log"
	"backend/model"
)

type User struct {
	model.User
}

// BeforeSave check state cannot be negative
/*
 * @return err: error
 */
func (user *User) BeforeSave() (err error) {
	if user.State < 0 {
		log.Error(consts.Dao, "State cannot be negative!")
		return err
	}
	return nil
}

// Authentication check user's name and password
/*
 * @param name: user's name
 * @param password: user's password
 * @return user: user
 * @return bool: whether the user is authenticated
 * @return err: error
 */
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

// InsertUser insert user
/*
 * @param user: user
 * @return err: error
 */
func InsertUser(user model.User) error {
	err := global.DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// IsUsernameExist check whether the username exists
/*
 * @param name: user's name
 * @return bool: whether the username exists
 */
func IsUsernameExist(name string) bool {
	var user model.User
	result := global.DB.Where("name = ?", name).First(&user)
	if result.Error != nil {
		log.Warning(consts.Dao, "This user does not exist")
		return false
	}
	return true
}

// GetUserInfo get user's information
/*
 * @param name: user's name
 * @return user: user
 * @return bool: whether the user exists
 */
func GetUserInfo(id string) (model.User, bool) {
	var user model.User
	result := global.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		log.Error(consts.Dao, "Query error when executing GetUserInfo")
		return user, false
	}
	return user, true
}

// UpdateUserState update user's state
/*
 * @param name: user's name
 * @param state: user's state
 * @return err: error
 */
func UpdateUserState(id string, state int32) error {
	var user model.User
	result := global.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		log.Error(consts.Dao, "Query error when executing UpdateUserState")
		return result.Error
	}
	result = global.DB.Model(&user).Update("state", state)
	if result.Error != nil {
		log.Error(consts.Dao, "Update error when executing UpdateUserState")
		return result.Error
	}
	return nil
}

// UpdateUserWallet update user's wallet
/*
 * @param name: user's name
 * @param wallet: user's wallet
 * @return err: error
 */
func UpdateUserWallet(id string, balance float32, bill float32) error {
	var user model.User
	result := global.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		log.Error(consts.Dao, "Query error when executing UpdateUserWallet")
		return result.Error
	}
	result = global.DB.Model(&user).Update("balance", balance).Update("bill", bill)
	if result.Error != nil {
		log.Error(consts.Dao, "Update error when executing UpdateUserWallet")
		return result.Error
	}
	log.Info(consts.Dao, "Update user's wallet successfully")
	return nil
}
