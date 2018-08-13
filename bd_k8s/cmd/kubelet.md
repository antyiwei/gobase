


# kubelet

* kubelet大致说明

kubelet是在每个上运行的主要“节点代理”节点。 kubelet适用于PodSpec。 PodSpec是YAML或JSON对象描述一个pod。 kubelet采用通过提供的一组PodSpecs各种机制（主要通过apiserver）并确保容器在那些PodSpecs中描述的运行和健康。 kubelet无法管理不是由Kubernetes创造的容器。 

除了来自apiserver的PodSpec之外，容器有三种方式清单可以提供给Kubelet。

File：Path在命令行上作为标志传递。将监视此路径下的文件定期更新。监控时间默认为20秒，可配置通过一面旗帜。

HTTP端点：在命令行上作为参数传递的HTTP端点。这个端点每20秒检查一次（也可以用标志配置）。

HTTP服务器：kubelet还可以侦听HTTP并响应简单的API（目前未提及）提交新清单。

* kubelet源码分析流程

1. main入口：k8s.io\kubernetes\cmd\kubelet\kubelet.go

```go
    func main() {
        rand.Seed(time.Now().UTC().UnixNano()) // 获取一个随机数
    
        command := app.NewKubeletCommand(server.SetupSignalHandler()) //初始化一个kubeletCommand的入口
        logs.InitLogs() // 初始化日志
        defer logs.FlushLogs()// 日志结束
    
        if err := command.Execute(); err != nil { // command初始化后，执行
            fmt.Fprintf(os.Stderr, "%v\n", err)
            os.Exit(1)
        }
    }
```
这个main函数是不是大家看着很亲切的感觉，我有这种感觉，就像发现新大陆一样

既然有这个main这个主入口，我们就分析代码，上面有注释，这里主要学习 NewKubeletCommand 函数


#### 参考文件

[k8s源代码分析-----kubelet（1）主要流程](https://www.cnblogs.com/slgkaifa/p/7308368.html)