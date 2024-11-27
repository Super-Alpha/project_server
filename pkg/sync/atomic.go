package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// 公共变量
	var val int64
	wp := &sync.WaitGroup{}

	wp.Add(3)

	for i := 0; i < 3; i++ {
		go func(v int) {
			atomic.AddInt64(&val, int64(v)) // 0+1+2
			wp.Done()
		}(i)
	}

	wp.Wait()

	fmt.Printf("val = %d\n", atomic.LoadInt64(&val))
}
