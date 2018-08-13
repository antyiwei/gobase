package factory

type Operater interface {
	Operate(int, int) int
}

type AddOperate struct {
}

func (this *AddOperate) Operate(rhs int, lhs int) int {
	return rhs + lhs
}

type MultipleOperate struct {
}

func (this *MultipleOperate) Operate(rhs int, lhs int) int {
	return rhs * lhs
}

type OperateFactory struct {
}

func NewOperateFactory() *OperateFactory {
	return &OperateFactory{}
}

func (this *OperateFactory) CreateOperate(operatename string) Operater {
	switch operatename {
	case "+":
		return &AddOperate{}
	case "*":
		return &MultipleOperate{}
	default:
		panic("无效运算符号")
		return nil
	}

}

/* ======================== golang 中工厂方法模式 ======================== */
type Operation struct {
	a, b float64
}

type OperstionI interface {
	GetResult() float64
	SetA(float64)
	SetB(float64)
}

func (op *Operation) SetA(a float64) {
	op.a = a
}

func (op *Operation) SetB(b float64) {
	op.b = b
}

type AddOperation struct {
	Operation
}

func (this *AddOperation) GetResult() float64 {
	return this.a + this.b
}

type SubOperation struct {
	Operation
}

func (this *SubOperation) GetResult() float64 {
	return this.a - this.b
}

type MulOperation struct {
	Operation
}

func (this *MulOperation) GetResult() float64 {
	return this.a * this.b
}

type DivOperation struct {
	Operation
}

func (this *DivOperation) GetResult() float64 {
	if this.b == 0 {
		return 0
	}
	return this.a / this.b
}

type IFactory interface {
	CreateOperation() Operation
}
type AddFactory struct {
}

func (this *AddFactory) CreateOperation() OperstionI {
	return &(AddOperation{})
}

type SubFactory struct {
}

func (this *SubOperation) CreateOperation() OperstionI {
	return &(SubOperation{})
}

type MulFactory struct {
}

func (this *MulFactory) CreateOperation() OperstionI {
	return &(MulOperation{})
}

type DivFactory struct {
}

func (this *DivFactory) CreateOperation() OperstionI {
	return &(DivOperation{})
}

/* ======================== golang 中抽象工厂模式 ======================== */
type GirlFriend struct {
	nationality string
	eyesColor   string
	language    string
}

type AbstractFactory interface {
	CreateMyLove() GirlFriend
}

type IndianGirlFriendFactory struct {
}

type KoreanGirlFriendFactory struct {
}

func (a IndianGirlFriendFactory) CreateMyLove() GirlFriend {
	return GirlFriend{"Indian", "Black", "Hindi"}
}

func (a KoreanGirlFriendFactory) CreateMyLove() GirlFriend {
	return GirlFriend{"Korean", "Brown", "Korean"}
}

func GetGirlFriend(typeGf string) GirlFriend {
	var gffact AbstractFactory
	switch typeGf {
	case "Indian":
		gffact = IndianGirlFriendFactory{}
		return gffact.CreateMyLove()
	case "Korean":
		gffact = KoreanGirlFriendFactory{}
		return gffact.CreateMyLove()
	}
	return GirlFriend{}
}
