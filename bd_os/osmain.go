package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// {
	// 	b, err := ioutil.ReadFile("test.log")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(b)
	// 	str := string(b)
	// 	fmt.Println(str)
	// }
	// {
	// 	d1 := []byte("hello\ngo\n")
	// 	err := ioutil.WriteFile("test.txt", d1, 0644)
	// 	check(err)
	// }
	// {
	// 	path := "test.txt"
	// 	fi, err := os.Open(path)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer fi.Close()
	// }

	// {

	// 	f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0755)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if err := f.Close(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// {

	// 	f, err := os.Open("test.log")
	// 	check(err)

	// 	b1 := make([]byte, 5)
	// 	n1, err := f.Read(b1)
	// 	check(err)
	// 	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// 	o2, err := f.Seek(6, 0)
	// 	check(err)
	// 	b2 := make([]byte, 2)
	// 	n2, err := f.Read(b2)
	// 	check(err)
	// 	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	// 	o3, err := f.Seek(6, 0)
	// 	check(err)
	// 	b3 := make([]byte, 2)
	// 	n3, err := io.ReadAtLeast(f, b3, 2)
	// 	check(err)
	// 	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 	_, err = f.Seek(0, 0)
	// 	check(err)

	// 	r4 := bufio.NewReader(f)
	// 	b4, err := r4.Peek(5)
	// 	check(err)
	// 	fmt.Printf("5 bytes: %s\n", string(b4))

	// 	f.Close()

	// }

	// {
	// 	/*   写方法处理  */
	// 	f, err := os.Create("dat2")
	// 	check(err)

	// 	defer f.Close()

	// 	d2 := []byte{115, 111, 109, 101, 10}
	// 	n2, err := f.Write(d2)
	// 	check(err)
	// 	fmt.Printf("wrote %d bytes\n", n2)

	// 	n3, err := f.WriteString("writes\n")
	// 	fmt.Printf("wrote %d bytes\n", n3)

	// 	f.Sync()

	// 	w := bufio.NewWriter(f)
	// 	n4, err := w.WriteString("buffered\n")
	// 	fmt.Printf("wrote %d bytes\n", n4)

	// 	w.Flush()

	// }

	{
		for i := 0; i < 5; i++ {
			file := "test.log"

			start := time.Now()

			read0(file)
			t0 := time.Now()
			fmt.Printf("Cost time %v\n", t0.Sub(start))

			read1(file)
			t1 := time.Now()
			fmt.Printf("Cost time %v\n", t1.Sub(t0))

			read2(file)
			t2 := time.Now()
			fmt.Printf("Cost time %v\n", t2.Sub(t1))

			read3(file)
			t3 := time.Now()
			fmt.Printf("Cost time %v\n", t3.Sub(t2))

			fmt.Println()
		}

	}
}

func read0(path string) string {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}
	return string(f)
}

func read1(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	return string(chunks)
}

func read2(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)

	chunks := make([]byte, 1024, 1024)

	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	return string(chunks)
}

func read3(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}
