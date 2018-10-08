package M19_strategy

import "testing"

func ExamplePayByCash() {
	ctx := NewPaymentContext("Ada", "", 123, &Cash{})
	ctx.Pay()
}

func ExamplePayByBank() {
	ctx := NewPaymentContext("Bob", "0002", 888, &Bank{})
	ctx.Pay()
}

func TestBank_Pay(t *testing.T) {
	ExamplePayByBank()
}

func TestCash_Pay(t *testing.T) {
	ExamplePayByCash()
}
