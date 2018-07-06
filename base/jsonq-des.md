 ## [原文来自https://hashnode.com/post/how-to-access-deeply-nested-json-data-using-go-lang-cjidni8w3000cxms18hn7f4sy][1]
   
大多数情况下，开发人员需要使用来自其他服务的JSON数据并对其进行查询。查询JSON文档非常耗时。在过去的几天里，我正在为Golang编写一个包，以便轻松查询JSON数据。这个想法和灵感来自 Nahid Bin Azhar的PHP-JSONQ。

我们来看一个示例JSON数据：

```json
{
   "name":"computers",
   "description":"List of computer products",
   "vendor":{
      "name":"Star Trek",
      "email":"info@example.com",
      "website":"www.example.com",
      "items":[
         {"id":1, "name":"MacBook Pro 13 inch retina","price":1350},
         {"id":2, "name":"MacBook Pro 15 inch retina", "price":1700},
         {"id":3, "name":"Sony VAIO", "price":1200},
         {"id":4, "name":"Fujitsu", "price":850},
         {"id":5, "name":"HP core i5", "price":850, "key": 2300},
         {"id":6, "name":"HP core i7", "price":950},
         {"id":null, "name":"HP core i3 SSD", "price":850}
      ],
      "prices":[
         2400,
         2100,
         1200,
         400.87,
         89.90,
         150.10
     ]
   }
}
```

让我们找到一个深度嵌套的属性并正确处理错误，在这种情况下，我们将尝试<font color=red>name</font>从<font color=red>items</font>数组的第二个元素进行访问，注意：<font color=red>items</font>是<font color=red>vendor</font>对象的属性。
请参阅以下示例：

```go
package main

import (
    "fmt"
    "log"

    "github.com/thedevsaddam/gojsonq"
)

func main() {
    jq := gojsonq.New().File("./sample-data.json")
    res := jq.Find("vendor.items.[1].name")

    if jq.Error() != nil {
        log.Fatal(jq.Errors())
    }

    fmt.Println(res)
}
```

Yahooooo！很简单吧？它看起来像使用ORMJSON数据。让我们看一些更多示例来查询示例数据。

## 例1

查询： <font color=red>select * from vendor.items where price > 1200 or id null</font>

使用[ gojsonq ][2]我们可以执行以下查询：

```go
package main

import (
    "fmt"

    "github.com/thedevsaddam/gojsonq"
)

func main() {
    jq := gojsonq.New().File("./sample-data.json")
    res := jq.From("vendor.items").Where("price", ">", 1200).OrWhere("id", "=", nil).Get()
    fmt.Println(res)
    // output: [map[price:1350 id:1 name:MacBook Pro 13 inch retina] map[id:2 name:MacBook Pro 15 inch retina price:1700] map[id:<nil> name:HP core i3 SSD price:850]]
}
```

## 例2
查询：<font color=red> select name, price from vendor.items where price > 1200 or id null</font>

使用[ gojsonq ][2]我们可以执行以下查询：

```go
package main

import (
    "fmt"

    "github.com/thedevsaddam/gojsonq"
)

func main() {
    jq := gojsonq.New().File("./sample-data.json")
    res := jq.From("vendor.items").Where("price", ">", 1200).OrWhere("id", "=", nil).Only("name", "price")
    fmt.Println(res)
    // output: [map[name:MacBook Pro 13 inch retina price:1350] map[name:MacBook Pro 15 inch retina price:1700] map[name:HP core i3 SSD price:850]]
}
```

## 例3
查询：<font color=red> select sum(price) from vendor.items where price > 1200 or id null</font>

使用[ gojsonq ][2]我们可以执行以下查询：

```go
package main

import (
    "fmt"

    "github.com/thedevsaddam/gojsonq"
)

func main() {
    jq := gojsonq.New().File("./sample-data.json")
    res := jq.From("vendor.items").Where("price", ">", 1200).OrWhere("id", "=", nil).Sum("price")
    fmt.Println(res)
    // output: 3900
}
```

## 例4
查询： <font color=red>select price from vendor.items where price > 1200</font>

使用[ gojsonq ][2]我们可以执行以下查询：

```go
package main

import (
    "fmt"

    "github.com/thedevsaddam/gojsonq"
)

func main() {
    jq := gojsonq.New().File("./sample-data.json")
    res := jq.From("vendor.items").Where("price", ">", 1200).Pluck("price")
    fmt.Println(res)
    // output: [1350 1700]
}
```

## 例5
查询：<font color=red> select * from vendor.items order by price</font>

使用[ gojsonq ][2]我们可以执行以下查询：

```go
package main

import (
    "fmt"

    "github.com/thedevsaddam/gojsonq"
)

func main() {
    jq := gojsonq.New().File("./sample-data.json")
    res := jq.From("vendor.items").SortBy("price").Get()
    fmt.Println(res)
    // output: [map[id:<nil> name:HP core i3 SSD price:850] map[id:4 name:Fujitsu price:850] map[id:5 name:HP core i5 price:850 key:2300] map[id:6 name:HP core i7 price:950] map[id:3 name:Sony VAIO price:1200] map[id:1 name:MacBook Pro 13 inch retina price:1350] map[id:2 name:MacBook Pro 15 inch retina price:1700]]
}
```

## 例6
使用[ gojsonq ][2]您可以正确处理错误，请参阅下面的代码段：

```go
package main

import (
    "log"

    "github.com/thedevsaddam/gojsonq"
)

func main() {
    jq := gojsonq.New().File("./invalid-file.xjsn")
    err := jq.Error()
    if err != nil {
        log.Fatal(err)
        // 2018/06/25 00:48:58 gojsonq: open ./invalid-file.xjsn: no such file or directory
        // exit status 1
    }
}
```

## 例7
假设我们有一个像这样的JSON文档

```json
{
  "users":[
    {
      "id":1,
      "name":{
        "first":"John",
        "last":"Ramboo"
      }
    },
    {
      "id":2,
      "name":{
        "first":"Ethan",
        "last":"Hunt"
      }
    },
    {
      "id":3,
      "name":{
        "first":"John",
        "last":"Doe"
      }
    }
  ]
}
```

我们想要运行这样的查询：

查询：<font color=red> select * from users where name.first=John</font>

使用该软件包可以轻松进行查询，请参阅下面的代码段：

```go
package main

import (
    "fmt"

    "github.com/thedevsaddam/gojsonq"
)

func main() {
    jq := gojsonq.New().File("./data.json")
    res := jq.From("users").WhereEqual("name.first", "John").Get()
    fmt.Println(res) //output: [map[id:1 name:map[first:John last:Ramboo]] map[id:3 name:map[first:John last:Doe]]]
}
```

您可以使用DOT（。）访问嵌套级别属性，例如Where/GroupBy/SortBy etc

注意：还有一些其他有用的方法可以让生活更轻松！如果您喜欢该软件包，请不要忘记与您的社区分享并为存储库加注星标




  [1]: https://hashnode.com/post/how-to-access-deeply-nested-json-data-using-go-lang-cjidni8w3000cxms18hn7f4sy
  [2]:https://github.com/thedevsaddam/gojsonq