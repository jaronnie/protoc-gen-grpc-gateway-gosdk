package main

import (
	"testing"

	"github.com/jaronnie/autosdk/fake"
	"github.com/jaronnie/autosdk/pb/corev1"
	fakecorev1 "github.com/jaronnie/autosdk/typed/corev1/fake"
)

func TestInitMachine(t *testing.T) {
	t.Run("test init machine", func(t *testing.T) {
		fakecorev1.FakeReturnInitMachine = &corev1.Machine{
			Id:   2,
			Type: "colocation",
		}
		fakeClient := &fake.Clientset{}
		machine, err := InitMachine(fakeClient)
		if err != nil {
			t.Log(err.Error())
			return
		}
		t.Log(machine)
	})
}
