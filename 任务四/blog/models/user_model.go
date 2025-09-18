package models

type UserModel struct {
	Model
	Username string `gorm:"size:64" json:"user_name"`
	Password string `gorm:"size:255" json:"password"`
	Email    string `gorm:"size:64" json:"email"`
}
