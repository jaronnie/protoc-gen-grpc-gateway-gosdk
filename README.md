# protoc-gen-grpc-gateway-gosdk

[![GitHub release](https://img.shields.io/github/release/jaronnie/protoc-gen-grpc-gateway-gosdk.svg?style=flat-square)](https://github.com/jaronnie/protoc-gen-grpc-gateway-gosdk/releases/latest)
[![Build Status](https://img.shields.io/github/actions/workflow/status/jaronnie/protoc-gen-grpc-gateway-gosdk/ci.yaml?branch=main&label=protoc-gen-grpc-gateway-gosdk-golint&logo=github&style=flat-square)](https://github.com/jaronnie/protoc-gen-grpc-gateway-gosdk/actions?query=workflow%3Aprotoc-gen-grpc-gateway-gosdk-golint)
[![Build Status](https://img.shields.io/github/actions/workflow/status/jaronnie/protoc-gen-grpc-gateway-gosdk/ci.yaml?branch=main&label=goreleaser-protoc-gen-grpc-gateway-gosdk&logo=github&style=flat-square)](https://github.com/jaronnie/protoc-gen-grpc-gateway-gosdk/actions?query=workflow%3Agoreleaser-protoc-gen-grpc-gateway-gosdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/jaronnie/protoc-gen-grpc-gateway-gosdk?style=flat-square)](https://goreportcard.com/report/github.com/jaronnie/protoc-gen-grpc-gateway-gosdk)
[![codecov](https://img.shields.io/codecov/c/github/jaronnie/protoc-gen-grpc-gateway-gosdk?logo=codecov&style=flat-square)](https://codecov.io/gh/jaronnie/protoc-gen-grpc-gateway-gosdk)

According to the function of grpc gateway supported by grpc-gateway plugin

to generate go http sdk similar to k8s.

```shell
$ tree modsdk
modsdk
├── clientset.go
├── fake
│   └── fake_clientset.go
├── go.mod
├── go.sum
├── pb
│   └── userpb
│       └── user.pb.go
├── rest
│   ├── client.go
│   ├── option.go
│   └── request.go
└── typed
    ├── direct_client.go
    ├── fake
    │   └── fake_direct_client.go
    └── userv1
        ├── fake
        │   ├── fake_user.go
        │   ├── fake_user_expansion.go
        │   └── fake_userv1_client.go
        ├── user.go
        ├── user_expansion.go
        └── userv1_client.go

8 directories, 16 files
```

## install

make sure you have already installed [task](https://github.com/go-task/task) and [goreleaser](https://github.com/goreleaser/goreleaser) tools first.

```shell
git clone https://github.com/jaronnie/protoc-gen-grpc-gateway-gosdk.git
cd protoc-gen-grpc-gateway-gosdk
task install
```

## examples

### grpc-restful

make sure you have already installed protoc and protoc-gen-go.

![2023-04-22_01-10-03](https://oss.jaronnie.com/2023-04-22_01-10-03.gif)
