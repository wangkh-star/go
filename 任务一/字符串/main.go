package main

import "fmt"

// 创建映射，用于匹配右括号和对应的左括号
var bracketMap = map[byte]byte{
	')': '(',
	'}': '{',
	']': '[',
}

func main() {
	var strs = "{)}"
	temp := make([]byte, 0)
	for i := 0; i < len(strs); i++ {
		var char = strs[i]
		if left, flag := bracketMap[char]; flag {
			if len(temp) == 0 || temp[len(temp)-1] != left {
				fmt.Println("不匹配")
			}
		} else {
			temp = append(temp, char)
		}
	}
}
