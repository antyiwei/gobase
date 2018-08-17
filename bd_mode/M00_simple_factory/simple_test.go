package M00_simple_factory

import (
	"fmt"
	"testing"
)

// test add
func TestType1(t *testing.T) {

	operator := NewOperateFactory().CreateOperate("+")

	r := operator.Operate(1, 2)
	fmt.Printf("add result is %d\n", r)
	if r != 3 {
		t.Fatal("Type1 test fail")
	} else {
		t.Log("Type1 test success")
	}

}

// 测试Multiple
func TestType2(t *testing.T) {

	operator := NewOperateFactory().CreateOperate("*")

	r := operator.Operate(1, 2)
	fmt.Printf("multiple result is %d\n", r)
	if r != 3 {
		t.Fatal("Type1 test fail")
	} else {
		t.Log("Type1 test success")
	}

}

// 测试其他命令函数
func TestTypeOther(t *testing.T) {

	operator := NewOperateFactory().CreateOperate("-")

	r := operator.Operate(1, 2)
	fmt.Printf("multiple result is %d\n", r)
	if r != 3 {
		t.Fatal("Type1 test fail")
	} else {
		t.Log("Type1 test success")
	}

}
