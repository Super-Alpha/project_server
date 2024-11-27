package design_pattern

//设计一个电脑主板架构，电脑包括（显卡，内存，CPU）3个固定的插口，显卡具有显示功能（display，功能实现只要打印出意义即可），内存具有存储功能（storage），cpu具有计算功能（calculate）。
//现有Intel厂商，nvidia厂商，Kingston厂商，均会生产以上三种硬件。
//要求组装两台电脑，
//1台（Intel的CPU，Intel的显卡，Intel的内存）
//1台（Intel的CPU， nvidia的显卡，Kingston的内存）
//用抽象工厂模式实现。

type AbstractComputerFactory interface {
	CreateGPU() AbstractGPU
	CreateSTO() AbstractSTO
	CreateCPU() AbstractSTO
}

type AbstractGPU interface {
	Display()
}

type AbstractSTO interface {
	Storage()
}

type AbstractCPU interface {
	Calculate()
}

type IntelGpu struct{}

func (g *IntelGpu) Display() {}

type IntelSto struct{}

func (s *IntelSto) Storage() {}

type IntelCpu struct{}

func (c *IntelCpu) Calculate() {}

type NvidiaGpu struct{}

func (g *NvidiaGpu) Display() {}

type NvidiaSto struct{}

func (s *NvidiaSto) Storage() {}

type NvidiaCpu struct{}

func (c *NvidiaCpu) Calculate() {}

type KingstonGpu struct{}

func (g *KingstonGpu) Display() {}

type KingstonSto struct{}

func (s *KingstonSto) Storage() {}

type KingstonCpu struct{}

func (c *KingstonCpu) Calculate() {}

type IntelComputeFactory struct{}

func (i *IntelComputeFactory) CreateGPU() AbstractGPU {
	var gpu AbstractGPU
	gpu = new(IntelGpu)

	return gpu
}

func (i *IntelComputeFactory) CreateSTO() AbstractSTO {
	var sto AbstractSTO
	sto = new(IntelSto)
	return sto
}

func (i *IntelComputeFactory) CreateCPU() AbstractCPU {
	var cpu AbstractCPU
	cpu = new(IntelCpu)
	return cpu
}

type NvidiaComputeFactory struct{}

func (i *NvidiaComputeFactory) CreateGPU() AbstractGPU {
	var gpu AbstractGPU
	gpu = new(NvidiaGpu)
	return gpu
}

func (i *NvidiaComputeFactory) CreateSTO() AbstractSTO {
	var sto AbstractSTO
	sto = new(NvidiaSto)
	return sto
}

func (i *NvidiaComputeFactory) CreateCPU() AbstractCPU {
	var cpu AbstractCPU
	cpu = new(NvidiaCpu)
	return cpu
}

type KingstonComputeFactory struct{}

func (i *KingstonComputeFactory) CreateGPU() AbstractGPU {
	var gpu AbstractGPU
	gpu = new(KingstonGpu)
	return gpu
}

func (i *KingstonComputeFactory) CreateSTO() AbstractSTO {
	var sto AbstractSTO
	sto = new(KingstonSto)
	return sto
}

func (i *KingstonComputeFactory) CreateCPU() AbstractCPU {
	var cpu AbstractCPU
	cpu = new(KingstonCpu)
	return cpu
}

type Computer struct {
	Gpu AbstractGPU
	Cpu AbstractCPU
	Sto AbstractSTO
}

func (c *Computer) ComFunc() {
	c.Cpu.Calculate()
	c.Sto.Storage()
	c.Gpu.Display()
}

func main() {
	intelFac := new(IntelComputeFactory)
	nvidiaFac := new(NvidiaComputeFactory)
	kingsFac := new(KingstonComputeFactory)

	computer1 := new(Computer)
	computer1.Gpu = intelFac.CreateGPU()
	computer1.Cpu = intelFac.CreateCPU()
	computer1.Sto = intelFac.CreateSTO()

	computer2 := new(Computer)
	computer2.Gpu = nvidiaFac.CreateGPU()
	computer2.Cpu = intelFac.CreateCPU()
	computer2.Sto = kingsFac.CreateSTO()
}
