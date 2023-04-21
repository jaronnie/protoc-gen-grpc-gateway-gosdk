package utilx

import (
	"strings"

	"github.com/jaronnie/protoc-gen-grpc-gateway-go/utilx/execx"
)

func getProtoToolsVersion(t string) (string, error) {
	// todo
	// add context to cancel

	switch t {
	case "protoc":
		return getProtocVersion()
	case "protoc-gen-go":
		return getProtocGenGoVersion()
	case "protoc-gen-grpc-gateway-go":
		return getProtocGenGRPCGatewayVersion()
	}
	return "", nil
}

func getProtocVersion() (string, error) {
	version, err := execx.Run("protoc --version", "")
	if err != nil {
		return "", err
	}
	fields := strings.Fields(version)
	if len(fields) > 1 {
		return fields[1], nil
	}
	return "", nil
}

func getProtocGenGoVersion() (string, error) {
	version, err := execx.Run("protoc-gen-go --version", "")
	if err != nil {
		return "", err
	}
	fields := strings.Fields(version)
	if len(fields) > 1 {
		return fields[1], nil
	}
	return "", nil
}

func getProtocGenGRPCGatewayVersion() (string, error) {
	version, err := execx.Run("protoc-gen-grpc-gateway-go version", "")
	if err != nil {
		return "", err
	}

	return version, nil
}
