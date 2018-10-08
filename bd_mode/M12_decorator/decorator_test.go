package M12_decorator

import (
	"fmt"
	"testing"
)

func ExampleDecorator() {
	var c Component = &ConcreateComponent{}

	c = WarpAddDecorator(c, 10)
	c = WarpMulDecorator(c, 8)

	res := c.Calc()
	fmt.Printf("res %d\n", res)
}

func TestDecorator(t *testing.T) {
	ExampleDecorator()
}
