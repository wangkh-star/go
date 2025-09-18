package settings

import (
	"blog/global"
	"blog/models"
	"blog/models/res"
	"blog/utils"

	"github.com/gin-gonic/gin"
)

type ComRequest struct {
	Content string `json:"content" binding:"required"`
	PostID  uint   `json:"post_id" binding:"required"`
	UserID  uint
}

func (Settingsapi) AddCommonts(c *gin.Context) {
	var addCom ComRequest
	if err := c.ShouldBindJSON(&addCom); err != nil {
		res.ResultFail("参数错误", c)
		return
	}
	userId, _ := utils.GetUserID(c)
	newComs := models.CommentsModel{
		Content: addCom.Content,
		PostID:  addCom.PostID,
		UserID:  userId,
	}
	result := global.DB.Create(&newComs)
	if result.Error != nil {
		global.Log.Error(result.Error)
		res.ResultFail("评论失败", c)
		return
	}
	res.ResultOk(newComs, "评论成功", c)

}

func (Settingsapi) QueryCommonsByPostId(c *gin.Context) {
	id := c.Param("id")
	var coms []models.CommentsModel
	result := global.DB.Where("post_id=?", id).Find(&coms)
	if result.Error != nil {
		global.Log.Error(result.Error)
		res.ResultFail("获取评论失败", c)
		return
	}
	res.ResultOk(coms, "获取评论成功", c)

}
