# protoc-gen-go-httpsdk


[![GitHub release](https://img.shields.io/github/release/jaronnie/protoc-gen-go-httpsdk.svg?style=flat-square)](https://github.com/jaronnie/protoc-gen-go-httpsdk/releases/latest)
[![Build Status](https://img.shields.io/github/actions/workflow/status/jaronnie/protoc-gen-go-httpsdk/ci.yaml?branch=main&label=protoc-gen-go-httpsdk-golint&logo=github&style=flat-square)](https://github.com/jaronnie/protoc-gen-go-httpsdk/actions?query=workflow%3Aprotoc-gen-go-httpsdk-golint)
[![Build Status](https://img.shields.io/github/actions/workflow/status/jaronnie/protoc-gen-go-httpsdk/ci.yaml?branch=main&label=goreleaser-protoc-gen-go-httpsdk&logo=github&style=flat-square)](https://github.com/jaronnie/protoc-gen-go-httpsdk/actions?query=workflow%3Agoreleaser-protoc-gen-go-httpsdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/jaronnie/protoc-gen-go-httpsdk?style=flat-square)](https://goreportcard.com/report/github.com/jaronnie/protoc-gen-go-httpsdk)
[![codecov](https://img.shields.io/codecov/c/github/jaronnie/protoc-gen-go-httpsdk?logo=codecov&style=flat-square)](https://codecov.io/gh/jaronnie/protoc-gen-go-httpsdk)

According to the function of grpc gateway supported by grpc-gateway plugin

to generate go http sdk similar to k8s.

## install 

make sure you have already installed [task](https://github.com/go-task/task) and [goreleaser](https://github.com/goreleaser/goreleaser) tools first.

```shell
git clone https://github.com/jaronnie/protoc-gen-go-httpsdk.git
cd protoc-gen-go-httpsdk
task install
```

## _example

make sure you have already installed protoc and protoc-gen-go.

```shell
cd _example/singleScopeVersionSingleService
make gensdk.debug.fmt
cd ../..
```