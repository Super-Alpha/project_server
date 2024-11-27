package main

import (
	"fmt"

	"github.com/pkg/errors"

	"golang.org/x/sync/errgroup"
)

/*
	等待所有开启的协程执行完成后，退出
*/

func main() {
	var g errgroup.Group

	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
	}

	for _, v := range urls {

		g.Go(func() error {
			if v != "" {
				fmt.Println(v)
			} else {
				return errors.New(v)
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("Failed fetched all URLs.")
	}
}
