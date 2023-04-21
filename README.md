# protoc-gen-grpc-gateway-gosdk

[![GitHub release](https://img.shields.io/github/release/jaronnie/protoc-gen-grpc-gateway-gosdk.svg?style=flat-square)](https://github.com/jaronnie/protoc-gen-grpc-gateway-gosdk/releases/latest)
[![Build Status](https://img.shields.io/github/actions/workflow/status/jaronnie/protoc-gen-grpc-gateway-gosdk/ci.yaml?branch=main&label=protoc-gen-grpc-gateway-gosdk-golint&logo=github&style=flat-square)](https://github.com/jaronnie/protoc-gen-grpc-gateway-gosdk/actions?query=workflow%3Aprotoc-gen-grpc-gateway-gosdk-golint)
[![Build Status](https://img.shields.io/github/actions/workflow/status/jaronnie/protoc-gen-grpc-gateway-gosdk/ci.yaml?branch=main&label=goreleaser-protoc-gen-grpc-gateway-gosdk&logo=github&style=flat-square)](https://github.com/jaronnie/protoc-gen-grpc-gateway-gosdk/actions?query=workflow%3Agoreleaser-protoc-gen-grpc-gateway-gosdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/jaronnie/protoc-gen-grpc-gateway-gosdk?style=flat-square)](https://goreportcard.com/report/github.com/jaronnie/protoc-gen-grpc-gateway-gosdk)
[![codecov](https://img.shields.io/codecov/c/github/jaronnie/protoc-gen-grpc-gateway-gosdk?logo=codecov&style=flat-square)](https://codecov.io/gh/jaronnie/protoc-gen-grpc-gateway-gosdk)

According to the function of grpc gateway supported by grpc-gateway plugin

to generate go http sdk similar to k8s.

## install

make sure you have already installed [task](https://github.com/go-task/task) and [goreleaser](https://github.com/goreleaser/goreleaser) tools first.

```shell
git clone https://github.com/jaronnie/protoc-gen-grpc-gateway-gosdk.git
cd protoc-gen-grpc-gateway-gosdk
task install
```

## _example

make sure you have already installed protoc and protoc-gen-go.

```shell
cd _example/singleScopeVersionSingleService
make gensdk.debug.fmt
cd ../..
```
