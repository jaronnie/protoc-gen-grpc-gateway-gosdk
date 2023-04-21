package typed

import (
	"fmt"
	"testing"

	"github.com/jaronnie/protoc-gen-grpc-gateway-gosdk/utilx"
)

func TestGenerateScopeClientFile(t *testing.T) {
	scopeClientData := ScopeClientData{
		ScopeVersion:   "corev1",
		UpScopeVersion: "Corev1",
		GoModule:       "github.com/jaronnie/autosdk",
		UpResources:    []string{"Credential", "Machine", "Chain"},
	}
	template, err := utilx.ParseTemplate(scopeClientData, []byte(ScopeClientTpl))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(template))
}
