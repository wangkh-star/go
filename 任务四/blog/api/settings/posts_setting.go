package settings

import (
	"blog/global"
	"blog/models"
	"blog/models/res"
	"blog/utils"

	"github.com/gin-gonic/gin"
)

type PostsRequest struct {
	ID      uint   `json:"id"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  uint
}
type UpdPostsRequest struct {
	ID      uint   `json:"id"  binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  uint
}

func (Settingsapi) AddPosts(c *gin.Context) {
	var post PostsRequest
	if err := c.ShouldBindJSON(&post); err != nil {
		res.ResultFail("参数错误", c)
		return
	}

	userId, _ := utils.GetUserID(c)
	newPosts := models.PostsModel{
		Title:   post.Title,
		Content: post.Content,
		UserID:  userId,
	}
	result := global.DB.Create(&newPosts)
	if result.Error != nil {
		global.Log.Error(result.Error)
		res.ResultFail("文章创建失败", c)
		return
	}
	res.ResultOk(newPosts, "文章创建成功", c)
}

func (Settingsapi) ReadPosts(c *gin.Context) {
	id := c.Param("id")
	newPosts := models.PostsModel{}
	result := global.DB.First(&newPosts, id)
	if result.Error != nil {
		global.Log.Error(result.Error)
		res.ResultFail("文章查询失败", c)
		return
	}
	res.ResultOk(newPosts, "文章查询成功", c)
}

func (Settingsapi) ReadPostsAll(c *gin.Context) {
	var newPosts []models.PostsModel
	result := global.DB.Where("").Find(&newPosts)
	if result.Error != nil {
		global.Log.Error(result.Error)
		res.ResultFail("文章查询失败", c)
		return
	}
	res.ResultOk(newPosts, "文章查询成功", c)
}

func (Settingsapi) UpdPosts(c *gin.Context) {
	var post UpdPostsRequest
	if err := c.ShouldBindJSON(&post); err != nil {
		res.ResultFail("参数错误", c)
		return
	}
	newPosts := models.PostsModel{}
	result := global.DB.First(&newPosts, post.ID)
	if result.Error != nil {
		global.Log.Error(result.Error)
		res.ResultFail("文章不存在", c)
		return
	}
	userId, _ := utils.GetUserID(c)
	if newPosts.UserID != userId {
		res.ResultFail("不能修改不属于自己的文章", c)
		return
	}

	newPosts = models.PostsModel{
		Title:   post.Title,
		Content: post.Content,
		UserID:  userId,
	}
	result = global.DB.Updates(&newPosts)
	if result.Error != nil {
		global.Log.Error(result.Error)
		res.ResultFail("文章修改失败", c)
		return
	}
	res.ResultOk(newPosts, "文章修改成功", c)
}

func (Settingsapi) DelPosts(c *gin.Context) {
	id := c.Param("id")
	newPosts := models.PostsModel{}
	result := global.DB.First(&newPosts, id)
	if result.Error != nil {
		global.Log.Error(result.Error)
		res.ResultFail("文章不存在", c)
		return
	}
	userId, _ := utils.GetUserID(c)
	if newPosts.UserID != userId {
		res.ResultFail("不能删除不属于自己的文章", c)
		return
	}
	result = global.DB.Delete(&newPosts)
	if result.Error != nil {
		global.Log.Error(result.Error)
		res.ResultFail("文章删除失败", c)
		return
	}
	res.ResultOk(newPosts, "文章删除成功", c)
}
