package main

import (
	"fmt"
	"unsafe"
)

type slice struct {
	array unsafe.Pointer // 指向数组的指针
	len   int            // 切片长度
	cap   int            // 底层数组容量
}

// 	s := make([]int, 5, 10) len为5（即可以使用下标s[0] ~ s[4]来操作里面的元素），capacity为10，表示后续向slice添加新的元素时可以不必重新分配内存，直接使用预留内存即可；

// 扩容：
// 1、如果当前切片的容量小于1024，则会将其容量扩展为原来的2倍
// 2、如果当前切片的容量大于1024，则会将其容量扩展为原来的1.25倍

// append追加元素：
//（1）假如Slice容量够用，则将新元素追加进去，Slice.len++，返回原Slice
//（2）原Slice容量不够，则将Slice先扩容，扩容后得到新Slice; 然后将新元素追加进新Slice，Slice.len++，返回新的Slice。

// 删除切片中的指定索引处的元素
func removeElement(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func one(s []int) { // 此时s与s1/s2指向同一内存地址

	s = append(s, 0) // 此时s发生扩容，与s1指向不同的内存地址,s=>[1,2,0]；（因s2其cap=4，所以此时不触发扩容机制，s与s2同指向[1,2,3,0]）

	for i := range s { // 此时s发生自增，不影响s1，所以s1仍为[1, 2], s=>[2,3,1]；（此时s=>[2,3,4,1]<=s2）
		s[i]++
	}
}

func two() {
	s1 := []int{1, 2}

	s2 := s1

	s2 = append(s2, 3) // 返回一个新的切片，此时s2=>[1,2,3],其len=3，cap=4

	one(s1) // 此时s1=>[1,2]

	one(s2) // 此时s2=>[1,2,3]

	fmt.Printf("%v,%v", s1, s2) // s1未被处理，所以s1=>[1 2];而s2此时len仍为3，所以s2=[2,3,4]
}

func main() {
	two()
}
