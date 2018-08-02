package example

import (
	"bytes"
	"fmt"
	"sort"
	"testing"
	"unicode"
)

/* MinRead是被Buffer.ReadFrom传递给Read调用的最小尺寸。只要该Buffer在保存内容之外有最少MinRead字节的余量，其ReadFrom方法就不会增加底层的缓冲。 */
func TestBytesConstants(t *testing.T) {
	t.Log(bytes.MinRead) // out:512
}

/* 如果内存中不能申请足够保存数据的缓冲，ErrTooLarge就会被传递给panic函数。 */
func TestBytesVariables(t *testing.T) {
	t.Log(bytes.ErrTooLarge) // out: bytes.Buffer: too large
}

/* Compare函数返回一个整数表示两个[]byte切片按字典序比较的结果（类同C的strcmp）。如果a==b返回0；如果a<b返回-1；否则返回+1。nil参数视为空切片。 */
func TestBytesCompare(t *testing.T) {
	a := []byte("a")
	b := a
	t.Log(bytes.Compare(a, b)) // out :0

	//
	{
		// Interpret Compare's result by comparing it to zero.
		var a, b []byte
		if bytes.Compare(a, b) < 0 {
			// a less b
		}
		if bytes.Compare(a, b) <= 0 {
			// a less or equal b
		}
		if bytes.Compare(a, b) > 0 {
			// a greater b
		}
		if bytes.Compare(a, b) >= 0 {
			// a greater or equal b
		}
		// Prefer Equal to Compare for equality comparisons.
		if bytes.Equal(a, b) {
			// a equal b
			t.Log(bytes.Equal(a, b)) // out:true
		}
		if !bytes.Equal(a, b) {
			// a not equal b
			t.Log(!bytes.Equal(a, b))
		}
	}
	{
		// Binary search to find a matching byte slice.
		var needle []byte
		var haystack [][]byte // Assume sorted
		i := sort.Search(len(haystack), func(i int) bool {
			// Return haystack[i] >= needle.
			return bytes.Compare(haystack[i], needle) >= 0
		})
		if i < len(haystack) && bytes.Equal(haystack[i], needle) {
			// Found it!
		}
	}
}

/* 判断两个切片的内容是否完全相同。 */
func TestBytesEqual(t *testing.T) {
	var a, b []byte
	t.Log(bytes.Equal(a, b))
	t.Log(t.Name())
}

/* 判断两个utf-8编码切片（将unicode大写、小写、标题三种格式字符视为相同）是否相同。 */
func TestBytesEqualFold(t *testing.T) {
	var a, b []byte
	a = []byte("abc")
	b = []byte("ddd")
	t.Log(bytes.Equal(a, b))
}

/* Runes函数返回和s等价的[]rune切片。（将utf-8编码的unicode码值分别写入单个rune） */
func TestBytesRunes(t *testing.T) {

	var a []byte
	a = []byte("abc")

	t.Log(bytes.Runes(a))

}

/* 判断s是否有前缀切片prefix。 */
func TestBytesHasPrefix(t *testing.T) {
	var a, b, c []byte
	a = []byte("qwerty")
	b = []byte("q")
	t.Log(bytes.HasPrefix(a, b))

	//
	c = []byte("你好")
	t.Log(c)
	t.Log(bytes.HasPrefix(c, []byte("你")))
	t.Log([]byte("好"))
}

/* 判断s是否有后缀切片suffix。 */
func TestBytesHasSuffix(t *testing.T) {
	var a, b, c []byte
	a = []byte("qwerty")
	b = []byte("A")
	t.Log(bytes.HasSuffix(a, b))

	//
	c = []byte("你好")
	t.Log(c)
	t.Log(bytes.HasSuffix(c, []byte("好")))
	t.Log([]byte("好"))
}

/* 判断切片b是否包含子切片subslice。 */
func TestBytesContains(t *testing.T) {
	var a, b []byte
	a, b = []byte("你好"), []byte("好")
	t.Log(bytes.HasSuffix(a, b))
}

/* Count计算s中有多少个不重叠的sep子切片。 */
func TestBytesCount(t *testing.T) {
	var a, b []byte
	a, b = []byte("abdcdabafafabjwrhfjasfa"), []byte("ab")
	t.Log(bytes.Count(a, b))
}

/* 子切片sep在s中第一次出现的位置，不存在则返回-1。 */
func TestBytesIndex(t *testing.T) {
	var s, sep []byte
	s, sep = []byte("sdsssaaadbddsaf"), []byte("a")
	t.Log(bytes.Index(s, sep))

}

/* 字符c在s中第一次出现的位置，不存在则返回-1。 */
func TestBytesIndexByte(t *testing.T) {

	var s, _ []byte
	s, _ = []byte("sdsssaaadbddsaf"), []byte("a")

	t.Log(bytes.IndexByte(s, 'd'))
}

/* unicode字符r的utf-8编码在s中第一次出现的位置，不存在则返回-1。 */
func TestBytesIndexRune(t *testing.T) {
	var s, _ []byte
	s = []byte("sdsssaaadbddsaf")

	t.Log(bytes.IndexRune(s, '好'))
}

