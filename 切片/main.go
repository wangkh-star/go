package main

import (
	"fmt"
	"sort"
)

func main() {
	//给定非空数组
	var num []int = []int{1, 2, 3, 4, 5, 2, 3, 4, 5}
	fmt.Println(getlen(num))

	var num2 [][]int = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(num2))
}

func getlen(num []int) int {
	var t int = 0 // 慢指针，

	for i := 1; i < len(num); i++ {
		if num[i] != num[t] {
			t++
			num[t] = num[i]
		}
	}
	return (t + 1)

}

func merge(num [][]int) [][]int {
	sort.Slice(num, func(i, j int) bool {
		return num[i][0] < num[j][0]
	})

	temp := make([][]int, 0)

	for _, v := range num {
		if len(temp) == 0 || temp[len(temp)-1][1] < v[0] {
			temp = append(temp, v)
		} else {
			// 合并区间：更新最后一个区间的结束位置
			if v[1] > temp[len(temp)-1][1] {
				temp[len(temp)-1][1] = v[1]
			}
		}
	}
	return temp
}
