package main

import (
	"blog/core"
	"blog/flag"
	"blog/global"
	"blog/routers"
)

func main() {
	//配置文件的加载
	core.InitConf()
	//初始化数据库连接
	core.InitGorm()

	//初始化表
	option := flag.Parse()
	core.Info(option)
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	//初始化路由
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	core.Info("程序运行 ", addr)
	err := router.Run(addr)
	if err != nil {
		core.Error("程序运行异常 ", err)
	}
}
