package main

import "fmt"

func main() {
	var a int = 10
	FuncA(&a)
	fmt.Println(a)

	var num = []int{1, 2, 3, 4, 5, 6, 7}
	FuncB(&num)
	fmt.Println(num)
}

func FuncA(num *int) {
	*num += 10
}

func FuncB(num *[]int) {

	for i := 0; i < len(*num); i++ {
		(*num)[i] *= 2
	}
}
