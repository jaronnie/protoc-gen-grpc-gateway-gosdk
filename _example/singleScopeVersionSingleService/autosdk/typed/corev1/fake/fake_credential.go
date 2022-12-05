package fake

import (
	"context"

	"github.com/jaronnie/autosdk/pb/corev1"
)

type FakedCredential struct {
	Fake *FakedCorev1
}

func (x *FakedCredential) UpdateCredential(ctx context.Context, param *corev1.Credential) (*corev1.Credential, error) {
	return &corev1.Credential{}, nil
}

type CredentialGetter interface {
	Credential() CredentialInterface
}

type CredentialInterface interface {
	InitCredential(ctx context.Context, param *corev1.Credential) (*corev1.Credential, error)
	UpdateCredential(ctx context.Context, param *corev1.Credential) (*corev1.Credential, error)
}

func (x *FakedCredential) InitCredential(ctx context.Context, param *corev1.Credential) (*corev1.Credential, error) {
	return &corev1.Credential{}, nil
}
