package goroutine

import (
	"fmt"
	"sync"
)

func go1(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func go2(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case i := <-ch:
			fmt.Println(i)
			if i == 9 {
				return
			}
		default:
			fmt.Println("default")
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}

	ch := make(chan int, 10)

	wg.Add(1)
	go go1(ch, wg)

	wg.Add(1)
	go go2(ch, wg)

	defer close(ch)

	wg.Wait()
}
