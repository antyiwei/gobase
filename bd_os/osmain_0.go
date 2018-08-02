package main

// func main() {
// 	// {
// 	// 	/* 创建文件夹 */
// 	// 	err := os.Mkdir("test_go", 0777)
// 	// 	if err != nil {
// 	// 		panic(err.Error())
// 	// 	}
// 	// }
// 	// {
// 	// 	/* 删除文件夹 */
// 	// 	err := os.Remove("test_go")
// 	// 	if err != nil {
// 	// 		fmt.Println(err)
// 	// 	}
// 	// 	os.RemoveAll("test_go")
// 	// }
// 	// {
// 	// 	os.MkdirAll("test_go/go1/go2", 0777)
// 	// }

// 	// {
// 	// 	/* 创建文件 */
// 	// 	file, err := os.Create("test_go.txt")
// 	// 	if err != nil {
// 	// 		fmt.Println(err)
// 	// 		return
// 	// 	}
// 	// 	file.Write([]byte("hello world ss"))

// 	// 	dat, err := ioutil.ReadFile("test_go.txt")
// 	// 	if err != nil {
// 	// 		log.Fatal(err)
// 	// 	}
// 	// 	fmt.Print(string(dat))
// 	// }
// 	// {
// 	// 	f, err := os.OpenFile("file.go", os.O_RDWR|os.O_CREATE, 0755)
// 	// 	if err != nil {
// 	// 		log.Fatal(err)
// 	// 	}
// 	// 	if err := f.Close(); err != nil {
// 	// 		log.Fatal(err)
// 	// 	}
// 	// }
// 	{
// 		/* 创建文件 */
// 		userFile := "test_go.txt"
// 		f1, err := os.Open(userFile)
// 		if err != nil {
// 			fmt.Println(userFile, err)
// 			return
// 		}
// 		defer f1.Close()
// 		buf := make([]byte, 1024)
// 		for {
// 			n, _ := f1.Read(buf)
// 			if 0 == n {
// 				break
// 			}
// 			os.Stdout.Write(buf[:n])
// 		}
// 	}
// }

// func check(e error) {
// 	if e != nil {
// 	panic(e)
// 	}
// }
