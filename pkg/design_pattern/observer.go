package design_pattern

import "fmt"

/*
 * @Description:观察者模式（发布订阅模式）一个对象状态发生改变，则所有依赖于它的对象都会得到通知并自动更新
 */

// 抽象层 被观察者
type Subject interface {
	Register(observer Observer)
	Remove(observer Observer)
	Notify()
}

// 观察者
type Observer interface {
	Update()
}

// 实现层
type ConcreteSubject struct {
	observers []Observer
}

func (c *ConcreteSubject) Register(observer Observer) {
	c.observers = append(c.observers, observer)
}
func (c *ConcreteSubject) Remove(observer Observer) {
	for i, v := range c.observers {
		if v == observer {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			break
		}
	}
}
func (c *ConcreteSubject) Notify() {
	for _, v := range c.observers {
		v.Update()
	}
}

type ConcreteObserver struct {
	Count int
}

func (c *ConcreteObserver) Update() {
	c.Count += 1
	fmt.Printf("count: %d\n", c.Count)
}

func main() {
	subject := &ConcreteSubject{}
	observer := &ConcreteObserver{}
	subject.Register(observer)
	subject.Notify()
}
