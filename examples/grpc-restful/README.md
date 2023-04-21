# 入门简单例子, 以 go-zero 的 grpc-restful 为例

## 安装工具

```shell
go install github.com/zeromicro/go-zero/tools/goctl@latest
go install github.com/golang/protobuf/protoc-gen-go@v1.3.2
```

## 生成 gosdk

```shell
# 生成的 gosdk 直接在服务端
mkdir -p pkgsdk/pb
protoc -I./proto --go_out=./pkgsdk/pb --grpc-gateway-gosdk_out=logtostderr=true,v=1,env_file=etc/pkgsdk.yaml:pkgsdk proto/user.proto

# 生成的 gosdk 独立 module
mkdir -p modsdk/pb
protoc -I./proto --go_out=./modsdk/pb --grpc-gateway-gosdk_out=logtostderr=true,v=1,env_file=etc/modsdk.yaml:modsdk proto/user.proto
```
