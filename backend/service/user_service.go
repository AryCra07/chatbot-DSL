package service

import (
	"backend/consts"
	"backend/dao"
	"backend/log"
	"backend/model"
	"backend/utils"
)

func NewUser(username string, password string) (string, bool) {
	user, flag, err := dao.Authentication(username, utils.GetMd5(password))
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

// Login 用户登录 + 注册
func Login(username string, password string) (string, bool) {
	// 判断数据库当中 username 是否有重复
	flag := dao.IsUsernameExist(username)
	if flag {
		log.Warning(consts.Service, "User has exist")
		id, ok := NewUser(username, password)
		return id, ok
	}

	// 插入数据库
	user := model.User{
		Id:       utils.GetUUID(),
		Name:     username,
		Password: utils.GetMd5(password),
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
