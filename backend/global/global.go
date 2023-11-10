package global

import (
	"backend/config"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Config config.YamlConfig
)
