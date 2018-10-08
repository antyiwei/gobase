package M09_proxy

import "testing"

func TestProxy(t *testing.T) {
	var sub Subject
	sub = &Proxy{}

	res := sub.Do()

	println(res)
	if res != "pre:real:after" {
		t.Fail()
	}
}
