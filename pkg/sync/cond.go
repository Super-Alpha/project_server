package main

import (
	"fmt"
	"sync"
	"time"
)

/*
     广播唤醒开启的协程

	(1) 5个Goroutine 通过 sync.Cond.Wait 等待特定条件的满足；
	(2) 1个Goroutine 会调用 sync.Cond.Broadcast 唤醒所有陷入等待的 Goroutine；

	sync.Cond.Wait 在调用之前一定要使用获取互斥锁，否则会触发程序崩溃；
	sync.Cond.Signal 唤醒的 Goroutine 都是队列最前面、等待最久的 Goroutine；
	sync.Cond.Broadcast 会按照一定顺序广播通知等待的全部 Goroutine；
*/

var done = false

func signal(c *sync.Cond) {
	done = true
	c.Signal()
}

func broadcast(c *sync.Cond) {
	done = true
	c.Broadcast()
}

func listen(c *sync.Cond, signal int) {
	// Wait()调用前，一定要使用互斥锁，否则会触发程序崩溃
	c.L.Lock()

	for !done {
		c.Wait()
	}
	fmt.Println("listen", signal)

	c.L.Unlock()
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})

	for i := 0; i < 5; i++ {
		go listen(cond, i)
	}

	time.Sleep(1 * time.Second)

	// 唤醒一个协程
	//go signal(cond)
	//time.Sleep(2 * time.Second)

	// 唤醒所有堵塞的协程
	go broadcast(cond)

	time.Sleep(3 * time.Second)
}
