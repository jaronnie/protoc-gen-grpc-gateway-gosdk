package main

import (
	"testing"

	"github.com/jaronnie/autosdk/fake"
	"github.com/jaronnie/autosdk/pb/corev1"
	fakecorev1 "github.com/jaronnie/autosdk/typed/corev1/fake"
)

func TestInitMachine(t *testing.T) {
	t.Run("test init credential", func(t *testing.T) {
		fakecorev1.FakeReturnInitCredential = &corev1.Credential{
			Id:   2,
			Name: "colocation",
		}
		fakeClient := &fake.Clientset{}
		machine, err := InitCredential(fakeClient)
		if err != nil {
			t.Log(err.Error())
			return
		}
		t.Log(machine)
	})
}
