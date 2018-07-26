
# 转 Golang 单元测试和性能测试

开发程序其中很重要的一点是测试，我们如何保证代码的质量，如何保证每个函数是可运行，运行结果是正确的，又如何保证写出来的代码性能是好的，我们知道单元测试的重点在于发现程序设计或实现的逻辑错误，使问题及早暴露，便于问题的定位解决，而性能测试的重点在于发现程序设计上的一些问题，让线上的程序能够在高并发的情况下还能保持稳定。本小节将带着这一连串的问题来讲解Go语言中如何来实现单元测试和性能测试。

Go语言中自带有一个轻量级的测试框架testing和自带的go test命令来实现单元测试和性能测试，testing框架和其他语言中的测试框架类似，你可以基于这个框架写针对相应函数的测试用例，也可以基于该框架写相应的压力测试用例，那么接下来让我们一一来看一下怎么写。

# 如何编写测试用例

由于go test命令只能在一个相应的目录下执行所有文件，所以我们接下来新建一个项目目录gotest,这样我们所有的代码和测试代码都在这个目录下。

接下来我们在该目录下面创建两个文件：gotest.go和gotest_test.go

1.gotest.go:这个文件里面我们是创建了一个包，里面有一个函数实现了除法运算:

```go
    package gotest

    import (
        "errors"
    )

    func Division(a, b float64) (float64, error) {
        if b == 0 {
            return 0, errors.New("除数不能为0")
        }

        return a / b, nil
    }
```

2.gotest_test.go:这是我们的单元测试文件，但是记住下面的这些原则：

1. 文件名必须是_test.go结尾的，这样在执行go test的时候才会执行到相应的代码
2. 你必须import testing这个包
3. 所有的测试用例函数必须是Test开头
4. 测试用例会按照源代码中写的顺序依次执行
5. 测试函数TestXxx()的参数是testing.T，我们可以使用该类型来记录错误或者是测试状态
6. 测试格式：func TestXxx (t *testing.T),Xxx部分可以为任意的字母数字的组合，但是首字母不能是小写字母[a-z]，例如Testintdiv是错误的函数名。
7. 函数中通过调用testing.T的Error, Errorf, FailNow, Fatal, FatalIf方法，说明测试不通过，调用Log方法用来记录测试的信息。

下面是我们的测试用例的代码：

```go
    package gotest

    import (
        "testing"
    )

    func Test_Division_1(t *testing.T) {
        if i, e := Division(6, 2); i != 3 || e != nil { //try a unit test on function
            t.Error("除法函数测试没通过") // 如果不是如预期的那么就报错
        } else {
            t.Log("第一个测试通过了") //记录一些你期望记录的信息
        }
    }

    func Test_Division_2(t *testing.T) {
        t.Error("就是不通过")
    }
```

我们在项目目录下面执行go test,就会显示如下信息：

```go
--- FAIL: Test_Division_2 (0.00 seconds)
    gotest_test.go:16: 就是不通过
FAIL
exit status 1
FAIL    gotest  0.013s
```

从这个结果显示测试没有通过，因为在第二个测试函数中我们写死了测试不通过的代码t.Error，那么我们的第一个函数执行的情况怎么样呢？默认情况下执行go test是不会显示测试通过的信息的，我们需要带上参数go test -v，这样就会显示如下信息：

```go
=== RUN Test_Division_1
--- PASS: Test_Division_1 (0.00 seconds)
    gotest_test.go:11: 第一个测试通过了
=== RUN Test_Division_2
--- FAIL: Test_Division_2 (0.00 seconds)
    gotest_test.go:16: 就是不通过
FAIL
exit status 1
FAIL    gotest  0.012s
```

上面的输出详细的展示了这个测试的过程，我们看到测试函数1Test_Division_1测试通过，而测试函数2Test_Division_2测试失败了，最后得出结论测试不通过。接下来我们把测试函数2修改成如下代码：

```go
func Test_Division_2(t *testing.T) {
    if _, e := Division(6, 0); e == nil { //try a unit test on function
        t.Error("Division did not work as expected.") // 如果不是如预期的那么就报错
    } else {
        t.Log("one test passed.", e) //记录一些你期望记录的信息
    }
}  
```

然后我们执行go test -v，就显示如下信息，测试通过了：

```go
=== RUN Test_Division_1
--- PASS: Test_Division_1 (0.00 seconds)
    gotest_test.go:11: 第一个测试通过了
=== RUN Test_Division_2
--- PASS: Test_Division_2 (0.00 seconds)
    gotest_test.go:20: one test passed. 除数不能为0
PASS
ok      gotest  0.013s
```

