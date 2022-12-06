package internal

import (
	"github.com/spf13/viper"

	"github.com/jaronnie/protoc-gen-go-httpsdk/utilx"
)

type PluginEnv struct {
	GoVersion                         string   `validate:"required"`
	GoModule                          string   `validate:"required"`
	ScopeVersion                      string   `validate:"required"` // scopeVersion
	ScopeVersions                     []string // scopeVersions used for clientSet
	GatewayPrefix                     string   // microservice gateway prefix
	IsWarpHttpResponse                bool     // is warped code, data, message
	IsResourceExpansionCreateOrUpdate bool     // is to create or update resource expansion
}

func getPluginEnv() (*PluginEnv, error) {
	var c PluginEnv

	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	if err := utilx.ValidateStruct(c); err != nil {
		return nil, err
	}

	return &c, nil
}
