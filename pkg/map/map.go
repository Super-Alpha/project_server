package _map

import (
	"fmt"
	"runtime"
)

// Map（采用拉链法解决的hash冲突），并不是并发安全的，不可以边读边写
// 1、读取数据：
//（1）根据 key取hash 值；
//（2）根据hash值的最后B位（B为桶数组长度的指数），确定所在的桶；
//（3）然后根据hash值的高8位，确定在桶中的位置，若在原桶中没有找到目标值，则沿着桶链表依次遍历溢出桶，基于hash值的高8位查找目标值；
//（4）若命中相同的key，则返回value；倘若key不存在，则返回零值.

// 2、写入数据
// （1）根据key获取hash值；
//（2）根据hash值的最后B位（B为桶数组长度的指数），确定所在的桶；
//（3）然后根据hash值的高8位，确定在桶中的位置，若在原桶中没有找到目标值，则沿着桶链表依次遍历溢出桶，基于hash值的高8位查找目标值；若存在相同的key-value对，则替换value；
//     若不存在，则添加，添加时，若桶中数量超过8个，则将其置于溢出桶中，当溢出桶越来越多，装载因子(元素数量/桶数量 > 6.5时)也会逐渐增大，则触发扩容机制，进行渐进式扩容。

/*
map 扩容机制的核心点包括：
（1）扩容分为增量扩容和等量扩容；

（2）当桶内 key-value 总数/桶数组长度 > 6.5 时发生‘增量扩容’，新桶数组长度增长为原值的两倍；（例如：原来数组长度为5，扩容后数组长度为10）

（3）当桶内溢出桶数量大于等于 2^B 时(B为桶数组长度的指数，B最大取15)，发生‘等量扩容’，新桶的长度保持为原值；

（4）采用渐进扩容的方式，当桶被实际操作时（即删除或增加元素时），由使用者负责完成数据迁移（将旧桶中的数据搬迁到新桶中，并非一次性搬移，少量多次），避免因为一次性的全量数据迁移引发性能抖动.
*/

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/(1024*1024))
}

func main() {
	n := 1000000
	var m = make(map[int][128]byte)

	printAlloc()

	for i := 0; i < n; i++ { // Adds 1 million elements
		m[i] = [128]byte{}
	}
	printAlloc()

	for i := 0; i < n; i++ { // Deletes 1 million elements
		delete(m, i)
	}

	runtime.GC() // Triggers a manual(手动的) GC

	printAlloc()

	runtime.KeepAlive(m) // Keeps a reference(引用) to m, so that the map isn’t collected,确保不被释放掉
}
