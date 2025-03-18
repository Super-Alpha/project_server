package main

import "fmt"

// 匿名函数
var add = func(x, y int) {
	fmt.Println(x + y)
}

// 闭包
func adder() func(y int) int {
	x := 10
	return func(y int) int {
		return x + y
	}
}

// defer
func f1() int {
	x := 10

	defer func() {
		x += 2
		fmt.Println(x)
	}()

	defer func() {
		x += 1
		fmt.Println(x)
	}()

	return x
}

func main() {
	// 通过变量调用匿名函数
	//add(10, 20)

	// 变量f引用了外部作用域中的变量x，此时f就是一个闭包，在f的生命周期内，x的值也一直有效
	//f := adder()
	//fmt.Println(f(20))

	fmt.Println(f1())
}
