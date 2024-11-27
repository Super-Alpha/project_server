package design_pattern

import "fmt"

// 适配器模式：将一个类的接口转换成用户希望的另外一个接口。使得原本由于接口不兼容而不能一起工作的类可以一起工作。

// 适配的目标（适配器实现该接口）
type V5 interface {
	Use5V()
}

// 业务类,依赖V5接口
type Phones struct {
	v V5 // 适配器对象
}

func NewPhone(v V5) *Phones {
	return &Phones{v}
}

func (p *Phones) Charge() {
	fmt.Println("Phone进行充电...")
	p.v.Use5V()
}

// 被适配的角色，适配者
type V220 struct{}

func (v *V220) Use220V() {
	fmt.Println("使用220V的电压")
}

// 适配器
type Adapter struct {
	v220 *V220 // 被适配对象作为内部属性
}

func (a *Adapter) Use5V() {
	fmt.Println("使用适配器进行充电")
	//调用适配者的方法
	a.v220.Use220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}

// ------- 业务逻辑层 -------
func mains() {
	iphone := NewPhone(NewAdapter(new(V220)))

	iphone.Charge()
}
