package crypto

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

// 根据请求参数，获取结构体变量的hash值，并使用redis中的hash存储键值，
// 可作为接口幂等性的解决方案

type Message struct {
	Name string
	Age  int
}

// 获取结构体类型变量的hash值
func main() {
	md := md5.New()
	m := Message{"xiaoming", 28}

	bytes, _ := json.Marshal(m)

	md.Write(bytes)

	res := fmt.Sprintf("%X", md.Sum(nil)) // 大写十六进制

	fmt.Println(res)
}
