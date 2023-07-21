package fake

type DirectClientData struct {
	GoModule string // github.com/jaronnie/autosdk
}

var FakeDirectClientTpl = `
// Code generated by protoc-gen-grpc-gateway-gosdk. DO NOT EDIT.
// versions:
//    protoc-gen-grpc-gateway-gosdk {{getProtoToolsVersion "protoc-gen-grpc-gateway-gosdk"}}
// type: fake_direct_client

package fake

import (
	"{{.GoModule}}/rest"
)

type FakeDirect struct {}

func (f *FakeDirect) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
`