# 如何编写压力测试

压力测试用来检测函数(方法）的性能，和编写单元功能测试的方法类似,此处不再赘述，但需要注意以下几点：

压力测试用例必须遵循如下格式，其中XXX可以是任意字母数字的组合，但是首字母不能是小写字母

```go
func BenchmarkXXX(b *testing.B) { ... }
```

go test不会默认执行压力测试的函数，如果要执行压力测试需要带上参数-test.bench，语法:-test.bench="test_name_regex",例如go test -test.bench=".*"表示测试全部的压力测试函数

在压力测试用例中,请记得在循环体内使用testing.B.N,以使测试可以正常的运行
文件名也必须以_test.go结尾

下面我们新建一个压力测试文件webbench_test.go，代码如下所示：

```go
package gotest

import (
    "testing"
)

func Benchmark_Division(b *testing.B) {
    for i := 0; i < b.N; i++ { //use b.N for looping 
        Division(4, 5)
    }
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
    b.StopTimer() //调用该函数停止压力测试的时间计数

    //做一些初始化的工作,例如读取文件数据,数据库连接之类的,
    //这样这些时间不影响我们测试函数本身的性能

    b.StartTimer() //重新开始时间
    for i := 0; i < b.N; i++ {
        Division(4, 5)
    }
}
```

我们执行命令go test -test.bench=".*"，可以看到如下结果：

```go
PASS
Benchmark_Division  500000000            7.76 ns/op
Benchmark_TimeConsumingFunction 500000000            7.80 ns/op
ok      gotest  9.364s  
```

上面的结果显示我们没有执行任何TestXXX的单元测试函数，显示的结果只执行了压力测试函数，第一条显示了Benchmark_Division执行了500000000次，每次的执行平均时间是7.76纳秒，第二条显示了Benchmark_TimeConsumingFunction执行了500000000，每次的平均执行时间是7.80纳秒。最后一条显示总共的执行时间。
我们执行命令go test -test.bench=".*" -count=5，可以看到如下结果： （使用-count可以指定执行多少次）

```go
PASS
Benchmark_Division-2                 300000000             4.60 ns/op
Benchmark_Division-2                 300000000             4.57 ns/op
Benchmark_Division-2                 300000000             4.63 ns/op
Benchmark_Division-2                 300000000             4.60 ns/op
Benchmark_Division-2                 300000000             4.63 ns/op
Benchmark_TimeConsumingFunction-2    300000000             4.64 ns/op
Benchmark_TimeConsumingFunction-2    300000000             4.61 ns/op
Benchmark_TimeConsumingFunction-2    300000000             4.60 ns/op
Benchmark_TimeConsumingFunction-2    300000000             4.59 ns/op
Benchmark_TimeConsumingFunction-2    300000000             4.60 ns/op
ok      _/home/diego/GoWork/src/app/testing    18.546s
```

```go
go test -run=文件名字 -bench=bench名字 -cpuprofile=生产的cprofile文件名称 文件夹
```

例子：
testBenchMark下有个popcnt文件夹，popcnt中有文件popcunt_test.go

```go
➜  testBenchMark ls
popcnt
```

popcunt_test.go的问价内容：

```go
package popcnt

import (
    "testing"
)

const m1 = 0x5555555555555555
const m2 = 0x3333333333333333
const m4 = 0x0f0f0f0f0f0f0f0f
const h01 = 0x0101010101010101

func popcnt(x uint64) uint64 {
    x -= (x >> 1) & m1
    x = (x & m2) + ((x >> 2) & m2)
    x = (x + (x >> 4)) & m4
    return (x * h01) >> 56
}

func BenchmarkPopcnt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        x := i
        x -= (x >> 1) & m1
        x = (x & m2) + ((x >> 2) & m2)
        x = (x + (x >> 4)) & m4
        _ = (x * h01) >> 56
    }
}
```

然后运行go test -bench=".*" -cpuprofile=cpu.profile ./popcnt

```go
➜  testBenchMark go test -bench=".*" -cpuprofile=cpu.profile ./popcnt
testing: warning: no tests to run
PASS
BenchmarkPopcnt-8    1000000000             2.01 ns/op
ok      app/testBenchMark/popcnt    2.219s
➜  testBenchMark ll
total 6704
drwxr-xr-x  5 diego  staff      170  5  6 13:57 .
drwxr-xr-x  3 diego  staff      102  5  6 11:12 ..
-rw-r--r--  1 diego  staff     5200  5  6 13:57 cpu.profile
drwxr-xr-x  4 diego  staff      136  5  6 11:47 popcnt
-rwxr-xr-x  1 diego  staff  3424176  5  6 13:57 popcnt.test
➜  testBenchMark
```

生产 cpu.profile问价和popcnt.test 文件

```go
➜  testBenchMark ll
total 6704
drwxr-xr-x  5 diego  staff      170  5  6 13:57 .
drwxr-xr-x  3 diego  staff      102  5  6 11:12 ..
-rw-r--r--  1 diego  staff     5200  5  6 13:57 cpu.profile
drwxr-xr-x  3 diego  staff      102  5  6 14:01 popcnt
-rwxr-xr-x  1 diego  staff  3424176  5  6 13:57 popcnt.test
➜  testBenchMark
```

```go
go tool pprof popcnt.test cpu.profile 进入交互模式
```

```go
➜  testBenchMark go tool pprof popcnt.test cpu.profile
Entering interactive mode (type "help" for commands)
(pprof) top
1880ms of 1880ms total (  100%)
      flat  flat%   sum%        cum   cum%
    1790ms 95.21% 95.21%     1790ms 95.21%  app/testBenchMark/popcnt.BenchmarkPopcnt
      90ms  4.79%   100%       90ms  4.79%  runtime.usleep
         0     0%   100%     1790ms 95.21%  runtime.goexit
         0     0%   100%       90ms  4.79%  runtime.mstart
         0     0%   100%       90ms  4.79%  runtime.mstart1
         0     0%   100%       90ms  4.79%  runtime.sysmon
         0     0%   100%     1790ms 95.21%  testing.(*B).launch
         0     0%   100%     1790ms 95.21%  testing.(*B).runN
(pprof)
```

go tool pprof --web popcnt.test cpu.profile 进入web模式

```go
$ go tool pprof --text mybin http://myserver:6060:/debug/pprof/profile
````

这有几个可用的输出类型，最有用的几个为： --text，--web 和 --list 。运行 go tool pprof 来得到最完整的列表。

格式形如：
```go
go test [-c] [-i] [build flags] [packages] [flags for test binary]

参数解读：
-c : 编译go test成为可执行的二进制文件，但是不运行测试。

-i : 安装测试包依赖的package，但是不运行测试。

关于build flags，调用go help build，这些是编译运行过程中需要使用到的参数，一般设置为空

关于packages，调用go help packages，这些是关于包的管理，一般设置为空

关于flags for test binary，调用go help testflag，这些是go test过程中经常使用到的参数

-test.v : 是否输出全部的单元测试用例（不管成功或者失败），默认没有加上，所以只输出失败的单元测试用例。

-test.run pattern: 只跑哪些单元测试用例

-test.bench patten: 只跑那些性能测试用例

-test.benchmem : 是否在性能测试的时候输出内存情况

-test.benchtime t : 性能测试运行的时间，默认是1s

-test.cpuprofile cpu.out : 是否输出cpu性能分析文件

-test.memprofile mem.out : 是否输出内存性能分析文件

-test.blockprofile block.out : 是否输出内部goroutine阻塞的性能分析文件

-test.memprofilerate n : 内存性能分析的时候有一个分配了多少的时候才打点记录的问题。这个参数就是设置打点的内存分配间隔，也就是profile中一个sample代表的内存大小。默认是设置为512 * 1024的。如果你将它设置为1，则每分配一个内存块就会在profile中有个打点，那么生成的profile的sample就会非常多。如果你设置为0，那就是不做打点了。

你可以通过设置memprofilerate=1和GOGC=off来关闭内存回收，并且对每个内存块的分配进行观察。

-test.blockprofilerate n: 基本同上，控制的是goroutine阻塞时候打点的纳秒数。默认不设置就相当于-test.blockprofilerate=1，每一纳秒都打点记录一下

-test.parallel n : 性能测试的程序并行cpu数，默认等于GOMAXPROCS。

-test.timeout t : 如果测试用例运行时间超过t，则抛出panic

-test.cpu 1,2,4 : 程序运行在哪些CPU上面，使用二进制的1所在位代表，和nginx的nginx_worker_cpu_affinity是一个道理

-test.short : 将那些运行时间较长的测试用例运行时间缩短
```