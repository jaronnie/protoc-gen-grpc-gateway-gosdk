package fake

import (
	"github.com/jaronnie/autosdk/typed"
	"github.com/jaronnie/autosdk/typed/corev1"
	"github.com/jaronnie/autosdk/typed/corev1/fake"
)

type Clientset struct {
	direct *typed.DirectClient
}

func (x *Clientset) Direct() typed.DirectInterface {
	return x.direct
}

func (x *Clientset) Corev1() corev1.Corev1Interface {
	return &fake.FakedCorev1{}
}
