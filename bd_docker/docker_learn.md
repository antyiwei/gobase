# docker 学习

## 二.容器相关
2.1 最基本的启动
```
docker run -it ubuntu
```
参数-it的含义，可以用docker run --help查看，就不展开了

2.2 启动后执行命令
```
Docker run -it ubuntu echo 'hello world'
```

2.3 启动指定容器名称
```
docker run -it --name 'myubuntu' ubuntu
```
容器名称是一个很有意思的东东，后面马上拼到。上面的命令运行完以后，先用exit退出，以便后面学习其他命令。

2.4 查看最近运行过的所有容器
```
docker ps -a
```