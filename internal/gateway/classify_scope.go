package gateway

import (
	"github.com/jaronnie/protoc-gen-grpc-gateway-go/internal/vars"
	"github.com/jaronnie/protoc-gen-grpc-gateway-go/utilx"
)

// GetAllScopeVersionsMap {"corev1":"CoreV1", "corev2":"CoreV2"}
func GetAllScopeVersionsMap(scopeResourceGws vars.ScopeResourceGateway) map[string]string {
	scopeVersionMap := make(map[string]string, 0)

	for k := range scopeResourceGws {
		scopeVersionMap[string(k)] = utilx.FirstUpper(string(k))
	}
	return scopeVersionMap
}
