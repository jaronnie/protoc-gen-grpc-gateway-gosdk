// Code generated by protoc-gen-grpc-gateway-gosdk. DO NOT EDIT.
// versions:
//    protoc-gen-grpc-gateway-gosdk 1.8.1-next-f3e4de
// type: fake_clientset

package fake

import (
	"github.com/jaronnie/autosdk/typed"
	av1 "github.com/jaronnie/autosdk/typed/av1"
	fakeav1 "github.com/jaronnie/autosdk/typed/av1/fake"
	bv1 "github.com/jaronnie/autosdk/typed/bv1"
	fakebv1 "github.com/jaronnie/autosdk/typed/bv1/fake"
	"github.com/jaronnie/autosdk/typed/fake"
)

type Clientset struct{}

func (f *Clientset) Direct() typed.DirectInterface {
	return &fake.FakeDirect{}
}

func (f *Clientset) Av1() av1.Av1Interface {
	return &fakeav1.FakeAv1{}
}

func (f *Clientset) Bv1() bv1.Bv1Interface {
	return &fakebv1.FakeBv1{}
}