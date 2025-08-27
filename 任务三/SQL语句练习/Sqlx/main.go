package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	ID        int
	Name      string
	Age       int
	Grade     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	dsn := "root:wangkh@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("链接异常")
	}
	db.AutoMigrate(&Student{})
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})

}
