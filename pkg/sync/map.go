package main

import (
	"fmt"
	"strconv"
	"sync"
)

func syncMap() {
	m := sync.Map{}
	for i := 0; i < 10; i++ {
		m.Store("label"+strconv.Itoa(i), i)
	}

	m.Range(func(key, value interface{}) bool {
		fmt.Printf("%s = %d\n", key, value)
		if value == 8 {
			return false
		}
		return true
	})
}

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
