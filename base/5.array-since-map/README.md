# 数组Array

- 定义数组的格式：var <varName> [n]<type>，n>=0
- 数组长度也是类型的一部分，因此具有不同长度的数组为不同类型
- 注意区分指向数组的指针和指针数组
- 数组在Go中为值类型
- 数组之间可以使用==或!=进行比较，但不可以使用<或>
- 可以使用new来创建数组，此方法返回一个指向数组的指针
- Go支持多维数组


- Go语言版冒泡排序

## 切片Slice

- 其本身并不是数组，它指向底层的数组
- 作为变长数组的替代方案，可以关联底层数组的局部或全部
- 为引用类型
- 可以直接创建或从底层数组获取生成
- 使用len()获取元素个数，cap()获取容量
- 一般使用make()创建
- 如果多个slice指向相同底层数组，其中一个的值改变会影响全部

- make([]T, len, cap)
- 其中cap可以省略，则和len的值相同
- len表示存数的元素个数，cap表示容量


 Slice与底层数组的对应关系
    ![Alt text](/img/since.png "Slice与底层数组的对应关系")

- Reslice
    - Reslice时索引以被slice的切片为准
    - 索引不可以超过被slice的切片的容量cap()值
    - 索引越界不会导致底层数组的重新分配而是引发错误
- Append
    - 可以在slice尾部追加元素
    - 可以将一个slice追加在另一个slice尾部
    - 如果最终长度未超过追加到slice的容量则返回原始slice
    - 如果超过追加到的slice的容量则将重新分配数组并拷贝原始数据
- Copy 


## Map

- 类似其它语言中的哈希表或者字典，以key-value形式存储数据
- Key必须是支持==或!=比较运算的类型，不可以是函数、map或slice
- Map查找比线性搜索快很多，但比使用索引访问数据的类型慢100倍
- Map使用make()创建，支持 := 这种简写方式

- make([keyType]valueType, cap)，cap表示容量，可省略
- 超出容量时会自动扩容，但尽量提供一个合理的初始值
- 使用len()获取元素个数

- 键值对不存在时自动添加，使用delete()删除某键值对
- 使用 for range 对map和slice进行迭代操作



## 课堂作业

- 根据在 for range 部分讲解的知识，尝试将类型为map[int]string的键和值进行交换，变成类型map[string]int
- 程序正确运行后应输出如下结果：

![Alt text](/img/map_result.png "Slice与底层数组的对应关系")

```go
package main

import "fmt"

func main() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "f", 7: "g", 8: "h", 9: "i", 0: "j"}
	fmt.Println(m1)
	m2 := make(map[string]int)
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m2)
}
```