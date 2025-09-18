package models

type CommentsModel struct {
	Model
	Content string `gorm:"size:256" json:"content"`

	// 使用默认命名约定：外键字段名 = 关联模型名 + ID
	PostID uint       `json:"post_id" gorm:"index"`
	Post   PostsModel `gorm:"foreignKey:PostID" json:"post"`

	UserID uint      `json:"user_id" gorm:"index"`
	User   UserModel `gorm:"foreignKey:UserID" json:"user"`
}
