package M17_iterator

import "testing"

func ExampleIterator() {

	var aggregate Aggregate

	aggregate = NewNumbers(1, 10)
	IteratorPrint(aggregate.Iterator())
}

func TestIteratorPrint(t *testing.T) {
	ExampleIterator()
}
