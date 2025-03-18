package main

import (
	"fmt"
	"sync"
)

// sync.pool 管理临时对象的复用，能够有效地减少内存分配的开销，提高性能，减少内存的分配和GC的压力。

// New：当Pool中没有可用对象时，会调用此函数生成新对象。
// Put：将对象放回Pool中
// Get：从Pool中获取对象

func main() {
	pool := &sync.Pool{
		//设置New函数，会在池中没有可用对象时调用
		New: func() any {
			fmt.Println("New")
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

	for v := range ch1 {
		println(v)
		if v == 3 {
			pool.Put(ch1)
			fmt.Println("end1")
			break
		}
	}

	// -------------------------------

	ch2 := pool.Get().(chan int)
	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}(ch2)

	for v := range ch2 {
		println(v)
		if v == 5 {
			pool.Put(ch2)
			fmt.Println("end2")
			return
		}
	}

}
