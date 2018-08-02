

## 原文来自：[理解Go中的context包-Understanding the context package in golang][1]

go中的context包在与API和慢进程[客户端]交互时可以派上用场，特别是在提供Web请求的生产级系统中。 在哪里，您可能想要通知所有goroutines停止工作并返回。 这是一个基本教程，介绍如何在项目中使用它以及一些最佳实践和陷阱。

要了解上下文包，您应该熟悉两个概念。

在转到上下文之前，我将尝试简要介绍这些内容，如果您已经熟悉这些内容，则可以直接转到上下文部分。

### Goroutine
从官方go文档：“goroutine是一个轻量级的执行线程”。 Goroutines比线程更轻，因此管理它们的资源相对较少。

码:[https://play.golang.org/p/-TDMgnkJRY6](https://play.golang.org/p/-TDMgnkJRY6)
```go
package main
import "fmt"
  
//function to print hello
func printHello() {
    fmt.Println("Hello from printHello")
}
func main() {
    //inline goroutine. Define a function inline and then call it.
    go func(){fmt.Println("Hello inline")}()
    //call a function as goroutine
    go printHello()
    fmt.Println("Hello from main")
}
```

如果你运行上面的程序，你可能只看到它打印出来自main的Hello，因为它会激活几个goroutine并且主函数在它们完成之前退出。 为了确保主要等待goroutine完成，你需要一些方法让goroutines告诉它它们已经完成执行，这是通道可以帮助我们的地方。

### Channels
这些是goroutines之间的沟通渠道。 当您想要将结果或错误或任何其他类型的信息从一个goroutine传递到另一个goroutine时，将使用通道。 通道有类型，可以有int类型的通道接收整数或错误接收错误等。

假设存在类型为int的通道ch如果要向通道发送内容，则语法为ch < - 1如果要从通道接收某些内容，则为var：= < - ch。 这将从通道中获取并将值存储在var中。

以下程序说明了使用通道来确保goroutine完成并将值从它们返回到main。

注意：等待组（https://golang.org/pkg/sync/#WaitGroup）也可用于同步，但由于我们稍后在上下文部分讨论了频道，因此我将在此博客的代码示例中选择它们岗位

码：[https://play.golang.org/p/3zfQMox5mHn](https://play.golang.org/p/3zfQMox5mHn)

```go
package main
import "fmt"
  
//prints to stdout and puts an int on channel
func printHello(ch chan int) {
    fmt.Println("Hello from printHello")
    //send a value on channel
    ch <- 2
}
func main() {
    //make a channel. You need to use the make function to create channels.
    //channels can also be buffered where you can specify size. eg: ch := make(chan int, 2)
    //that is out of the scope of this post.
    ch := make(chan int)
    //inline goroutine. Define a function and then call it.
    //write on a channel when done
    go func(){
      fmt.Println("Hello inline")
      //send a value on channel
      ch <- 1
    }()
    //call a function as goroutine
    go printHello(ch)
    fmt.Println("Hello from main")
//get first value from channel.
    //and assign to a variable to use this value later
    //here that is to print it
    i := <- ch
    fmt.Println("Recieved ",i)
    //get the second value from channel
    //do not assign it to a variable because we dont want to use that
    <- ch
}
```

考虑上下文包的方法是允许您将“上下文”传递给您的程序。 上下文，如超时或截止时间或指示停止工作和返回的通道。 例如，如果您正在执行Web请求或运行系统命令，那么对生产级系统进行超时通常是个好主意。 因为，如果您依赖的API运行缓慢，则您不希望在系统上备份请求，因为它可能最终会增加负载并降低您所服务的所有请求的性能。 导致级联效应。 这是超时或截止日期上下文可以派上用场的地方。

### Creating context
上下文包允许以下列方式创建和派生上下文：

#### context.Background() ctx Context

此函数返回空上下文。这应仅用于高级别（在主要或顶级请求处理程序中）。这可以用于推导我们稍后讨论的其他上下文。

例如：
```go
    ctx, cancel := context.Background()
```
#### context.TODO() ctx Context

此函数返回空上下文。这应仅用于高级别（在主要或顶级请求处理程序中）。这可以用于推导我们稍后讨论的其他上下文

```go
    ctx, cancel := context.Background()
```

#### context.TODO() ctx Context

此函数还会创建一个空的上下文。 这也应该仅用于高级别，或者当您不确定要使用的上下文或者是否未更新函数以接收上下文时。 这意味着您（或维护者）计划在将来为该功能添加上下文。

```go
ctx, cancel := context.TODO()
```
有趣的是，查看代码（https://golang.org/src/context/context.go），它与背景完全相同。 不同的是，静态分析工具可以使用它来验证上下文是否正确传递，这是一个重要的细节，因为静态分析工具可以帮助在早期发现潜在的错误，并且可以连接到CI中 / CD管道。
From https://golang.org/src/context/context.go:

```go
 var ( background = new(emptyCtx) todo = new(emptyCtx) )
```

#### context.WithValue(parent Context, key, val interface{}) (ctx Context, cancel CancelFunc)

此函数接受上下文并返回派生上下文，其中值val与key关联，并通过上下文树与上下文一起流动。 这意味着一旦获得带有值的上下文，从中派生的任何上下文都会获得此值。 不建议使用上下文值传递关键参数，而是函数应接受签名中的那些值，使其显式化。

ctx := context.WithValue(context.Background(), key, "test")

#### context.WithCancel(parent Context) (ctx Context, cancel CancelFunc)
这是它开始变得有趣的地方。此函数创建从传入的父上下文派生的新上下文。父可以是后台上下文或传递给函数的上下文。


这将返回派生上下文和取消功能。 只有创建它的函数才应调用cancel函数来取消此上下文。 如果您愿意，可以传递取消功能，但是，强烈建议不要这样做。 这可能导致取消的调用者没有意识到取消上下文的下游影响可能是什么。 可能存在源自此的其他上下文，这可能导致程序以意外的方式运行。 简而言之，永远不要传递取消功能。

```go
        ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2 * time.Second))
```


#### context.WithDeadline(parent Context, d time.Time) (ctx Context, cancel CancelFunc)
 
此函数返回其父项的派生上下文，当截止日期超过或取消函数被取消时，该上下文将被取消。 例如，您可以创建一个将在以后的某个时间自动取消的上下文，并在子函数中传递它。 当因为截止日期耗尽而取消该上下文时，获取上下文的所有函数都会收到通知以停止工作并返回。

ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2 * time.Second))

