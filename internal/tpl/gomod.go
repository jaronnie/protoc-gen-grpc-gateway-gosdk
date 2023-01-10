package tpl

type GoModData struct {
	GoModule  string
	GoVersion string
}

var GoModTpl = `module {{.GoModule}}

go {{.GoVersion}}

require (
	github.com/bitly/go-simplejson v0.5.0
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/websocket v1.5.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cast v1.5.0
	google.golang.org/genproto v0.0.0-20230106154932-a12b697841d9
)

require (
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	google.golang.org/grpc v1.51.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

`
