# nsq 学习

NSQ is a realtime distributed messaging platform designed to operate at scale, handling billions of messages per day.

NSQ是一个实时分布式消息传递平台，旨在大规模运行，每天处理数十亿条消息

It promotes distributed and decentralized topologies without single points of failure, enabling fault tolerance and high availability coupled with a reliable message delivery guarantee. See features & guarantees.

它可以在没有单点故障的情况下促进分布式和分散式拓扑，实现容错和高可用性以及可靠的消息传递保证。查看功能和保证。

Operationally, NSQ is easy to configure and deploy (all parameters are specified on the command line and compiled binaries have no runtime dependencies). For maximum flexibility, it is agnostic to data format (messages can be JSON, MsgPack, Protocol Buffers, or anything else). Official Go and Python libraries are available out of the box (as well as many other client libraries) and, if you're interested in building your own, there's a protocol spec.

在操作上, NSQ 易于配置和部署 (在命令行上指定了所有参数, 并且编译的二进制文件没有运行时依赖项)。为了获得最大的灵活性, 数据格式 (消息可以是 JSON、MsgPack、协议缓冲区或其他任何内容) 是不可知的。正式的去和 Python 库可用的框 (以及许多其他客户端库) 和 如果你有兴趣建立自己的，有一个协议规范。

We publish binary releases for linux, darwin, freebsd and windows as well as an official Docker image.

我们发布了linux，darwin，freebsd和windows的二进制版本以及官方Docker镜像

NOTE: master is our development branch and may not be stable at all times.
注: master 是我们的发展分支, 可能在任何时候都不稳定。

```go
fmt





```