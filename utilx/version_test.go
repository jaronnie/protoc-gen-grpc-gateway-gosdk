package utilx

import (
	"fmt"
	"testing"
)

func TestGetProtocVersion(t *testing.T) {
	s, err := getProtocVersion()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}

func TestGetProtocGenGoVersion(t *testing.T) {
	s, err := getProtocGenGoVersion()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}

func TestGetProtocGenGRPCGatewayVersion(t *testing.T) {
	s, err := getProtocGenGRPCGatewayGoSdkVersion()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}
