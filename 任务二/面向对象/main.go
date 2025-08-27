package main

import "fmt"

//定义接口
type Shape interface {
	Area() int
	Perimeter() int
}

//构造体
type Rectangle struct {
	A int
}

//构造体
type Circle struct {
	B int
}

//实现接口
func (c Circle) Area() int {
	return c.B + 10
}
func (c Circle) Perimeter() int {
	return c.B + 11
}

func (c Rectangle) Area() int {
	return c.A + 12
}
func (c Rectangle) Perimeter() int {
	return c.A + 13
}

func main() {
	var s Shape
	s = Circle{1}
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
	s = Rectangle{2}
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())

}
