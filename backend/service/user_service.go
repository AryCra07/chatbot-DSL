package service

import (
	"backend/consts"
	"backend/dao"
	"backend/log"
	"backend/model"
	"backend/utils"
)

// Login user login, if user not exist, insert into database
/*
 * @param username: user name
 * @param password: user password
 * @return id: user id
 * @return ok: login success or not
 */
func Login(username string, password string) (string, bool) {
	// check user exist or not
	flag := dao.IsUsernameExist(username)
	if flag {
		// user exist, login directly
		log.Warning(consts.Service, "User has exist")
		id, ok := enterExistUser(username, password)
		return id, ok
	}

	// build new user and insert into database
	user := model.User{
		Id:       utils.GetUUID(),
		Name:     username,
		Password: utils.GetSha256(password),
		State:    0,
		Balance:  0,
		Bill:     0,
	}

	err := dao.InsertUser(user)
	if err != nil {
		log.Error(consts.Service, "Insert Fail")
		return consts.NotExistId, false
	} else {
		log.Info(consts.Service, "Login successfully")
		return user.Id, true
	}
}

// enterExistUser login for existed user, check password
/*
 * @param username: user name
 * @param password: user password
 * @return id: user id
 * @return ok: login success or not
 */
func enterExistUser(username string, password string) (string, bool) {
	user, flag, err := dao.Authentication(username, utils.GetSha256(password))
	if err != nil {
		log.Error(consts.Service, "Query error")
		return consts.NotExistId, false
	}
	if !flag {
		log.Error(consts.Service, "Wrong username or password")
		return consts.NotExistId, false
	} else {
		log.Info(consts.Service, "Login successfully")
		return user.Id, true
	}
}
