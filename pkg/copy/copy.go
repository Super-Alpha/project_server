package copy

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
1、深拷贝（Deep Copy）
拷贝的是数据本身，创造一个样的新对象，新创建的对象与原对象不共享内存，新创建的对象在内存中开辟一个新的内存地址，新对象值修改时不会影响原对象值。既然内存地址不同，释放内存地址时，可分别释放

值类型的数据，默认全部都是深复制，Array、Int、String、Struct、Float，Bool
*/

/*2、浅拷贝（Shallow Copy）
拷贝的是数据地址，只复制指向的对象的指针，此时新对象和老对象指向的内存地址是一样的，新对象值修改时老对象也会变化。释放内存地址时，同时释放内存地址

引用类型的数据，默认全部都是浅复制，Slice，Map
*/

// 本质区别: 是否真正获取（复制）对象实体，而不是引用

type Person struct {
	Name string
	Age  int
}

// 深拷贝（值类型数据）DeepCopy

func DeepCopy() {
	person1 := Person{
		Name: "wang",
		Age:  18,
	}

	person2 := person1

	person2.Age = 20

	//fmt.Printf("%p\n", &person1)
	//fmt.Printf("%p\n", &person2)

	fmt.Println(person1)
	fmt.Println(person2)
}

// 浅拷贝（引用类型数据）ShallowCopy

func ShallowCopy() {
	p1 := []int{1}
	p2 := p1

	fmt.Printf("%p\n", &p1)
	fmt.Printf("%p\n", &p2)

	fmt.Println("person1", (*reflect.SliceHeader)(unsafe.Pointer(&p1))) // 只能区分切片类型的地址
	fmt.Println("person2", (*reflect.SliceHeader)(unsafe.Pointer(&p2)))
}
