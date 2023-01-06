// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package corev1

import (
	"context"

	"github.com/jaronnie/autosdk/pb/corev1"
	"github.com/jaronnie/autosdk/rest"
)

type MachineGetter interface {
	Machine() MachineInterface
}

type MachineInterface interface {
	InitMachine(ctx context.Context, param *corev1.Machine) (*corev1.Machine, error)
	DownloadMachine(param *corev1.Machine) (*rest.Request, error)

	MachineExpansion
}

type machineClient struct {
	client rest.Interface
}

func newMachineClient(c *Corev1Client) *machineClient {
	return &machineClient{
		client: c.RESTClient(),
	}
}

func (x *machineClient) InitMachine(ctx context.Context, param *corev1.Machine) (*corev1.Machine, error) {
	var resp corev1.Machine
	err := x.client.Verb("POST").
		SubPath(
			"/gateway/core/api/v1/machine/init",
		).
		Params().
		Body(param).
		Do(ctx).
		Into(&resp, false)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (x *machineClient) DownloadMachine(param *corev1.Machine) (*rest.Request, error) {
	request := x.client.Verb("POST").
		SubPath(
			"/gateway/core/api/v1/machine/download",
		).
		Params().
		Body(param)

	return request, nil
}
