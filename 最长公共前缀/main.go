package main

import (
	"fmt"
	"strings"
)

func main() {
	var str []string = []string{"1", "2", "4"}
	first := str[0]
	for i := 1; i < len(str); i++ {
		for strings.Index(str[i], first) != 0 {
			first = first[:len(first)-1]
			if first == "" {
				fmt.Println("无")
			}
		}
	}
	fmt.Println(first)
}
