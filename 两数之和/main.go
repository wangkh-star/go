package main

import "fmt"

func main() {
	//给定整数数组
	var num []int = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var target int = 10

	var tempMap = make(map[int]int)
	for k, v := range num {
		var tempUnm = target - v
		if index, isext := tempMap[tempUnm]; isext {
			fmt.Printf("找到和为%d的两个整数: nums[%d]=%d 和 nums[%d]=%d\n",
				target, index, tempUnm, k, v)
		}
		tempMap[v] = k
	}

	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			if num[i]+num[j] == target {
				fmt.Println("1和为目标值的那两个整数", num[i], "-", num[j])
			}
		}
	}

}
