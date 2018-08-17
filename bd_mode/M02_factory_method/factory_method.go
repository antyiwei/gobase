package M02_factory_method

/* ======================== golang 中工厂方法模式 ======================== */

// Operation 是OperstionI 接口实现的基类，封装公用方法
type Operation struct {
	a, b float64
}

// OperstionI 是被封装的实际类接口
type OperstionI interface {
	GetResult() float64
	SetA(float64)
	SetB(float64)
}

//SetA 设置 A
func (op *Operation) SetA(a float64) {
	op.a = a
}

//SetB 设置 B
func (op *Operation) SetB(b float64) {
	op.b = b
}

// AddOperation  OperstionI 的实际加法实现
type AddOperation struct {
	Operation
}

// Result 获取结果
func (this *AddOperation) GetResult() float64 {
	return this.a + this.b
}

// SubOperation  OperstionI 的实际减法实现
type SubOperation struct {
	Operation
}

// Result 获取结果
func (this *SubOperation) GetResult() float64 {
	return this.a - this.b
}

// MulOperation  OperstionI 的实际乘法实现
type MulOperation struct {
	Operation
}

// Result 获取结果
func (this *MulOperation) GetResult() float64 {
	return this.a * this.b
}

// DivOperation  OperstionI 的实际除法实现
type DivOperation struct {
	Operation
}

// Result 获取结果
func (this *DivOperation) GetResult() float64 {
	if this.b == 0 {
		return 0
	}
	return this.a / this.b
}

// IFactory  是工厂接口
type IFactory interface {
	CreateOperation() OperstionI
}

// AddFactory 是 AddOperation 的工厂类
type AddFactory struct {
}

func (this *AddFactory) CreateOperation() OperstionI {
	return &(AddOperation{})
}

// SubFactory 是 SubOperation 的工厂类
type SubFactory struct {
}

func (this *SubFactory) CreateOperation() OperstionI {
	return &(SubOperation{})
}

// MulFactory 是 MulOperation 的工厂类
type MulFactory struct {
}

func (this *MulFactory) CreateOperation() OperstionI {
	return &(MulOperation{})
}

// DivFactory 是 DivOperation 的工厂类
type DivFactory struct {
}

func (this *DivFactory) CreateOperation() OperstionI {
	return &(DivOperation{})
}
