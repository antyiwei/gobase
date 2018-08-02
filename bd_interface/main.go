package main

import (
	"fmt"
	"reflect"
)

type Skills interface {
	Running()
	GetName() string
}
type Test interface {
	Skills
	sleeping()
}
type Student struct {
	Name string
	Age  int
}

type Teacher struct {
	Name   string
	Salary int
}

// func (p Student) GetName() string {
// 	fmt.Println(p.Name)
// 	return p.Name
// }

// func (p Student) Running() {
// 	fmt.Printf("%s running...\n", p.Name)
// }

func (p Teacher) GetName() string {
	fmt.Println(p.Name)
	return p.Name
}

func (p Teacher) Running() {
	fmt.Printf("%s running...\n", p.Name)
}

func (this *Student) SetName(name string) {
	this.Name = name
	fmt.Printf("set name %s\n", this.Name)
}

func (this *Student) SetAge(age int) {
	this.Age = age
	fmt.Printf("set age %d\n", age)
}

func (this *Student) String() string {
	fmt.Printf("this is %s\n", this.Name)
	return this.Name
}

func main() {
	/* 	{
	   		var h1 Skills
	   		var p Student
	   		var t1 Teacher
	   		t1.Name = "wang xiao shan"

	   		p.Name = "antyiwei"
	   		p.Age = 28
	   		h1 = p
	   		h1.Running()
	   		h1 = t1
	   		t1.Running()
	   	}
	*/
	{

		// reflect
		stu1 := Student{Name: "antyiwei", Age: 90}
		inf := new(Skills)
		stu_type := reflect.TypeOf(stu1)
		inf_type := reflect.TypeOf(inf).Elem()

		fmt.Println(stu_type.String())
		fmt.Println(stu_type.Name())
		fmt.Println(stu_type.PkgPath())
		fmt.Println(stu_type.Kind())
		fmt.Println(stu_type.Size())
		fmt.Println(inf_type.NumMethod())

		fmt.Println(inf_type.Method(1), inf_type.Method(1).Name)
		fmt.Println(inf_type.MethodByName("running"))
	}

	{

		stu1 := Student{Name: "antyiwei", Age: 28}
		stuType := reflect.TypeOf(stu1)
		fmt.Println(stuType.NumField())
		fmt.Println(stuType.Field(0))
		fmt.Println(stuType.FieldByName("Age"))
	}

	{

		/* reflect.Value */
		str := "antyiwei"
		val := reflect.ValueOf(str).Kind()
		fmt.Println(val) // string
	}

	{
		str := "antyiwei"
		age := 29
		fmt.Println(reflect.ValueOf(str).String) //获取str的值，结果wd
		fmt.Println(reflect.ValueOf(age).Int)    //获取age的值，结果age
		str2 := reflect.ValueOf(&str)            //获取Value类型
		str2.Elem().SetString("jack")            //设置值
		fmt.Println(str2.Elem(), age)            //jack 11
	}

	{
		stu1 := Student{Name: "antyiwei", Age: 28}
		stuVal := reflect.ValueOf(stu1)
		fmt.Println(stuVal.NumField())
		fmt.Println(stuVal.Field(0), stuVal.Field(1))
		fmt.Println(stuVal.FieldByName("Age"))
		stuVal2 := reflect.ValueOf(&stu1).Elem()
		stuVal2.FieldByName("Age").SetInt(33)
		fmt.Println(stu1.Age)
	}
	{
		stu1 := &Student{Name: "wd", Age: 22}
		val := reflect.ValueOf(stu1)         //获取Value类型，也可以使用reflect.ValueOf(&stu1).Elem()
		val.MethodByName("String").Call(nil) //调用String方法

		params := make([]reflect.Value, 1)
		params[0] = reflect.ValueOf(18)
		val.MethodByName("SetAge").Call(params) //通过名称调用方法

		params[0] = reflect.ValueOf("jack")
		val.Method(1).Call(params) //通过方法索引调用

		fmt.Println(stu1.Name, stu1.Age)

	}
}
