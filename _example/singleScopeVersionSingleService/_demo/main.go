package main

import (
	"context"
	"fmt"

	"github.com/jaronnie/autosdk"
	"github.com/jaronnie/autosdk/pb/pb/corev1"
	"github.com/jaronnie/autosdk/rest"
)

var ClientSet *autosdk.Clientset

func main() {
	machine, err := ClientSet.Corev1().Machine().InitMachine(context.Background(), &corev1.Machine{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(machine)
}

func init() {
	var err error
	ClientSet, err = autosdk.NewClientWithOptions(
		rest.WithProtocol("http"),
		rest.WithAddr("127.0.0.1"),
		rest.WithPort("8090"),
		rest.WithHeaders(map[string][]string{
			"Content-Type": {"application/json"},
		}))
	if err != nil {
		panic(err)
	}
}
