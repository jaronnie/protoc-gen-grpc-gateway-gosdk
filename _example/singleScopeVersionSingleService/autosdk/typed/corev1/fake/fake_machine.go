package fake

import (
	"context"

	"github.com/jaronnie/autosdk/pb/corev1"
)

var (
	Machines = &corev1.Machine{}
)

type FakedMachine struct {
	Fake *FakedCorev1
}

type MachineGetter interface {
	Machine() MachineInterface
}

type MachineInterface interface {
	InitMachine(ctx context.Context, param *corev1.Machine) (*corev1.Machine, error)
}

func (x *FakedMachine) InitMachine(ctx context.Context, param *corev1.Machine) (*corev1.Machine, error) {
	return Machines, nil
}
