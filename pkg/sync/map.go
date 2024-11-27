package main

import (
	"fmt"
	"sync"
)

func sum(s []int, tag string, c chan *sync.Map, wg *sync.WaitGroup, once *sync.Once) {

	sum := 0
	for _, v := range s {
		sum += v
	}

	sm := &sync.Map{}

	sm.Store(tag, sum)

	c <- sm

	defer wg.Done()
}

func main() {

	var wg sync.WaitGroup
	var once sync.Once

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan *sync.Map, 1)

	wg.Add(1)
	go sum(s[:len(s)/2], "G1", c, &wg, &once)

	wg.Add(1)
	go sum(s[len(s)/2:], "G2", c, &wg, &once)

	x, y := <-c, <-c
	fmt.Println(x, y)

	close(c)

	wg.Wait()
}
