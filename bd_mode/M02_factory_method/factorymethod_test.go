package M02_factory_method

import (
	"testing"
)

// Compute 计算数据信息
func Compute(factory IFactory, a, b float64) float64 {
	oper := factory.CreateOperation()
	oper.SetA(a)
	oper.SetB(b)
	return oper.GetResult()

}

func TestAddOperationI(t *testing.T) {

	var (
		factory IFactory
	)

	factory = &AddFactory{}
	if r := Compute(factory, 1, 2); r != 3 {
		t.Fatal("error with factory method pattern")
	} else {
		t.Logf("a + b = %v", r)
	}

	factory = &SubFactory{}
	if r := Compute(factory, 6.5, 3.5); r != 3 {
		t.Fatal("error with factory method pattern")
	} else {
		t.Logf("a - b = %v", r)
	}

	factory = &MulFactory{}
	if r := Compute(factory, 1, 2); r != 2 {
		t.Fatal("error with factory method pattern")
	} else {
		t.Logf("a * b = %v", r)
	}

	factory = &DivFactory{}
	if r := Compute(factory, 1, 2); r != 0.5 {
		t.Fatal("error with factory method pattern")
	} else {
		t.Logf("a / b = %v", r)
	}
}
