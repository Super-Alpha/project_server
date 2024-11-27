package main

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"golang.org/x/sync/semaphore"
)

/*
	信号量是在并发编程中比较常见的一种同步机制，它会保证持有的计数器在0到初始化的权重之间，
	每次获取资源时都会将信号量中的计数器减去对应的数值，在释放时重新加回来，
	当遇到计数器大于信号量大小时就会进入休眠等待其他进程释放信号。
*/

func main() {
	ctx := context.Background()
	maxWorker := runtime.NumCPU()                  // 8
	sem := semaphore.NewWeighted(int64(maxWorker)) // 可同时执行的协程数量为8

	// 开启20个goroutine，但只允许8个同时执行
	for i := 1; i <= 20; i++ {
		// 请求信号量（即获取可执行操作的资源, 若获取不到，则阻塞）
		if err := sem.Acquire(ctx, 1); err != nil {
			fmt.Printf("goroutine %d acquire semaphore failed: %v\n", i, err)
			break
		}

		//if sem.TryAcquire(1) {} // 也可以用 TryAcquire 来获取资源，若获取不到，则返回 false，非阻塞

		go func(i int) {
			// 执行任务
			fmt.Printf("goroutine %d start running\n", i)
			time.Sleep(time.Second)
			fmt.Printf("goroutine %d stop running\n", i)
			// 释放信号量
			sem.Release(1)
		}(i)
	}
	// 等待所有 goroutine 执行完成
	time.Sleep(time.Second * 9)
}
