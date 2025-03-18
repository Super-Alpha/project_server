package main

import (
	"fmt"
	"sync"
	"time"
)

/*
    互斥锁（Mutex）是一种独占锁。它用于保护共享资源，使得同一时刻只有一个goroutine能访问该资源。其有两种模式，正常模式、饥饿模式
	在正常模式下，锁的等待者会按照先进先出的顺序获取锁，但是刚被唤醒的Goroutine与新创建的Goroutine竞争时，大概率会获取不到锁，为了减少这种情况的出现，
	一旦有goroutine超过1ms没有获取到锁，则会进入饥饿模式，防止部分Goroutine被饿死。

	在饥饿模式中，互斥锁会直接交给等待队列最前面的 Goroutine。新的 Goroutine 在该状态下不能获取锁、也不会进入自旋状态，它们只会在队列的末尾等待。
	如果一个 Goroutine 获得了互斥锁并且它在队列的末尾或者它等待的时间少于 1ms，那么当前的互斥锁就会切换回正常模式。
*/

/*
    读写锁（RWMutex）是一种读多写少锁。它用于保护被多个 goroutine 同时读取但只有一个 goroutine 写入的共享资源。
	在读操作时，可以允许多个goroutine并发访问，并发性能更高。而在写操作时，只允许一个 goroutine 进入临界区域，其他读写请求都会被阻塞。
	读写锁适用于读操作非常频繁，写操作相对较少的场景。读写锁通过 RLock()、RUnlock()、Lock() 和 Unlock() 方法来获取和释放锁。
	读写互斥锁在互斥锁之上提供了额外的更细粒度的控制，能够在读操作远远多于写操作时提升性能。
*/

var (
	v      int64
	wg     sync.WaitGroup
	mLock  sync.Mutex   // 互斥锁（悲观锁类型）
	rwLock sync.RWMutex // 读写锁（乐观锁类型）
)

func write() {
	mLock.Lock() // 加互斥锁
	//rwLock.Lock() // 加写锁

	v = v + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒

	mLock.Unlock() // 解互斥锁
	//rwLock.Unlock() // 解写锁

	wg.Done()
}

func read() {
	mLock.Lock() // 加互斥锁
	//rwLock.RLock() // 加读锁

	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	fmt.Printf("数值:%d\n", v)

	mLock.Unlock() // 解互斥锁
	//rwLock.RUnlock() // 解读锁

	wg.Done()
}

func main() {
	start := time.Now()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()

	end := time.Now()
	fmt.Printf("耗时:%fs\n", end.Sub(start).Seconds())
}
