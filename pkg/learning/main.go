package main

/*
#include <string.h>
#include <stdlib.h>  // 引入 free 函数
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	str := C.CString("hello")
	defer C.free(unsafe.Pointer(str))
	fmt.Println(C.strlen(str))
}
