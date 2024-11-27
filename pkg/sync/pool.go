package main

import (
	"fmt"
	"sync"
)

// sync.pool 是用来存储被分配了但是没有被使用的对象，并且可能会再次使用的对象，可以减少内存的分配和GC的压力。

// New：当Pool中没有可用对象时，会调用此函数生成新对象。
// Put：将对象放回Pool中
// Get：从Pool中获取对象

func main() {
	pool := &sync.Pool{
		New: func() any {
			return make(chan int)
		},
	}

	ch1 := pool.Get().(chan int)
	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}(ch1)

	ch2 := pool.New().(chan int)
	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}(ch2)

	for v := range ch1 {
		println(v)
		if v == 3 {
			pool.Put(ch1)
			break
		}
	}

	fmt.Println("end")

	for v := range ch2 {
		println(v)
		if v == 5 {
			pool.Put(ch2)
			return
		}
	}
}
