# protoc-gen-grpc-gateway-go

[![GitHub release](https://img.shields.io/github/release/jaronnie/protoc-gen-grpc-gateway-go.svg?style=flat-square)](https://github.com/jaronnie/protoc-gen-grpc-gateway-go/releases/latest)
[![Build Status](https://img.shields.io/github/actions/workflow/status/jaronnie/protoc-gen-grpc-gateway-go/ci.yaml?branch=main&label=protoc-gen-grpc-gateway-go-golint&logo=github&style=flat-square)](https://github.com/jaronnie/protoc-gen-grpc-gateway-go/actions?query=workflow%3Aprotoc-gen-grpc-gateway-go-golint)
[![Build Status](https://img.shields.io/github/actions/workflow/status/jaronnie/protoc-gen-grpc-gateway-go/ci.yaml?branch=main&label=goreleaser-protoc-gen-grpc-gateway-go&logo=github&style=flat-square)](https://github.com/jaronnie/protoc-gen-grpc-gateway-go/actions?query=workflow%3Agoreleaser-protoc-gen-grpc-gateway-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/jaronnie/protoc-gen-grpc-gateway-go?style=flat-square)](https://goreportcard.com/report/github.com/jaronnie/protoc-gen-grpc-gateway-go)
[![codecov](https://img.shields.io/codecov/c/github/jaronnie/protoc-gen-grpc-gateway-go?logo=codecov&style=flat-square)](https://codecov.io/gh/jaronnie/protoc-gen-grpc-gateway-go)

According to the function of grpc gateway supported by grpc-gateway plugin

to generate go http sdk similar to k8s.

## install

make sure you have already installed [task](https://github.com/go-task/task) and [goreleaser](https://github.com/goreleaser/goreleaser) tools first.

```shell
git clone https://github.com/jaronnie/protoc-gen-grpc-gateway-go.git
cd protoc-gen-grpc-gateway-go
task install
```

## _example

make sure you have already installed protoc and protoc-gen-go.

```shell
cd _example/singleScopeVersionSingleService
make gensdk.debug.fmt
cd ../..
```
