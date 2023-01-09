package main

import (
	"context"
	"fmt"

	"github.com/jaronnie/autosdk"
	"github.com/jaronnie/autosdk/pb/corev1"
	"github.com/jaronnie/autosdk/rest"
)

func main() {
	var clientSet autosdk.Interface
	var err error
	clientSet, err = autosdk.NewClientWithOptions(
		rest.WithProtocol("http"),
		rest.WithAddr("127.0.0.1"),
		rest.WithPort("8090"),
		rest.WithHeaders(map[string][]string{
			"Content-Type": {"application/json"},
		}),
		rest.WithGatewayPrefix("/gateway"),
	)

	if err != nil {
		panic(err)
	}

	machine, err := InitCredential(clientSet)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(machine)
}

func InitCredential(clientSet autosdk.Interface) (*corev1.Credential, error) {
	return clientSet.Corev1().Credential().InitCredential(context.Background(), &corev1.Credential{})
}
