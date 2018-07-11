package jsonant

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

// type Person struct {
//     name string
//     age int
// }

// func TestStruct2Json(t *testing.T) {
//     jsonStr := `
//     {
//         "name":"liangyongxing",
//         "age":12
//     }
//     `
//     var person Person
//     json.Unmarshal([]byte(jsonStr), &person)
//     t.Log(person)
// }

/* ================================================= */

// type Person struct {
// 	Name string
// 	Age  int
// }

// func TestStruct2Json(t *testing.T) {
// 	// jsonStr := `
// 	// {
// 	//     "name":"liangyongxing",
// 	//     "age":12
// 	// }
// 	// `
// 	jsonStr := `
//     {
//         "NaMe":"liangyongxing",
//         "agE":12
//     }
//     `
// 	var person Person
// 	json.Unmarshal([]byte(jsonStr), &person)
// 	t.Log(person)
// }

/* ================================================= */

// //这里对应的 N 和 A 不能为小写，首字母必须为大写，这样才可对外提供访问，具体 json 匹配是通过后面的 tag 标签进行匹配的，与 N 和 A 没有关系
// //tag 标签中 json 后面跟着的是字段名称，都是字符串类型，要求必须加上双引号，否则 golang 是无法识别它的类型
// type Person struct {
// 	N string `json:"name"`
// 	A int    `json:"age"`
// }

// func TestStruct2Json(t *testing.T) {
// 	jsonStr := `
//     {
//         "name":"liangyongxing",
//         "age":12
//     }
//     `
// 	var person Person
// 	json.Unmarshal([]byte(jsonStr), &person)
// 	t.Log(person)
// }

/* ================================================= */
// type Person struct {
// 	Name string
// 	Age  int
// }

// func TestStruct2Json(t *testing.T) {
// 	p := Person{
// 		Name: "liangyongxing",
// 		Age:  29,
// 	}

// 	t.Logf("Person 结构体打印结果:%v", p)

// 	//Person 结构体转换为对应的 Json
// 	jsonBytes, err := json.Marshal(p)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("转换为 json 串打印结果:%s", string(jsonBytes))
// }

// /* ================================================= */
// func TestJson2Map(t *testing.T) {
// 	jsonStr := `
//     {
//         "name":"liangyongxing",
//         "age":12
//     }
//     `
// 	var mapResult map[string]interface{}
// 	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
// 	if err := json.Unmarshal([]byte(jsonStr), &mapResult); err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(mapResult)
// }
/* ================================================= */

// func TestMap2Json(t *testing.T) {
// 	mapInstance := make(map[string]interface{})
// 	mapInstance["Name"] = "liang637210"
// 	mapInstance["Age"] = 28
// 	mapInstance["Address"] = "北京昌平区"

// 	jsonStr, err := json.Marshal(mapInstance)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Logf("Map2Json 得到 json 字符串内容:%s", jsonStr)
// }

/* ================================================= */

func TestMap2Struct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	mapInstance := make(map[string]interface{})
	mapInstance["Name"] = "liang637210"
	mapInstance["age"] = 28

	jsonBytes, _ := json.Marshal(mapInstance)

	var person Person
	//将 map 转换为指定的结构体
	// if err := mapstructure.Decode(mapInstance, &person); err != nil {
	// 	t.Fatal(err)
	// }

	if err := json.Unmarshal(jsonBytes, &person); err != nil {
		t.Fatal(err)
	}
	t.Logf("map2struct后得到的 struct 内容为:%v", person)
}

/* ================================================= */
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func TestStruct2Map(t *testing.T) {
	user := User{5, "zhangsan", "password"}
	data := Struct2Map(user)
	t.Logf("struct2map得到的map内容为:%v", data)
}

/* ================================================= */
func TestOsHostname(t *testing.T) {
	t.Log(os.Hostname())
}

/* ================================================= */

/* ================================================= */

/* ================================================= */

/* ================================================= */

/* ================================================= */
