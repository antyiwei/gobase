package M23_chain_of_responsibility

import "testing"

func ExampleChain() {

	c1 := NewProjectManagerChain()
	c2 := NewDepManagerChain()
	c3 := NewGeneraManagerChain()

	c1.SetSuccess(c2)
	c2.SetSuccess(c3)

	var c Manager = c1

	c.HandleFeeRequest("bob", 400)
	c.HandleFeeRequest("tom", 1400)
	c.HandleFeeRequest("ada", 10000)
	c.HandleFeeRequest("floar", 400)
}

func TestRequestChain_HaveRight(t *testing.T) {
	ExampleChain()
}
