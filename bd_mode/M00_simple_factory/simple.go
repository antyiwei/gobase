package M00_simple_factory

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
		panic("无效运算符号") // 这里强制程序奔溃
		return nil
	}

}
