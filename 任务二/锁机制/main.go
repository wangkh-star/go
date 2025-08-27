package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var count = 0
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				count++
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)

	var count1 int64 = 0
	var wg1 sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&count1, 1)
			}
		}()
	}
	wg1.Wait()
	fmt.Println(count1)
}
