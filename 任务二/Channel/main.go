package main

import (
	"fmt"
)

// 发送channel
func send(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("无缓冲通道发送channel", i)
	}
	close(ch)
}

// 读取channel
func read(ch <-chan int) {
	for v := range ch {
		fmt.Println("无缓冲通道接收读取channel", v)
	}
}
func main() {

	var ch = make(chan int)

	go send(ch)

	read(ch)

	var ch2 = make(chan int, 10)

	go send(ch2)

	read(ch2)

	// var timeOut = time.After(1 * time.Second)

	// for {
	// 	select {
	// 	case v, ok := <-ch:
	// 		if ok {
	// 			fmt.Println("关闭")
	// 		}
	// 		fmt.Println("接收到", v)

	// 	case <-timeOut:
	// 		fmt.Println("操作超时")
	// 		return
	// 	default:
	// 		fmt.Println("等待")
	// 		time.Sleep(500 * time.Millisecond)
	// 	}

	// }

}