#### context.WithTimeout(parent Context, timeout time.Duration) (ctx Context, cancel CancelFunc)

此函数类似于context.WithDeadline。 不同之处在于它将持续时间作为输入而不是时间对象。 此函数返回派生上下文，如果调用cancel函数或超出超时持续时间，则会取消该派生上下文。

```go
    ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2 * time.Second))
```

#### 在您的函数中接受和使用上下文

现在我们知道了如何创建上下文（Background和TODO）以及如何派生上下文（WithValue，WithCancel，Deadline和Timeout），让我们讨论如何使用它们。

在下面的示例中，您可以看到接受上下文的函数启动goroutine并等待返回该goroutine或取消该上下文。 select语句帮助我们选择先发生的任何情况并返回。

<-ctx.Done（）完成Done通道后，选择<-ctx.Done（）：大小写。 一旦发生这种情况，该功能应该放弃工作并准备返回。 这意味着您应该关闭所有打开的管道，释放资源并从函数返回。 有些情况下，释放资源可以阻止返回，比如做一些挂起的清理等等。在处理上下文返回时，你应该注意任何这样的可能性。

本节后面的示例有一个完整的go程序，它说明了超时和取消功能。

```go 
//Function that does slow processing with a context
//Note that context is the first argument
func sleepRandomContext(ctx context.Context, ch chan bool) {
//Cleanup tasks
  //There are no contexts being created here
  //Hence, no canceling needed
  defer func() {
    fmt.Println("sleepRandomContext complete")
    ch <- true
  }()
//Make a channel
  sleeptimeChan := make(chan int)
//Start slow processing in a goroutine
  //Send a channel for communication
  go sleepRandom("sleepRandomContext", sleeptimeChan)
//Use a select statement to exit out if context expires
  select {
  case <-ctx.Done():
    //If context expires, this case is selected
    //Free up resources that may no longer be needed because of aborting the work
    //Signal all the goroutines that should stop work (use channels)
    //Usually, you would send something on channel,
    //wait for goroutines to exit and then return
    //Or, use wait groups instead of channels for synchronization
    fmt.Println("Time to return")
  case sleeptime := <-sleeptimeChan:
    //This case is selected when processing finishes before the context is cancelled
    fmt.Println("Slept for ", sleeptime, "ms")
  }
}
```

#### Example

到目前为止，我们已经看到使用上下文可以设置截止日期，超时或调用cancel函数来通知所有使用任何派生上下文的函数来停止工作和返回。以下是它如何工作的示例：

main function:

- 使用cancel创建上下文
- 随机超时后调用取消功能

doWorkContext function

- 派生超时上下文
- 此上下文将在取消时取消
- 主要调用cancelFunction或
- 超时过去或
- doWorkContext调用其cancelFunction
- 启动goroutine以通过派生上下文执行一些缓慢的处理
- 等待goroutine完成或上下文被主要取消，以先发生者为准

sleepRandomContext function

- 启动goroutine来进行缓慢处理
- 等待goroutine完成或，
- 等待主要，超时或取消调用其自己的cancelFunction取消上下文


sleepRandom function
- 睡觉时间随机 此示例使用sleep来模拟随机处理时间，
- 在实际示例中，您可以使用通道来通知此函数以开始清理并在通道上等待它以确认清理已完成。

游乐场：https：//play.golang.org/p/grQAUN3MBlg（看起来像我使用的随机种子，时间，在游乐场中并没有真正改变。你可能必须在本地机器上执行此操作以查看随机性）


