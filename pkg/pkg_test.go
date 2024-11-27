package pkg

import (
	"bytes"
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNumPrint(t *testing.T) {
	wg := sync.WaitGroup{}
	lock := new(sync.Mutex)
	var a int32 = 0
	var b int32 = 2
	for i := 0; i < 5; i++ {
		go func() {
			if a > b {
				fmt.Println("done")
				return
			}
			lock.Lock()
			defer lock.Unlock()
			a++
			fmt.Printf("i: %d a: %d \n", i, a)
		}()
	}
	wg.Wait()
}

func TestSlicePrint(t *testing.T) {
	a := []byte("AAAA/BBBBB")
	index := bytes.IndexByte(a, '/') // 4

	fmt.Println(cap(a))

	b := a[:index]   // AAAA
	c := a[index+1:] // BBBBB

	b = append(b, "CCC"...)

	fmt.Println(string(a))
	fmt.Println(string(b))
	fmt.Println(string(c))
}

func main() {
	// 创建一个带有超时的父 context
	parentCtx, cancelParent := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelParent() // 确保在 main 函数结束时释放资源

	// 创建一个子 context，继承父 context 的超时
	childCtx, cancelChild := context.WithCancel(parentCtx)
	defer cancelChild() // 确保在不再需要子 context 时释放资源

	// 启动一个 goroutine 来监听父 context 的取消
	go func() {
		select {
		case <-parentCtx.Done():
			fmt.Println("Parent context canceled:", parentCtx.Err())
		case <-time.After(10 * time.Second):
			// 这行代码实际上不会被执行，因为 parentCtx 会在 5 秒后被取消
		}
	}()

	// 启动一个 goroutine 来监听子 context 的取消
	go func() {
		select {
		case <-childCtx.Done():
			fmt.Println("Child context canceled:", childCtx.Err())
		case <-time.After(10 * time.Second):
			// 同样，这行代码也不会被执行
		}
	}()

	// 等待一段时间，观察输出
	time.Sleep(6 * time.Second)
}
