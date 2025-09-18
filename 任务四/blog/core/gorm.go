package core

import (
	"blog/global"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {

	if global.Config.Mysql.Host == "" {
		log.Println("未配置db")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()
	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	log.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: mysqlLogger})

	if err != nil {
		log.Fatalf("mysql链接失败")
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour * 4)
	global.DB = db
	log.Println("mysql链接成功")
	return db
}
