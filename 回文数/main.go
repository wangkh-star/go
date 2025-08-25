package main

import "fmt"

func main() {
	//给定非空数组
	var num []int = []int{1, 2, 3, 4, 5, 2, 3, 4, 5}

	//统计出现次数
	var count = make(map[int]int)

	for _, value := range num {
		if count[value] == 0 {
			count[value] = 1
		} else {
			count[value] = count[value] + 1
		}
	}
	for k, v := range count {
		fmt.Println("数字:", k, "出现次数", v)
	}

	/*运行结果*/
	// 数字: 5 出现次数 2
	// 数字: 1 出现次数 1
	// 数字: 2 出现次数 2
	// 数字: 3 出现次数 2
	// 数字: 4 出现次数 2

}
