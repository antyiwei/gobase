package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) GetName() string {
	return p.Name
}
func (p Person) SetName(s string) {
	p.Name = s
}

type MyStruct struct {
	name string
}

func (this *MyStruct) GetName() string {
	return this.name
}

// func main() {
// 	// ReflectZero()
// 	// RelectOne()
// 	RelectTwo()
// 	// ReflectThree()
// }
func ReflectZero() {
	a := &Person{"XiaoMing", 10}
	t := reflect.TypeOf(*a) //必须取值，否则类型为空
	fmt.Println(t.Name())

	fmt.Println(reflect.ValueOf(a))
	v := reflect.ValueOf(a).Elem() //a需要是引用
	fmt.Println(v)
	k := v.Type()
	for i := 0; i < v.NumField(); i++ {
		key := k.Field(i)
		val := v.Field(i)
		fmt.Println(key.Name, val.Type(), val.Interface())
	}

	v.FieldByName("Name").Set(reflect.ValueOf("antyiwei"))
	fmt.Println(a.Name)
	name := v.MethodByName("GetName").Call([]reflect.Value{})
	fmt.Println(name)
}
func RelectOne() {

	s := "this is string"
	fmt.Println(reflect.TypeOf(s))
	fmt.Println("-------------------")

	fmt.Println(reflect.ValueOf(s))
	var x float64 = 3.4
	fmt.Println(reflect.ValueOf(x))
	fmt.Println("-------------------")

	a1 := new(MyStruct)
	a1.name = "yejianfeng"
	typ := reflect.TypeOf(a1)

	fmt.Println(typ.NumMethod())
	fmt.Println("-------------------")

	b := reflect.ValueOf(a1).MethodByName("GetName").Call([]reflect.Value{})
	fmt.Println(b[0])
}
func RelectTwo() {

	fmt.Println("--------------")
	var a MyStruct
	b := new(MyStruct)
	fmt.Println(reflect.ValueOf(a))
	fmt.Println(reflect.ValueOf(b))

	fmt.Println("--------------")
	a.name = "yejianfeng"
	b.name = "yejianfeng"
	val := reflect.ValueOf(a).FieldByName("name")

	//painc: val := reflect.ValueOf(b).FieldByName("name")
	fmt.Println(val)

	fmt.Println("--------------")
	fmt.Println(reflect.ValueOf(a).FieldByName("name").CanSet())
	fmt.Println(reflect.ValueOf(&(a.name)).Elem().CanSet())

	fmt.Println("--------------")
	var c string = "yejianfeng"
	p := reflect.ValueOf(&c)
	fmt.Println(p.CanSet())        //false
	fmt.Println(p.Elem().CanSet()) //true
	p.Elem().SetString("9999999")
	fmt.Println(c)

}

type IStruct interface {
	GetName() string
}

func ReflectThree() {
	// TypeOf
	s := "this is string"
	fmt.Println(reflect.TypeOf(s)) // output: "string"

	// object TypeOf
	a := new(MyStruct)
	a.name = "yejianfeng"
	typ := reflect.TypeOf(a)
	fmt.Println(typ)        // output: "*main.MyStruct"
	fmt.Println(typ.Elem()) // output: "main.MyStruct"

	// reflect.Type Base struct
	fmt.Println(typ.NumMethod())                   // 1
	fmt.Println(typ.Method(0))                     // {GetName  func(*main.MyStruct) string <func(*main.MyStruct) string Value> 0}
	fmt.Println(typ.Name())                        // ""
	fmt.Println(typ.PkgPath())                     // ""
	fmt.Println(typ.Size())                        // 8
	fmt.Println(typ.String())                      // *main.MyStruct
	fmt.Println(typ.Elem().String())               // main.MyStruct
	fmt.Println(typ.Elem().FieldByIndex([]int{0})) // {name main string  0 [0] false}
	fmt.Println(typ.Elem().FieldByName("name"))    // {name main string  0 [0] false} true

	fmt.Println(typ.Kind() == reflect.Ptr)                              // true
	fmt.Println(typ.Elem().Kind() == reflect.Struct)                    // true
	fmt.Println(typ.Implements(reflect.TypeOf((*IStruct)(nil)).Elem())) // true

	fmt.Println(reflect.TypeOf(12.12).Bits()) // 64, 因为是float64

	cha := make(chan int)
	fmt.Println(reflect.TypeOf(cha).ChanDir()) // chan

	var fun func(x int, y ...float64) string
	var fun2 func(x int, y float64) string
	fmt.Println(reflect.TypeOf(fun).IsVariadic())  // true
	fmt.Println(reflect.TypeOf(fun2).IsVariadic()) // false
	fmt.Println(reflect.TypeOf(fun).In(0))         // int
	fmt.Println(reflect.TypeOf(fun).In(1))         // []float64
	fmt.Println(reflect.TypeOf(fun).NumIn())       // 2
	fmt.Println(reflect.TypeOf(fun).NumOut())      // 1
	fmt.Println(reflect.TypeOf(fun).Out(0))        // string

	mp := make(map[string]int)
	mp["test1"] = 1
	fmt.Println(reflect.TypeOf(mp).Key()) //string

	arr := [1]string{"test"}
	fmt.Println(reflect.TypeOf(arr).Len()) // 1

	fmt.Println(typ.Elem().NumField()) // 1

	// MethodByName, Call
	b := reflect.ValueOf(a).MethodByName("GetName").Call([]reflect.Value{})
	fmt.Println(b[0]) // output: "yejianfeng"

}
