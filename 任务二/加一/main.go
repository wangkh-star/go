package main

import "fmt"

func main() {
	var str []int = []int{1, 2, 3, 4, 9}
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] < 9 {
			str[i]++
			fmt.Println(str)
			return
		}
		str[i] = 0 // 当前位设为0，进位到前一位
	}
	str = append([]int{1}, str...)
	fmt.Println(str)
}
