package main

import (
	"fmt"
	"time"
)

func main() {
	go onePtoTen()
	go onePtoTen2()
	time.Sleep(100 * time.Millisecond)
}

func onePtoTen() {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
func onePtoTen2() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
