package typed

type DirectClientData struct {
	GoModule string // github.com/jaronnie/autosdk
}

var DirectClientTpl = `
// Code generated by protoc-gen-grpc-gateway-go. DO NOT EDIT.
// versions:
//    protoc-gen-grpc-gateway-go {{getProtoToolsVersion "protoc-gen-grpc-gateway-go"}}
// type: direct_client

package typed

import (
	"{{.GoModule}}/rest"
)


type DirectInterface interface {
	RESTClient() rest.Interface
}

type DirectClient struct {
	restClient rest.Interface
}

func (x *DirectClient) RESTClient() rest.Interface {
	if x == nil {
		return nil
	}
	return x.restClient
}

// NewForConfig creates a new DirectClient for the given config.
func NewForConfig(x *rest.RESTClient) (*DirectClient, error) {
	config := *x
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &DirectClient{client}, nil
}
`