package internal

import (
	"flag"

	"github.com/golang/glog"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type PluginEnv struct {
	GoVersion          string
	GoModule           string
	GatewayPrefix      string   // microservice gateway prefix
	ScopeVersion       string   // scopeVersion
	ScopeVersions      []string // scopeVersions used for clientSet
	IsWarpHttpResponse bool     // is warped code, data, message
}

func getPluginEnv() (*PluginEnv, error) {
	goVersion := viper.GetString("goVersion")
	if goVersion == "" {
		return nil, errors.New("empty go version, please set your go version")
	}
	goModule := viper.GetString("goModule")
	if goModule == "" {
		return nil, errors.New("empty go go module, please set your go module")
	}
	gatewayPrefix := viper.GetString("gatewayPrefix")

	scopeVersion := flag.Lookup("scopeVersion").Value.String()

	scopeVersions := viper.GetStringSlice("scopeVersions")

	if scopeVersion == "" {
		glog.Warningf("empty scope version, now generate empty go sdk")
		scopeVersions = nil
	}

	isWarpHttpResponse := viper.GetBool("isWarpHttpResponse")

	return &PluginEnv{GoVersion: goVersion,
		GoModule:           goModule,
		GatewayPrefix:      gatewayPrefix,
		ScopeVersion:       scopeVersion,
		ScopeVersions:      scopeVersions,
		IsWarpHttpResponse: isWarpHttpResponse,
	}, nil
}
