package main

import (
	"fmt"
	"sync"
)

func writeInt(ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func readInt(ch chan int, flag string, wg *sync.WaitGroup) {
	for {
		val, ok := <-ch
		if !ok {
			wg.Done()
			break
		}
		fmt.Println(flag, val)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 10)

	wg.Add(3)
	go writeInt(ch, &wg)

	//time.Sleep(1000 * time.Millisecond)

	go readInt(ch, "demo1", &wg)
	go readInt(ch, "demo2", &wg)

	wg.Wait()
}
