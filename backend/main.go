package main

import (
	"backend/config"
	"backend/consts"
	"backend/dao"
	"backend/global"
	"backend/log"
	"backend/router"
)

func main() {
	err := config.Init("config/config.yaml")
	if err != nil {
		log.Error(consts.Config, "Parse config.yaml error")
		return
	}

	global.DB, _ = dao.InitGorm()
	global.Config = config.GetYamlConfig()
	router.CreateServer()
}
