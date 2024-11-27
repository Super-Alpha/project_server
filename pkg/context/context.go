package context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Context 使用原则
1、不要把Context放在结构体中，要以参数的方式传递
2、以Context作为参数的函数方法，应该把Context作为第一个参数，放在第一位。
3、给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用`context.TODO`
4、Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递
5、Context是线程安全的，可以放心的在多个goroutine中传递
*/

func goroutine1(ctx context.Context, wg *sync.WaitGroup, duration time.Duration) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("goroutine1", ctx.Err())
			return
		case <-time.After(duration):
			fmt.Println("goroutine1 have done")
			return
		}
	}
}

func goroutine2(ctx context.Context, wg *sync.WaitGroup, duration time.Duration) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("goroutine2", ctx.Err())
			return
		case <-time.After(duration):
			fmt.Println("goroutine2 have done")
			return
		}
	}
}

func goroutine3(ctx context.Context, wg *sync.WaitGroup, duration time.Duration) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("goroutine3", ctx.Err())
			return
		case <-time.After(duration):
			fmt.Printf("goroutine3 value = %d\n", ctx.Value("label"))
			return
		}
	}
}

// WithCancelApplication 手动执行，退出相关协程
func WithCancelApplication() {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go goroutine1(ctx, wg, 100*time.Millisecond)

	time.Sleep(100 * time.Millisecond)

	defer cancel()

	wg.Wait()
}

// WithTimeoutApplication 手动设置超时时间，定时退出相关协程
func WithTimeoutApplication() {
	wg := &sync.WaitGroup{}

	// 设置将来的某一时刻，为退出相关协程的最后期限
	//ctx, cancel := context.WithDeadline(context.Background(), time.Date(2023, 11, 11, 23, 59, 59, 59, time.UTC))

	// 设置一段时间，为退出相关协程的最后期限;例：2s后终止协程
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	wg.Add(1)
	go goroutine2(ctx, wg, 1*time.Second)

	wg.Wait()
}

// WithValueApplication 父子协程之间传递数据
func WithValueApplication() {
	wg := &sync.WaitGroup{}

	p := context.WithValue(context.Background(), "label", 100)

	//ctx, cancel := context.WithTimeout(p, 1000*time.Millisecond)

	//defer cancel()

	wg.Add(1)
	go goroutine3(p, wg, 200*time.Millisecond)

	wg.Wait()
}
