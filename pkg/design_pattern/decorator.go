package design_pattern

import "fmt"

// 装饰器模式：通过将对象放入特殊封装对象(别名:装饰器)中来为原始对象增加新的行为;

// 以手机贴膜和安装手机壳为例

// ---------- 抽象层 ----------
type Phone interface {
	Show()
}

// ----------- 实现层 -----------
type HuaWei struct{}

func (hw *HuaWei) Show() {
	fmt.Println("秀出了HuaWei手机")
}

type XiaoMi struct{}

func (xm *XiaoMi) Show() {
	fmt.Println("秀出了XiaoMi手机")
}

// 具体的装饰器类
type MoDecorator struct {
	Phone
}

func (md *MoDecorator) Show() {
	md.Phone.Show()      //调用被装饰构件的原方法
	fmt.Println("贴膜的手机") //装饰额外的方法
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{phone}
}

type KeDecorator struct {
	Phone
}

func (kd *KeDecorator) Show() {
	kd.Phone.Show()       //调用被装饰构件的原方法
	fmt.Println("手机壳的手机") //装饰额外的方法
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{phone}
}

// ------------ 业务逻辑层 ---------
func main() {
	//获得裸机实例
	var huawei Phone
	huawei = new(HuaWei)
	huawei.Show()

	fmt.Println("---------")
	//用贴膜装饰器装饰裸机，得到贴膜新功能构件
	var moHuawei Phone
	moHuawei = NewMoDecorator(huawei) //通过HueWei ---> MoHuaWei
	moHuawei.Show()                   //调用装饰后新构件的方法

	fmt.Println("---------")
	//用手机壳装饰器装饰裸机，得到手机壳新功能构件
	var keHuawei Phone
	keHuawei = NewKeDecorator(huawei) //通过HueWei ---> KeHuaWei
	keHuawei.Show()

	fmt.Println("---------")
	//用手机壳装饰器装饰贴膜手机，得到贴膜手机壳新功能特性
	var keMoHuaWei Phone
	keMoHuaWei = NewMoDecorator(keHuawei) //通过KeHuaWei ---> KeMoHuaWei
	keMoHuaWei.Show()
}
