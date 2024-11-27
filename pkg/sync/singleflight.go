package main

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"golang.org/x/sync/singleflight"
)

/*
golang/sync/singleflight.Group 是 Go 语言扩展包中提供了另一种同步原语，它能够在一个服务中抑制对下游的多次重复请求。
一个比较常见的使用场景是：我们在使用 Redis 对数据库中的数据进行缓存，发生缓存击穿时，大量的流量都会打到数据库上进而影响服务的尾延时。
*/

var errorNotExist = errors.New("not exist")
var gsf singleflight.Group

// 获取数据
func getData(key string) (string, error) {
	data, err := getDataFromCache(key)
	if errors.Is(err, errorNotExist) {
		//模拟从db中获取数据
		data, err = getDataFromDB(key)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		//TOOD: set cache
	} else if err != nil {
		return "", err
	}
	return data, nil
}

// 获取数据
func getDatas(key string) (string, error) {
	data, err := getDataFromCache(key)
	if errors.Is(err, errorNotExist) {
		//模拟并发从db中获取数据,确保相同key的请求中，只有一个会请求成功(若某key请求已经执行，则后续不再执行该key请求)
		v, err, shared := gsf.Do(key, func() (interface{}, error) {
			return getDataFromDB(key)
		})
		if err != nil {
			log.Println(err)
			return "", err
		}
		if shared {
			data = v.(string)
		} else {
			return "", errors.New("shares is false")
		}
	} else if err != nil {
		return "", err
	}
	return data, nil
}

// 模拟从cache中获取值，cache中无该值
func getDataFromCache(key string) (string, error) {
	return "", errorNotExist
}

// 模拟从数据库中获取值
func getDataFromDB(key string) (string, error) {
	fmt.Printf("get %s from database", key)
	return "data", nil
}

func main() {
	var wg sync.WaitGroup
	//模拟10个并发，请求缓存数据
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//data,err := getData("key") // 会导致缓存击穿
			data, err := getDatas("key") // 会避免缓存击穿
			if err != nil {
				fmt.Print(err)
				return
			}
			fmt.Println(data)
		}()
	}
	wg.Wait()
}
