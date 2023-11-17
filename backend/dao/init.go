package dao

import (
	"backend/config"
	"backend/consts"
	"backend/global"
	"backend/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// InitGorm init gorm
/*
 * @return *gorm.DB: gorm db
 */
func InitGorm() (*gorm.DB, error) {
	dsn := config.Dsn(global.Config)
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

// CloseDB close db
/*
 * @param db: gorm db
 * @param err: error
 */
func CloseDB(db *gorm.DB, err error) {
	sqlDB, _ := db.DB()
	if err != nil {
		err := sqlDB.Close()
		if err != nil {
			return
		}
	}
}
