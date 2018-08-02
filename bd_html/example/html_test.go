package example

import (
	"html"
	"testing"
)

func TestHtml(t *testing.T) {
	s := "a>cl'kll"
	var str string
	if html.UnescapeString(html.EscapeString(s)) == s {
		t.Log(html.EscapeString(s))
		str = html.EscapeString(s)
	}

	t.Log(html.UnescapeString(str))
}
