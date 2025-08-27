package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id   uint
	Name string
	Age  uint
}

func main() {
	dsn := "root:wangkh@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//自动迁移
	db.AutoMigrate(&User{})
	u1 := User{Name: "张三111", Age: 14}
	db.Create(&u1) //创建

	var user []User
	db.Find(&user)
	fmt.Println(user)

}
