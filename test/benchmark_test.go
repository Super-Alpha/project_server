package main

import (
	"fmt"
	"strconv"
	"testing"
)

/*
一、benchmark基准测试：在一定时间(默认1秒)内重复运行测试代码，然后输出执行次数和内存分配结果。
1、go test -bench . 来运行基准测试（需要先进入该测试文件路径之下）
2、go test -bench. -benchmem -benchtime=5s -count=3 执行3次，运行5秒来运行基准测试，并输出每次执行内存分配结果、每次执行耗费时间、每次执行内存分配次数
*/

func BenchmarkStrconv(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = strconv.Itoa(n)
	}
}

func BenchmarkFmtSprint(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = fmt.Sprint(n)
	}
}
