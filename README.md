# protoc-gen-go-httpsdk

According to the function of grpc gateway supported by grpc-gateway plugin

to generate go http sdk similar to k8s.

## install 

```shell
git clone https://github.com/jaronnie/protoc-gen-go-httpsdk.git
cd protoc-gen-go-httpsdk
make build
```

## _example

make sure you have already installed protoc and protoc-gen-go.

```shell
cd _example/singleScopeVersionSingleService
make gensdk.debug.fmt
cd ../..
```