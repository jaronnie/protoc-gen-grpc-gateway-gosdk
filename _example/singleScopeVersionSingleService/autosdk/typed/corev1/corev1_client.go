// Code generated by protoc-gen-grpc-gateway-go. DO NOT EDIT.
// versions:
//    protoc-gen-grpc-gateway-go 1.7.1-next-bbbf72
// type: corev1_client

package corev1

import (
	"github.com/jaronnie/autosdk/rest"
)

type Corev1Interface interface {
	RESTClient() rest.Interface

	CredentialGetter
	MachineGetter
}

type Corev1Client struct {
	restClient rest.Interface
}

func (x *Corev1Client) RESTClient() rest.Interface {
	if x == nil {
		return nil
	}
	return x.restClient
}

func (x *Corev1Client) Credential() CredentialInterface {
	return newCredentialClient(x)
}

func (x *Corev1Client) Machine() MachineInterface {
	return newMachineClient(x)
}

// NewForConfig creates a new Corev1Client for the given config.
func NewForConfig(x *rest.RESTClient) (*Corev1Client, error) {
	config := *x
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &Corev1Client{client}, nil
}
