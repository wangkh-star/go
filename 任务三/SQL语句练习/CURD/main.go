package main

import (
	"fmt"
	"log"
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
type Account struct {
	ID          int
	AccountName string
	Balance     int64
}
type Transaction struct {
	ID            int
	FromAccountId int
	ToAccountId   int
	Amount        int64
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
	//单条插入
	// var student = Student{Name: "张山2", Age: 22, Grade: "三年级"}
	// result := db.Create(&student)
	// println(result)
	// println(student.ID)

	//批量插入
	// var students []Student = []Student{
	// 	{Name: "张22", Age: 9, Grade: "三年级"},
	// 	{Name: "王1", Age: 2, Grade: "五年级"},
	// 	{Name: "王2", Age: 22, Grade: "一年级"},
	// 	{Name: "王45", Age: 12, Grade: "二年级"}}
	// db.Create(&students)

	//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	// var students []Student
	// result := db.Where("age>?", 18).Find(&students)
	// if result.Error != nil {
	// 	log.Fatal("查询失败:", result.Error)
	// }
	// fmt.Println(students)

	//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	// result := db.Model(&Student{}).Where("name=?", "张山11111").Update("grade", "四年级")
	// if result.Error != nil {
	// 	log.Fatal("更新失败:", result.Error)
	// }
	// if result.RowsAffected == 0 {
	// 	fmt.Println("未找到数据")
	// }

	//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	// result := db.Delete(&Student{}, "age<?", 15)
	// if result.Error != nil {
	// 	log.Fatal("s删除失败:", result.Error)
	// }
	// if result.RowsAffected == 0 {
	// 	fmt.Println("未找到数据")
	// }

	// var sccounts []Account = []Account{
	// 	{Balance: 1200, AccountName: "A"},
	// 	{Balance: 0, AccountName: "B"}}
	// db.Create(&sccounts)

	//事务

	err = TransFerMoney(db, "B", "A", 100)
	if err != nil {
		log.Fatal("转账失败:", err)
	}
}

func TransFerMoney(db *gorm.DB, fromName string, toName string, amount int64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var account Account
		if err := tx.Where("account_name=?", fromName).First(&account).Error; err != nil {
			return fmt.Errorf("转出账户不存在")
		}
		if account.Balance < amount {
			return fmt.Errorf("余额不足")
		}

		var toAccount Account
		if err := tx.Where("account_name=?", toName).First(&toAccount).Error; err != nil {
			return fmt.Errorf("转入账户不存在")
		}

		//在插入记录  //先扣款  //在入款
		var transaction Transaction = Transaction{FromAccountId: account.ID, ToAccountId: toAccount.ID, Amount: amount}
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}
		//先扣款
		if err := tx.Model(&account).Update("balance", account.Balance-amount).Error; err != nil {
			return err
		}
		if err := tx.Model(&toAccount).Update("balance", toAccount.Balance+amount).Error; err != nil {
			return err
		}
		fmt.Printf("转账成功: %d -> %d, 金额: %d元\n",
			fromName, toName, amount)

		return nil
	})
}
