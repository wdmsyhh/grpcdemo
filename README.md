# grpcdemo

## 生成 pb.go 文件

```shell
./gen-stub
```

## 启动服务

```shell
cd service
go run main.go
```

## 访问接口

```shell
# 注意：这里的端口是 gateway 的端口，会代理请求到 grpc 服务
http://localhost:8080/v1/example/hello/get
```