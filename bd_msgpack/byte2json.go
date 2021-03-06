package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/vmihailenco/msgpack"
)

type AutoGenerated struct {
	V       string   `json:"v"`
	Mid     int      `json:"mid"`
	Time    int      `json:"time"`
	IP      string   `json:"ip"`
	Mac     string   `json:"mac"`
	Devices []string `json:"devices"`
}

func main() {

	b, err := ioutil.ReadFile("16hex2.txt")
	if err != nil {
		fmt.Print(err)
	}
	//fmt.Println(b)

	s := string(b)
	//fmt.Println(s)

	{
		test, _ := hex.DecodeString(s)

		var out map[string]interface{}
		err = msgpack.Unmarshal(test, &out)
		if err != nil {
			panic(err)
		}
		fmt.Println("IP =", out["ip"])
		fmt.Println("V =", out["v"])
		fmt.Println("Mid =", out["mid"])

		b, err := json.Marshal(out)
		if err != nil {
			fmt.Println("json.Marshal failed:", err)
			return
		}

		//fmt.Println("b:", string(b))

		var item AutoGenerated
		err = json.Unmarshal(b, &item)
		if err != nil {
			panic(err)
		}
		fmt.Println(item.V, item.Mid, item.Time, item.IP, item.Mac)
		for i := 0; i < len(item.Devices); i++ {
			fmt.Println(i+1, "item.Devices[", i, "] = ", item.Devices[i])
		}
	}

	//{
	//
	//	type Item struct {
	//		Foo string
	//	}
	//
	//	b, err := msgpack.Marshal(&Item{Foo: "bar"})
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	var item Item
	//	err = msgpack.Unmarshal(b, &item)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(item.Foo)
	//	fmt.Printf("%#v", &item)
	//}
	//
	//{
	//
	//	type Item struct {
	//		_msgpack struct{} `msgpack:",asArray"`
	//		Foo      string
	//		Bar      string
	//	}
	//
	//	var buf bytes.Buffer
	//	enc := msgpack.NewEncoder(&buf)
	//	err := enc.Encode(&Item{Foo: "foo", Bar: "bar"})
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	dec := msgpack.NewDecoder(&buf)
	//	v, err := dec.DecodeInterface()
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(v)
	//}
	//
	//{
	//
	//	in := map[string]interface{}{"foo": 1, "hello": "world"}
	//	b, err := msgpack.Marshal(in)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	var out map[string]interface{}
	//	err = msgpack.Unmarshal(b, &out)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	fmt.Println("foo =", out["foo"])
	//	fmt.Println("hello =", out["hello"])
	//}
}

func xtop(x string) string {
	base, _ := strconv.ParseInt(x, 16, 10)
	return strconv.FormatInt(base, 2)
}

func btox(b string) string {
	base, _ := strconv.ParseInt(b, 2, 10)
	return strconv.FormatInt(base, 2)
}
