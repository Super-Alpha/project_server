package main

import (
	"fmt"
	"sync"
)

// golang 直接将channel实例分配到堆上

/*
type hchan struct {
	qcount   uint           // channel中当前元素的个数
	dataqsiz uint           // 可以缓冲的数量,即channel中循环队列的长度
	buf      unsafe.Pointer // channel中缓冲区数据指针
	elemsize uint16         // 要发送或接收的数据类型大小
	closed   uint32         // 表示通道是否已关闭的标志，0表示未关闭，1表示已关闭
	elemtype *_type         // channel中元素的类型
	sendx    uint           // 当设置了缓冲时，数据区域即循环队列此时已发送数据的索引位置
	recvx    uint           // 当设置了缓冲时，数据区域即循环队列此时已接收数据的索引位置
	recvq    waitq          // 结构为双向链表，存储由于缓存不足而堵塞的执行接收操作的goroutine
	sendq    waitq          // 结构为双向链表，存储由于缓存不足而堵塞的执行发送操作的goroutine
	lock     mutex          // 互斥锁lock用于保护对hchan结构体的并发访问
}

type waitq struct {
	first  *sudog    // sudog表示等待队列中的goroutine
	last   *sudog
}
*/

/*
针对没有缓冲区的channel
	1、发送数据
    	G发送数据到channel时，先看channel中是否有阻塞的接收goroutine（recvq），如果有则直接拷贝数据给它，并唤醒它；如果没有则将该goroutine放入sendq中并挂起。
	2、接收数据
    	G从channel中接收数据时，先看channel中是否有阻塞的发送goroutine（sendq），如果有则直接从它里面拷贝数据，并唤醒它；如果没有则将该goroutine放入recvq中并挂起。
*/

/*
针对有缓存区的channel
 	1、发送数据
		先看是否有阻塞接收的G，如果有，则将数据拷贝给它，并唤醒它；
		如果没有，若缓存区未满，则将G1数据写进缓存区，若缓存区已满，则将G1挂靠到sendq队列。
 	2、接收数据
		若缓存区有数据，则G2从缓存区读取数据；若sendq队列有挂靠G1，则将该G1数据写进缓存区，并解除G1的休眠，等待调度；
		若缓存区没有数据，则将G1挂靠进recvq队列，等待唤醒调度。
*/

func G1(ch chan int, list []int, wg *sync.WaitGroup) {
	temp := 0

	defer wg.Done()

	for _, v := range list {
		temp += v
	}
	ch <- temp
}

func main() {
	res := 0
	ch := make(chan int, 2)
	wg := &sync.WaitGroup{}
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go G1(ch, list[i*5:(i+1)*5], wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		res += v
	}

	fmt.Printf("channel closed,val = %d\n", res)
}
