package copy

import "fmt"

// User 定义User结构体
type User struct {
	Name string
	Age  int
}

// 定义一个全局的User
var globalUser = User{
	"xiaoming",
	28,
}

// modifyUser 定义一个函数，参数为User结构体“对象”，将全局globalUser指向传递过来的User结构体“对象”
func modifyUser(user User) {
	fmt.Printf("参数user的地址 = %p\n", &user)
	fmt.Printf("globalUser修改前的地址 = %p\n", &globalUser)
	fmt.Println("globalUser修改前 = ", globalUser)
	// 修改指向
	globalUser = user // 值拷贝
	fmt.Printf("globalUser修改后的地址 = %p\n", &globalUser)
	fmt.Println("globalUser修改后 = ", globalUser)
}

func main() {
	u := User{
		"xiaohong",
		29,
	}
	fmt.Printf("将要传递的参数u的地址 = %p\n", &u)
	modifyUser(u)
}
