package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/jaronnie/protoc-gen-grpc-gateway-go/internal"
	"github.com/jaronnie/protoc-gen-grpc-gateway-go/internal/gen/codegenerator"
)

var (
	EnvFile            string
	GoVersion          string
	GoModule           string
	SdkDir             string
	ScopeVersion       string   // scopeVersion
	ScopeVersions      []string // scopeVersions used for clientSet
	GatewayPrefix      string   // microservice gateway prefix
	IsWarpHTTPResponse bool     // is warped code, data, message
	PluginOutputPath   string   // plugin output path
	SpecifiedMethods   []string // specified rpc methods
	Debug              bool     // used to debug
)

var (
	version string
	commit  string
)

type HTTPSdk struct{}

func main() {
	pflag.Usage = func() {
		pflag.PrintDefaults()
	}

	if len(os.Args) == 2 {
		if os.Args[1] == "version" {
			if commit != "" {
				commit = commit[:6]
			}
			fmt.Printf("%s-%s\n", version, commit)
			return
		}
	}

	bindFlag()

	hs := HTTPSdk{}
	protogen.Options{
		ParamFunc: pflag.CommandLine.Set,
	}.Run(hs.Generate)
}

func bindFlag() {
	pflag.StringVar(&EnvFile,
		"env_file",
		"",
		"set protoc go http sdk env file")
	pflag.StringVar(&GoVersion,
		"goVersion",
		"",
		"set go version")
	pflag.StringVar(&GoModule,
		"goModule",
		"",
		"set go module",
	)
	pflag.StringVar(&SdkDir,
		"sdkDir",
		"",
		"set sdk dir if in go module project",
	)
	pflag.StringVar(&ScopeVersion,
		"scopeVersion",
		"",
		"set scope version",
	)
	pflag.StringSliceVar(&ScopeVersions,
		"scopeVersions",
		nil,
		"set scope versions",
	)
	pflag.StringVar(&GatewayPrefix,
		"gatewayPrefix",
		"",
		"set gateway prefix",
	)
	pflag.BoolVar(&IsWarpHTTPResponse,
		"isWarpHTTPResponse",
		false,
		"isWarpHttpResponse",
	)
	pflag.StringVar(&PluginOutputPath,
		"pluginOutputPath",
		".",
		"set pluginOutputPath",
	)
	pflag.StringSliceVar(&SpecifiedMethods,
		"specifiedMethods",
		nil,
		"set SpecifiedMethods",
	)
	pflag.BoolVar(&Debug,
		"debug",
		false,
		"is debug",
	)

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
}

func (hs *HTTPSdk) Generate(plugin *protogen.Plugin) (err error) {
	if Debug {
		// into debug mode, you can attach to this process
		time.Sleep(time.Second * 10)
	}

	glog.V(1).Infof("get protoc go http sdk cmd flag logtostderr: [%v]", pflag.CommandLine.Lookup("logtostderr").Value)
	glog.V(1).Infof("get protoc go http sdk cmd flag env_file: [%v]", pflag.CommandLine.Lookup("env_file").Value)

	glog.V(1).Infof("get protoc go http sdk version: [%v-%v]", version, commit)

	codegenerator.SetSupportedFeaturesOnPluginGen(plugin)

	if EnvFile != "" {
		viper.SetConfigFile(EnvFile)

		err = viper.ReadInConfig()
		if err != nil {
			return err
		}
	}

	if err = viper.BindPFlags(pflag.CommandLine); err != nil {
		return err
	}

	err = internal.Generate(plugin)
	if err != nil {
		glog.Errorf("generate file: [%v]", err)
		return err
	}
	return nil
}
