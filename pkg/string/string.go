package string

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 对于标准转换（存在内存拷贝）而言，强制转换直接替换指针的指向，避免了申请新内存（零拷贝）；针对高性能要求，可以采用强制转换的处理方式

/* 标准转换 string <-> []byte */

func bytes2String(b []byte) string {
	return string(b)
}

// string 转 []byte
func string2Bytes(s string) (b []byte) {
	return []byte(s)
}

/*
强制转换 string <-> []byte
*/

// []byte 转 string
func bytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// string 转 []byte
func stringToBytes(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len

	return b
}

func main() {
	str := "hello"
	fmt.Println(bytesToString([]byte(str)))
}
