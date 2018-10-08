# 设计模式

## 1.单例模式

* [Golang中的设计模式：Singleton](http://blog.ralch.com/tutorial/design-patterns/golang-singleton/)

* In software engineering, the singleton pattern is a software design pattern that restricts the instantiation of a class to one object.

举个栗子： 
Windows 是多进程多线程的，在操作一个文件的时候，就不可避免地出现多个进程或线程同时操作一个文件的现象，所以所有文件的处理必须通过唯一的实例来进行。这就可以使用单例模式来解决该问题。



* golang中的单例
但是，在golang的世界中，没有private public static等关键字，也没有面向对象中类的概念。

* 那么golang是如何控制访问范围的呢？ 
首字母大写，代表对外部可见，首字母小写代表对外部不可见，适用于所有对象，包括函数、方法

* golang中全局变量 
可以使用全局变量，达到c++中static的效果

* golang标准库中的单例模式使用示例 
golang是开源的，当我们不知道怎么写代码的时候，完全可以读一读golang的源码。 
对于单例模式，例如net/http包中的 http.DefaultClient 和 http.DefaultServeMux。

http.DefaultClient
```go 
    type Client struct {
   
       Transport RoundTripper
   
       CheckRedirect func(req *Request, via []*Request) error
   
       Jar CookieJar
   
       Timeout time.Duration
     }
   var DefaultClient = &Client{}
```

http.DefaultServeMux

```go
// DefaultServeMux is the default ServeMux used by Serve.
  var DefaultServeMux = &defaultServeMux
```

可以通过这些代码查看信息理解singleton 

```go
package db

import "fmt"

type repository struct {
    items map[string]string
    mu    sync.RWMutex
}

func (r *repository) Set(key, data string) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.items[key] = data
}

func (r *repository) Get(key string) (string, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    item, ok := r.items[key]
    if !ok {
        return "", fmt.Errorf("The '%s' is not presented", key)
    }
    return item, nil
}

var (
    r    *repository
    once sync.Once
)

func Repository() *repository {
    once.Do(func() {
        r = &repository{
            items: make(map[string]string),
        }
    })

    return r
}
```