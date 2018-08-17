package M04_builder

import "testing"

func TestBuilder1_GetResult(t *testing.T) {
	builder := &Builder1{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	if res != "123" {
		t.Fatal("Builder1 fail expect 123 acture %s", res)
	} else {
		t.Logf("res = %s", res)
	}
}

func TestBuilder2_GetResult(t *testing.T) {
	builder := &Builder2{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	if res != 6 {
		t.Fatal("Builder2 fail expect 6 acture %d", res)
	} else {
		t.Logf("res = %d", res)
	}
}
