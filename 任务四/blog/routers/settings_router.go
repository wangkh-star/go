package routers

import (
	"blog/api"
	"blog/core"
)

func (r RouterGroup) SettingRouter() {
	settingsapi := api.ApiGroupApp.Settingsapi
	public := r.Group("/api/v1")
	{
		//系统登录注册模块
		public.POST("/login", settingsapi.Login)
		public.POST("/register", settingsapi.Register)

		//文章新增阅读修改模块
		public.POST("/addPost", core.JWTAuth(), settingsapi.AddPosts)
		public.GET("/readPost/:id", settingsapi.ReadPosts)
		public.GET("/readPostsAll", settingsapi.ReadPostsAll)
		public.POST("/updPost", core.JWTAuth(), settingsapi.UpdPosts)
		public.GET("/delPost/:id", settingsapi.DelPosts)

		//文章评论功能
		public.POST("/addCommonts", core.JWTAuth(), settingsapi.AddCommonts)
		public.GET("/queryCommonsByPostId", settingsapi.QueryCommonsByPostId)
	}

}