/* s中第一个满足函数f的位置i（该处的utf-8码值r满足f(r)==true），不存在则返回-1  */
func TestBytesIndexFunc(t *testing.T) {
	var s, _ []byte
	s = []byte("sdsssaaoadbddsaf")
	result := bytes.IndexFunc(s, func(a rune) bool {
		if a == 'o' {
			return true
		} else {
			return false
		}
	})
	t.Log(result)

}

/* s中最后一个满足函数f的unicode码值的位置i，不存在则返回-1。 */
func TestBytesLastIndexFunc(t *testing.T) {
	var s, _ []byte
	s = []byte("sdsssaaoadbddsaof")
	result := bytes.LastIndexFunc(s, func(a rune) bool {
		if a == 'o' {
			return true
		} else {
			return false
		}
	})
	t.Log(result)

}

/* 返回s中每个单词的首字母都改为标题格式的拷贝。BUG: Title用于划分单词的规则不能很好的处理Unicode标点符号。 */
func TestBytesTitle(t *testing.T) {
	var s []byte
	s = []byte("abd")
	sbyte := bytes.Title(s)
	t.Log(sbyte)
	t.Log(string(sbyte))
}

/* 返回将所有字母都转为对应的小写版本的拷贝。 */
func TestBytesToLower(t *testing.T) {
	var s []byte
	s = []byte("AEFSsdfDFA")
	b := bytes.ToLower(s)
	t.Log(b)
	t.Log(string(b))
}

/* 使用_case规定的字符映射，返回将所有字母都转为对应的小写版本的拷贝。 */
func TestBytesToLowerSpecial(t *testing.T) {
	lower := "abcçdefgğhıijklmnoöprsştuüvyz"
	// upper := "ABCÇDEFGĞHIİJKLMNOÖPRSŞTUÜVYZ"
	// u := ToUpperSpecial(unicode.TurkishCase, upper)
	// if u != upper {
	// 	t.Errorf("Upper(upper) is %s not %s", u, upper)
	// }
	// u = ToUpperSpecial(unicode.TurkishCase, lower)
	// if u != upper {
	// 	t.Errorf("Upper(lower) is %s not %s", u, upper)
	// }
	l := bytes.ToLowerSpecial(unicode.TurkishCase, []byte(lower))

	t.Errorf("Lower(lower) is %s not %s", l, lower)
}

/* 返回将所有字母都转为对应的大写版本的拷贝。 */
func TestBytesToUpper(t *testing.T) {
	lower := "abcçdefgğhıijklmnoöprsştuüvyz"
	s := []byte(lower)
	t.Log(string(bytes.ToUpper(s)))
}

/* 返回将s中前n个不重叠old切片序列都替换为new的新的切片拷贝，如果n<0会替换所有old子切片。 */
func TestBytesReplace(t *testing.T) {
	var s, old, new []byte
	s = []byte("abdcdefjhijklmn")
	old = []byte("c")
	new = []byte("C")
	result := bytes.Replace(s, old, new, -2)
	t.Log(result)
	t.Log(string(result))
}

/* 将一系列[]byte切片连接为一个[]byte切片，之间用sep来分隔，返回生成的新切片。 */
func TestBytesJoin(t *testing.T) {
	s := [][]byte{[]byte("你好"), []byte("世界")}
	sep := []byte(",")
	t.Log(string(bytes.Join(s, sep)))
	sep = []byte("#")
	t.Log(string(bytes.Join(s, sep)))
}

/* 返回count个b串联形成的新的切片。 */
func TestBytesRepeat(t *testing.T) {
	var s []byte
	s = []byte("abc")
	t.Log(string(bytes.Repeat(s, 7)))
}

/* 返回将s前后端所有cutset包含的unicode码值都去掉的子切片。（共用底层数组） */
func TestBytesTrimMethod(t *testing.T) {

	var s []byte
	s = []byte("abcdef kkafjaio qreqwrn ndsfakabc")

	t.Log(string(bytes.Trim(s, "abc"))) // def kkafjaio qreqwrn ndsfakabc
	t.Log(string(bytes.TrimSpace(s)))   //abcdef kkafjaio qreqwrn ndsfakabc

	t.Log(string(bytes.TrimFunc(s, func(r rune) bool {
		if r == 34 {
			return true
		} else {
			return false
		}
	})))
	t.Log(string(bytes.TrimLeft(s, "abc"))) // 左边abc删除

	var b = []byte("Goodbye,, world!")
	b = bytes.TrimPrefix(b, []byte("Goodbye,"))
	b = bytes.TrimPrefix(b, []byte("See ya,"))
	fmt.Printf("Hello%s", b)
}

func TestBytesFields(t *testing.T) {
	var s []byte
	s = []byte("sfslfjk")

	t.Log(bytes.Fields(s))

	lena := len(bytes.Fields(s))
	for i := 0; i < lena; i++ {

	}
}
