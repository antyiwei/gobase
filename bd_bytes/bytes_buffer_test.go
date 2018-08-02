package bytesbd

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestBytesBuffer(t *testing.T) {
	{
		result := bytes.Equal([]byte("a"), []byte("a"))
		fmt.Println(result)

	}

	{
		buffer := bytes.NewBuffer([]byte("helloworld"))

		fmt.Println(buffer.Len())
		for i := 0; i < buffer.Len(); i++ {
			fmt.Println(buffer)
			v, size, err := buffer.ReadRune()

			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Printf("位%d v=%v,size=%v,buffer.len()=%d\n", i+1, string(v), size, buffer.Len())

		}
	}
	{
		buffer := bytes.NewBuffer([]byte("helloworld"))

		fmt.Println(buffer.Len())
		for {
			v, size, err := buffer.ReadRune()
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}

			fmt.Printf(" v=%v,size=%v,buffer.len()=%d\n", string(v), size, buffer.Len())

		}
	}
}

/* Buffer是一个实现了读写方法的可变大小的字节缓冲。本类型的零值是一个空的可用于读写的缓冲。 */
func TestBytesBuffer2(t *testing.T) {

	var b1 bytes.Buffer
	b2 := bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	// 返回缓冲中未读取部分的字节长度；b.Len() == len(b.Bytes())。
	t.Log(b1.Len())

	//  Reset重设缓冲，因此会丢弃全部内容，等价于b.Truncate(0)。
	b2.Reset()
	t.Log(b2.String()) // Out:

	b2 = bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	t.Log(b2.String()) // Out:

	// 丢弃缓冲中除前n字节数据外的其它数据，如果n小于零或者大于缓冲容量将panic。
	b2.Truncate(b2.Len() - 2)
	t.Log(b2.String())
	// 必要时会增加缓冲的容量，以保证n字节的剩余空间。调用Grow(n)后至少可以向缓冲中写入n字节数据而无需申请内存。如果n小于零或者不能增加容量都会panic。

	b2.Grow(1000)
	t.Log(b2.String())

	//Read方法从缓冲中读取数据直到缓冲中没有数据或者读取了len(p)字节数据，将读取的数据写入p。返回值n是读取的字节数，除非缓冲中完全没有数据可以读取并写入p，此时返回值err为io.EOF；否则err总是nil。
	p := []byte("abcdef")
	n, err := b2.Read(p)
	if err != nil {
		t.Log(err.Error())
	}
	t.Logf("n:%d", n)
	t.Log(b2.String())

	t.Log(string(p))

	//返回未读取部分前n字节数据的切片，并且移动读取位置，就像调用了Read方法一样。如果缓冲内数据不足，会返回整个数据的切片。切片只在下一次调用b的读/写方法前才合法。
	b3 := b2.Next(2)
	t.Log(string(b3))

	t.Log(b2.String())

	// ReadByte读取并返回缓冲中的下一个字节。如果没有数据可用，返回值err为io.EOF。

}
