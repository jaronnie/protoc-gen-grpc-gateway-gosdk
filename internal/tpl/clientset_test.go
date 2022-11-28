package tpl

import (
	"fmt"
	"testing"

	"github.com/jaronnie/protoc-gen-go-httpsdk/utilx"
)

func TestGenerateClientSetFile(t *testing.T) {
	clientSetData := &ClientSetData{
		GoModule:      "github.com/jaronnie/autosdk",
		RootModule:    "bfsdk",
		ScopeVersions: map[string]string{"corev1": "CoreV1", "oauthv1": "OauthV1"},
	}

	template, err := utilx.ParseTemplate(clientSetData, []byte(ClientSetTpl))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(template))
}
