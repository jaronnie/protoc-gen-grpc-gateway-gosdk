# grpc-restful

## 安装工具

```
go install github.com/zeromicro/go-zero/tools/goctl@latest
go install github.com/golang/protobuf/protoc-gen-go@v1.3.2
```

## 编写 proto

```protobuf
syntax = "proto3";
option go_package = "./userpb";
package user;

import "google/api/annotations.proto";

message AddUserReq {
      string name = 1;
      int32 age = 2;
}

message AddUserResp {
      int32 id = 1;
}

service user {
      rpc Add(AddUserReq) returns (AddUserResp) {
            option (google.api.http) = {
                  get: "/api/v1.0/user/add"
            };
      };
}
```

## 编写配置文件

* modsdk.yaml

```yaml
scopeVersion: userv1
goModule: github.com/jaronnie/autosdk
goVersion: 1.18
```

* pkgsdk.yaml

```shell
scopeVersion: userv1
sdkDir: pkgsdk
```

## 生成 gosdk

```shell
git clone https://github.com/jaronnie/protoc-gen-grpc-gateway-gosdk.git
cd protoc-gen-grpc-gateway-gosdk
task install

cd examples/grpc-restful
# 生成的 gosdk 直接在服务端
mkdir -p pkgsdk/pb
protoc -I./proto --go_out=./pkgsdk/pb --grpc-gateway-gosdk_out=logtostderr=true,v=1,env_file=etc/pkgsdk.yaml:pkgsdk proto/user.proto

# 生成的 gosdk 独立 module
mkdir -p modsdk/pb
protoc -I./proto --go_out=./modsdk/pb --grpc-gateway-gosdk_out=logtostderr=true,v=1,env_file=etc/modsdk.yaml:modsdk proto/user.proto
```

