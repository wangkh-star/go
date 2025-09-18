package flag

import (
	"blog/core"
	"blog/global"
	"blog/models"
)

func MakeDb() {
	if global.DB == nil {
		core.Error("数据库连接未初始化")
		return
	}
	err := global.DB.Migrator().AutoMigrate(&models.UserModel{}, &models.CommentsModel{}, &models.PostsModel{})
	if err != nil {
		core.Error("生成表结构失败")
		return
	}
	core.Error("生成表结构成功")
}
