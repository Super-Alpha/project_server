package main

import (
	"fmt"
	"sync"
)

func once(flag string, wg *sync.WaitGroup, on *sync.Once) {
	// 只执行一次，多用于初始化操作
	on.Do(func() {
		fmt.Printf("flag:%s\n", flag)
	})

	wg.Done()
}

func main() {
	wp := &sync.WaitGroup{}
	on := &sync.Once{}

	wp.Add(10)

	for i := 0; i < 10; i++ {
		go once(fmt.Sprintf("demo%d", i+1), wp, on)
	}

	wp.Wait()
}
