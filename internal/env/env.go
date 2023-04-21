package env

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/jaronnie/protoc-gen-grpc-gateway-gosdk/utilx"
	"github.com/jaronnie/protoc-gen-grpc-gateway-gosdk/utilx/execx"
)

type PluginEnv struct {
	GoVersion          string
	GoModule           string `validate:"required"`
	SdkDir             string
	ScopeVersion       string   `validate:"required"` // scopeVersion
	ScopeVersions      []string // scopeVersions used for clientSet
	GatewayPrefix      string   // microservice gateway prefix
	IsWarpHTTPResponse bool     // is warped code, data, message
	// IsResourceExpansionUpdate bool     // is to update resource expansion
	PluginOutputPath string // plugin output path

	SpecifiedMethods []string // specified rpc methods
}

func (p *PluginEnv) IsNeedGenerateGoMod() bool {
	if p.GoVersion != "" && p.GoModule != "" {
		return true
	}
	return false
}

func GetPluginEnv() (*PluginEnv, error) {
	var c PluginEnv

	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	if pflag.CommandLine.Lookup("scopeVersion").Value.String() != "" {
		c.ScopeVersion = pflag.CommandLine.Lookup("scopeVersion").Value.String()
	}

	if c.GoVersion == "" && c.GoModule == "" {
		wd, _ := os.Getwd()
		mod, err := GetGoMod(wd)
		if err != nil {
			return nil, errors.Errorf("not in go module")
		}
		c.GoModule = mod.Path + "/" + c.SdkDir
	}

	if err := utilx.ValidateStruct(c); err != nil {
		return nil, err
	}

	return &c, nil
}

// Module contains the relative data of go module,
// which is the result of the command go list
type Module struct {
	Path      string
	Main      bool
	Dir       string
	GoMod     string
	GoVersion string
}

// GetGoMod is used to determine whether workDir is a go module project through command `go list -json -m`
func GetGoMod(workDir string) (*Module, error) {
	if len(workDir) == 0 {
		return nil, errors.New("the work directory is not found")
	}
	if _, err := os.Stat(workDir); err != nil {
		return nil, err
	}

	data, err := execx.Run("go list -json -m", workDir)
	if err != nil {
		return nil, nil
	}

	var m Module
	err = json.Unmarshal([]byte(data), &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}
