package global

import (
	"blog/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	//配置信息
	Config *config.Config
	Log    *logrus.Logger
	//连接数据库
	DB *gorm.DB
)
