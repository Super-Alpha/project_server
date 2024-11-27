package main

import (
	"fmt"
	"sync"
)

func once(flag string, wp *sync.WaitGroup, on *sync.Once) {
	// 只执行一次，多用于初始化操作
	on.Do(func() {
		fmt.Printf("flag:%s\n", flag)
	})

	wp.Done()
}

func main() {
	wp := &sync.WaitGroup{}
	on := &sync.Once{}

	wp.Add(2)

	go once("go1", wp, on)
	go once("go2", wp, on)

	wp.Wait()
}
