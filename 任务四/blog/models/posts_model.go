package models

type PostsModel struct {
	Model
	Title   string `gorm:"size:64" json:"title"`
	Content string `gorm:"size:256" json:"content"`

	// 修正外键配置
	UserID    uint      `json:"user_id" gorm:"index"`
	UserModel UserModel `gorm:"foreignKey:UserID;references:ID"`

	Comments []CommentsModel `gorm:"foreignKey:PostID;references:ID"`
}
