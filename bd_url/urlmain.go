package main

import (
	"fmt"
	"log"
	"net/url"
)

/*
URL基本格式如下：
scheme://[userinfo@]host/path[?query][#fragment]
*/
func main() {
	// modeParse()

	// mode ParseRequestURI 特殊
	// modeParseRequestURI()

	// mode IsAbs
	// modeIsAbs()

	// modeQuery()

	// modePort()

	// modeRequestURI()

	// modeResolveReference()

	// modeValues()

	modeLast()
}

func modeParse() {
	u, err := url.Parse("http://bing.com:80802/search?q=dotnet&c=233")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)        //http://bing.com/search?q=dotnet&c=233
	fmt.Println(u.Scheme) //http
	fmt.Println(u.Opaque) //
	println(u.Host)       // bing.com
	println(u.Path)       // search
	println(u.RawPath)    //
	println(u.ForceQuery) //false
	println(u.RawQuery)   //q=dotnet&c=233
	println(u.Fragment)   //
	println(u.Port())     // 端口
}

func modeParseRequestURI() {
	values, err := url.ParseRequestURI("https://www.baidu.com/s?wd=%E6%90%9C%E7%B4%A2&rsv_spt=1&issp=1&f=8&rsv_bp=0&rsv_idx=2&ie=utf-8&tn=baiduhome_pg&rsv_enter=1&rsv_sug3=7&rsv_sug1=6#jj=ljlj")
	fmt.Println(values)
	if err != nil {
		fmt.Println(err)
	}
	urlParam := values.RawQuery
	fmt.Println(urlParam)

	values, err = url.ParseRequestURI("https://www.baidu.com/s?wd=中文&account=18368870825&rsv_spt=1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(values.RequestURI())
}

func modeIsAbs() {
	url, _ := url.Parse("//example.com")
	if url.IsAbs() {
		fmt.Printf("%v is absolute.", url)
	} else {
		fmt.Printf("%v is relative. Are you sure?", url)
	}
}

func modeQuery() {
	u, _ := url.Parse("http://www.example.com/path?foo=32&f=2322&b=kjj&fj=")
	fmt.Println(u)
	u.RawQuery = u.Query().Encode()
	fmt.Println(u)
}

// modePort 返回端口号，如果host中没有端口号，接返回一个空的字符串
func modePort() {
	u, _ := url.Parse("http://www.example.com:8080P/path?foo")
	fmt.Println(u.Port())
}

// func modeParseRequestURI

func modeRequestURI() {

	values, err := url.ParseRequestURI("https://www.baidu.com/s?wd=%E6%90%9C%E7%B4%A2&rsv_spt=1&issp=1&f=8&rsv_bp=0&rsv_idx=2&ie=utf-8&tn=baiduhome_pg&rsv_enter=1&rsv_sug3=7&rsv_sug1=6#jj=ljlj")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(values.RequestURI())
}

// modeResolveReference 本方法根据一个绝对URI将一个URI补全为一个绝对URI，参见RFC 3986 节 5.2。参数ref可以是绝对URI或者相对URI。ResolveReference总是返回一个新的URL实例，即使该实例和u或者ref完全一样。如果ref是绝对URI，本方法会忽略参照URI并返回ref的一个拷贝。
func modeResolveReference() {
	u, err := url.Parse("../../..//search?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}
	base, err := url.Parse("http://example.com/directory/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(base.ResolveReference(u))
}

// modeValues Values将建映射到值的列表。它一般用于查询的参数和表单的属性。不同于http.Header这个字典类型，Values的键是大小写敏感的。
func modeValues() {
	v := url.Values{}
	v.Set("name", "antyiwei")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")

	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])

	v.Del("name")
	fmt.Println(v["friend"])

	v.Del("friend")
	fmt.Println(v["friend"])

}

func modeLast() {
	c := url.Values{"method": {"get", "post"}, "id": {"1"}}
	fmt.Println(c.Encode())
	fmt.Println(c.Get("method"))
}
