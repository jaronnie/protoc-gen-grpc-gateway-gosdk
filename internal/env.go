package internal

import (
	"github.com/jaronnie/protoc-gen-go-httpsdk/utilx"
	"github.com/spf13/viper"
)

type PluginEnv struct {
	GoVersion          string   `validate:"required"`
	GoModule           string   `validate:"required"`
	ScopeVersion       string   `validate:"required"` // scopeVersion
	ScopeVersions      []string `validate:"required"` // scopeVersions used for clientSet
	GatewayPrefix      string   // microservice gateway prefix
	IsWarpHttpResponse bool     // is warped code, data, message
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
