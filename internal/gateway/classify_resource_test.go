package gateway

import (
	"testing"

	"github.com/bmizerany/assert"

	"github.com/jaronnie/protoc-gen-grpc-gateway-gosdk/internal/vars"
)

func TestGetResourceByUri(t *testing.T) {
	resource, err := getResourceByURI("/api/v1.9.8/credential")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	assert.Equal(t, vars.Resource("credential"), resource)

	resource, err = getResourceByURI("/api/v1.0/credential")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	assert.Equal(t, vars.Resource("credential"), resource)

	resource, err = getResourceByURI("/api/v1/credential")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	assert.Equal(t, vars.Resource("credential"), resource)
}
