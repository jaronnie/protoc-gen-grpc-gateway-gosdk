// Code generated by protoc-gen-grpc-gateway-go. DO NOT EDIT.
// versions:
//    protoc-gen-grpc-gateway-go 1.7.1-next-5e31ea
// type: fake_machine

package fake

import (
	"github.com/jaronnie/autosdk/pb/corev1"
	"github.com/jaronnie/autosdk/rest"
)

var (
	FakeReturnInitMachine     = &rest.Request{}
	FakeReturnDownloadMachine = &rest.Request{}
)

type MachineGetter interface {
	Machine() MachineInterface

	FakeMachineExpansion
}

type MachineInterface interface {
	InitMachine(param *corev1.Machine) (*rest.Request, error)
	DownloadMachine(param *corev1.Machine) (*rest.Request, error)
}

type FakeMachine struct {
	Fake *FakeCorev1
}

func (f *FakeMachine) InitMachine(param *corev1.Machine) (*rest.Request, error) {
	return FakeReturnInitMachine, nil
}

func (f *FakeMachine) DownloadMachine(param *corev1.Machine) (*rest.Request, error) {
	return FakeReturnDownloadMachine, nil
}
