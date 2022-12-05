package fake

import (
	"github.com/jaronnie/autosdk/rest"
	"github.com/jaronnie/autosdk/typed/corev1"
)

type FakedCorev1 struct {
}

func (f *FakedCorev1) Credential() corev1.CredentialInterface {
	return &FakedCredential{Fake: f}
}

func (f *FakedCorev1) Machine() corev1.MachineInterface {
	return &FakedMachine{Fake: f}
}

func (f *FakedCorev1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