Github: https://github.com/pagnihotry/golang_samples/blob/master/go_context_sample.go
```go
package main
import (
  "context"
  "fmt"
  "math/rand"
  "time"
)
//Slow function
func sleepRandom(fromFunction string, ch chan int) {
  //defer cleanup
  defer func() { fmt.Println(fromFunction, "sleepRandom complete") }()
//Perform a slow task
  //For illustration purpose,
  //Sleep here for random ms
  seed := time.Now().UnixNano()
  r := rand.New(rand.NewSource(seed))
  randomNumber := r.Intn(100)
  sleeptime := randomNumber + 100
  fmt.Println(fromFunction, "Starting sleep for", sleeptime, "ms")
  time.Sleep(time.Duration(sleeptime) * time.Millisecond)
  fmt.Println(fromFunction, "Waking up, slept for ", sleeptime, "ms")
//write on the channel if it was passed in
  if ch != nil {
    ch <- sleeptime
  }
}
//Function that does slow processing with a context
//Note that context is the first argument
func sleepRandomContext(ctx context.Context, ch chan bool) {
//Cleanup tasks
  //There are no contexts being created here
  //Hence, no canceling needed
  defer func() {
    fmt.Println("sleepRandomContext complete")
    ch <- true
  }()
//Make a channel
  sleeptimeChan := make(chan int)
//Start slow processing in a goroutine
  //Send a channel for communication
  go sleepRandom("sleepRandomContext", sleeptimeChan)
//Use a select statement to exit out if context expires
  select {
  case <-ctx.Done():
    //If context is cancelled, this case is selected
    //This can happen if the timeout doWorkContext expires or
    //doWorkContext calls cancelFunction or main calls cancelFunction
    //Free up resources that may no longer be needed because of aborting the work
    //Signal all the goroutines that should stop work (use channels)
    //Usually, you would send something on channel, 
    //wait for goroutines to exit and then return
    //Or, use wait groups instead of channels for synchronization
    fmt.Println("sleepRandomContext: Time to return")
  case sleeptime := <-sleeptimeChan:
    //This case is selected when processing finishes before the context is cancelled
    fmt.Println("Slept for ", sleeptime, "ms")
  }
}
//A helper function, this can, in the real world do various things.
//In this example, it is just calling one function.
//Here, this could have just lived in main
func doWorkContext(ctx context.Context) {
//Derive a timeout context from context with cancel
  //Timeout in 150 ms
  //All the contexts derived from this will returns in 150 ms
  ctxWithTimeout, cancelFunction := context.WithTimeout(ctx, time.Duration(150)*time.Millisecond)
//Cancel to release resources once the function is complete
  defer func() {
    fmt.Println("doWorkContext complete")
    cancelFunction()
  }()
//Make channel and call context function
  //Can use wait groups as well for this particular case
  //As we do not use the return value sent on channel
  ch := make(chan bool)
  go sleepRandomContext(ctxWithTimeout, ch)
//Use a select statement to exit out if context expires
  select {
  case <-ctx.Done():
    //This case is selected when the passed in context notifies to stop work
    //In this example, it will be notified when main calls cancelFunction
    fmt.Println("doWorkContext: Time to return")
  case <-ch:
    //This case is selected when processing finishes before the context is cancelled
    fmt.Println("sleepRandomContext returned")
  }
}
func main() {
  //Make a background context
  ctx := context.Background()
  //Derive a context with cancel
  ctxWithCancel, cancelFunction := context.WithCancel(ctx)
//defer canceling so that all the resources are freed up 
  //For this and the derived contexts
  defer func() {
    fmt.Println("Main Defer: canceling context")
    cancelFunction()
  }()
//Cancel context after a random time
  //This cancels the request after a random timeout
  //If this happens, all the contexts derived from this should return
  go func() {
    sleepRandom("Main", nil)
    cancelFunction()
    fmt.Println("Main Sleep complete. canceling context")
  }()
  //Do work
  doWorkContext(ctxWithCancel)
}
```

### Gotchas
如果函数接受上下文，请确保检查它是如何遵循取消通知的。 例如，exec.CommandContext不会关闭读取器管道，直到命令执行了进程创建的所有分支（Github问题：https：//github.com/golang/go/issues/23019），这意味着 如果等待cmd.Wait（）直到外部命令的所有分支都已完成处理，则上下文取消不会立即返回此函数。 如果您使用超时或最后执行时间的最后期限，您可能会发现这不能按预期工作。 如果遇到任何此类问题，可以使用time.After实现超时。


#### Best practices
- context.Background只应在最高级别使用，作为所有派生上下文的根
- context.TODO应该用在不确定要使用什么的地方，或者是否将更新当前函数以便将来使用上下文
- 上下文取消是建议性的，功能可能需要一些时间来清理和退出
- context.Value应该很少使用，它永远不应该用于传递可选参数。 这使得API隐含并且可能引入错误。 相反，这些值应作为参数递。
- 不要在结构中存储上下文，在函数中显式传递它们，最好是作为第一个参数。
- 从不传递nil上下文，相反，如果您不确定要使用什么，请使用TODO。
- Context结构没有cancel方法，因为只有派生上下文的函数才能取消它。


[1]:
https://medium.com/@parikshit/understanding-the-context-package-in-golang-b1392c821d14