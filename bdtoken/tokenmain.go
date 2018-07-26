package main

import (
	"fmt"
	"go/token"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Hello antyiwei,learning token。。。")

	TestToken()
}

func TestToken() {
	{
		tokenstr := token.Lookup("==")
		fmt.Println(tokenstr)
		fmt.Println(tokenstr.IsKeyword()) // 判读goto是否时关键字
		fmt.Println(tokenstr.IsLiteral())
		fmt.Println(tokenstr.IsOperator())
		fmt.Println(tokenstr.Precedence())
		fmt.Println(tokenstr.String())
	}
	{

		pos := token.Position{
			Filename: "sb",
			Offset:   10,
			Line:     19,
			Column:   28,
		}

		fmt.Println(pos.IsValid())
		fmt.Println(pos.String())
	}
	{
		// 用go golang获取access_token
		v := url.Values{}
		v.Set("grant_type", "client_credentials") //规定，无需修改
		v.Add("client_id", "YM1OvG...K")          //根据自己申请填写
		v.Add("client_secret", "S926...Wq")       //根据自己申请填写
		res, err := http.PostForm("https://aip.baidubce.com/oauth/2.0/token", v)
		robots, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", robots)
	}
	{
		var substr = "abc"
		var str = "abcdabcadbabckihk"
		fmt.Println(strings.LastIndex(str, substr))
		fmt.Println(str[strings.LastIndex(str, substr)+len(substr):])
	}

}
