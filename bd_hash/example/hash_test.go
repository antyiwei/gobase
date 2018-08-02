package example

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash/adler32"
	"hash/crc32"
	"hash/fnv"
	"testing"
)

func TestHashMd5(t *testing.T) {

	m := md5.New()
	m.Write([]byte("hello, world"))
	s := hex.EncodeToString(m.Sum(nil))
	println(s)
}
func TestHashFnv(t *testing.T) {

	a := fnv.New32()
	a.Write([]byte("hello, world"))
	fmt.Println(hex.EncodeToString(a.Sum(nil)))
}

func TestHash(t *testing.T) {

	pwd := "123456"
	sha1h := sha1.New()
	sha1h.Write([]byte(pwd))
	cipherStr := sha1h.Sum(nil)
	fmt.Println(cipherStr)
	fmt.Printf("%s\n", hex.EncodeToString(cipherStr)) // 输出加密结果

	{
		m := md5.New()
		m.Write([]byte("123456"))
		s := hex.EncodeToString(m.Sum(nil))
		println(s)
	}
	{
		a := fnv.New32()
		a.Write([]byte("123456"))
		fmt.Println(hex.EncodeToString(a.Sum(nil)))
	}

	{
		for _, v := range []string{"aaaaaaaaaa", "3333sdfsdffsdffsd", "234esrewr234324",
			`An Adler-32 checksum is obtained by calculating two 16-bit checksums A and B and concatenating their bits into a 32-bit integer. A is the sum of all bytes in the stream plus one, and B is the sum of the individual values of A from each step.
					At the beginning of an Adler-32 run, A is initialized to 1, B to 0. The sums are done modulo 65521 (the largest prime number smaller than 216). The bytes are stored in network order (big endian), B occupying the two most significant bytes.
					The function may be expressed as
					A = 1 + D1 + D2 + ... + Dn (mod 65521)
					 B = (1 + D1) + (1 + D1 + D2) + ... + (1 + D1 + D2 + ... + Dn) (mod 65521)
					   = n×D1 + (n−1)×D2 + (n−2)×D3 + ... + Dn + n (mod 65521)
					 Adler-32(D) = B × 65536 + A
					where D is the string of bytes for which the checksum is to be calculated, and n is the length of D.`} {
			calc(ADLER32, []byte(v))
			calc(CRC32, []byte(v))
		}
	}
}

var ADLER32 int = 0
var CRC32 int = 1

func calc(t int, b []byte) {
	var ret uint32
	if ADLER32 == t {
		ret = adler32.Checksum([]byte(b))
		fmt.Printf("ADLER32 %15d  : %s...  \n", ret, string(b[:5]))
	} else if CRC32 == t {
		ret = crc32.ChecksumIEEE([]byte(b))
		fmt.Printf("CRC32   %15d  : %s...  \n", ret, string(b[:5]))
	} else {
		return
	}
}
