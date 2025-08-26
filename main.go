package main

import "fmt"

type PayMethod interface {
	Account
	Pay(amount int) bool
}

type Account interface {
	GetBalance() int
}

type CreditCard struct {
	balance int
	limit   int
}

func (c *CreditCard) Pay(amount int) bool {
	if c.balance+amount <= c.limit {
		c.balance += amount
		fmt.Println("信用卡支付成功：", amount)
		return true
	}
	fmt.Println("信用卡支付失败：余额不足")
	return false
}

func (c *CreditCard) GetBalance() int {
	return c.balance
}

type DebitCard struct {
	balance int
}

func (c *DebitCard) Pay(amount int) bool {
	if c.balance >= amount {
		c.balance -= amount
		fmt.Println("借记卡支付成功：", amount)
		return true
	}
	fmt.Println("借记卡支付失败：余额不足")
	return false
}

func (c *DebitCard) GetBalance() int {
	return c.balance
}

func purchaseItem(p PayMethod, price int) {
	if p.Pay(price) {
		fmt.Println("购买成功剩余：", p.GetBalance())
	} else {
		fmt.Println("购买失败")
	}
}

func main() {

	creditCard := &CreditCard{balance: 0, limit: 1000}
	fmt.Println("信用卡支付")
	purchaseItem(creditCard, 800)

	debitCard := &DebitCard{balance: creditCard.balance}
	fmt.Println("借记卡支付")
	purchaseItem(debitCard, 400)
	fmt.Println("借记卡支付")
	go purchaseItem(debitCard, 700)

	var account Account = debitCard
	fmt.Println(account.GetBalance())
}
