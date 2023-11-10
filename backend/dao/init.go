package dao

import (
	"backend/config"
	"backend/consts"
	"backend/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func InitGorm() (*gorm.DB, error) {
	dsn := config.Dsn(config.GetYamlConfig())
	var mysqlLogger logger.Interface

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		log.Error(consts.Dao, "DB Connect Error")
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	log.Info(consts.Dao, "Connect DB successfully")
	return db, err
}

func CloseDB(db *gorm.DB, err error) {
	sqlDB, _ := db.DB()
	if err != nil {
		err := sqlDB.Close()
		if err != nil {
			return
		}
	}
}
