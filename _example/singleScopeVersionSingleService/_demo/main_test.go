package main

import (
	"testing"

	"github.com/jaronnie/autosdk/fake"
	"github.com/jaronnie/autosdk/pb/corev1"
	fakeMachine "github.com/jaronnie/autosdk/typed/corev1/fake"
)

func TestInitMachine(t *testing.T) {
	fakeMachine.Machines = &corev1.Machine{
		Id:   2,
		Type: "colocation",
	}
	t.Run("test init machine", func(t *testing.T) {
		fakeClient := &fake.Clientset{}
		machine, err := InitMachine(fakeClient)
		if err != nil {
			t.Log(err.Error())
			return
		}
		t.Log(machine)
	})
}
