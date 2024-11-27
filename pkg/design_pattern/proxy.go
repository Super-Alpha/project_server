package design_pattern

import "fmt"

// 代理模式：为其他对象提供一种代理以控制对这个对象的访问

// 以海外代购为例

// --------抽象层--------
type Goods struct {
	Kind string // 种类
	Fact bool   // 真伪
}

type Shopping interface {
	Buy(goods *Goods)
}

// ------实现层-------
type KoreaShopping struct{}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国进行了购物, 买了 ", goods.Kind)
}

type AmericanShopping struct{}

func (as *AmericanShopping) Buy(goods *Goods) {
	fmt.Println("去美国进行了购物, 买了 ", goods.Kind)
}

type AfricaShopping struct{}

func (as *AfricaShopping) Buy(goods *Goods) {
	fmt.Println("去非洲进行了购物, 买了 ", goods.Kind)
}

// OverseasProxy 海外代购实例对象
type OverseasProxy struct {
	shopping Shopping
}

// Buy 代理实例：主体功能，调用被代理实例方法的同时，加以调用辅助功能
func (op *OverseasProxy) Buy(goods *Goods) {
	if op.distinguish(goods) == true {
		op.shopping.Buy(goods)
		op.check(goods)
	}
}

// 代理辅助功能：辨别真伪
func (op *OverseasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对[", goods.Kind, "]进行了辨别真伪.")
	if goods.Fact == false {
		fmt.Println("发现假货", goods.Kind, ", 不应该购买。")
	}
	return goods.Fact
}

// 代理辅助功能：海关检查
func (op *OverseasProxy) check(goods *Goods) {
	fmt.Println("对[", goods.Kind, "] 进行了海关检查， 成功的带回祖国")
}

// NewProxy 创建代理
func NewProxy(shopping Shopping) Shopping {
	return &OverseasProxy{shopping}
}

// -------业务逻辑层--------
func main() {
	koreaG1 := Goods{
		Kind: "韩国面膜",
		Fact: true,
	}

	koreaG2 := Goods{
		Kind: "韩国泡菜",
		Fact: false,
	}

	// 代购韩国商品的海外代理
	overseasProxy := NewProxy(new(KoreaShopping))

	overseasProxy.Buy(&koreaG1)

	overseasProxy.Buy(&koreaG2)
}
