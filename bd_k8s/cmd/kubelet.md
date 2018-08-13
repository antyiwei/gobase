


# kubelet
kubelet是在每个上运行的主要“节点代理”
节点。 kubelet适用于PodSpec。 PodSpec是YAML或JSON对象
描述一个pod。 kubelet采用通过提供的一组PodSpecs
各种机制（主要通过apiserver）并确保容器
在那些PodSpecs中描述的运行和健康。 kubelet无法管理
不是由Kubernetes创造的容器。

除了来自apiserver的PodSpec之外，容器有三种方式
清单可以提供给Kubelet。

File：Path在命令行上作为标志传递。将监视此路径下的文件
定期更新。监控时间默认为20秒，可配置
通过一面旗帜。

HTTP端点：在命令行上作为参数传递的HTTP端点。这个端点
每20秒检查一次（也可以用标志配置）。

HTTP服务器：kubelet还可以侦听HTTP并响应简单的API
（目前未提及）提交新清单。
