package design_pattern

import "sync"

// 单例模式

type Singleton struct{}

var singleton *Singleton

func GetInstance() *Singleton {

	var once sync.Once

	if singleton == nil {
		// 解决并发创建单例的情况
		once.Do(func() {
			singleton = &Singleton{}
		})
	}

	return singleton
}
