package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/jaronnie/protoc-gen-go-httpsdk/internal"
)

var (
	EnvFile      string
	ScopeVersion string
)

type HttpSdk struct{}

func main() {
	flag.StringVar(&EnvFile,
		"env_file",
		"./conf/cfg.toml",
		"set protoc go http sdk env file")
	flag.StringVar(&ScopeVersion,
		"scopeVersion",
		"",
		"set scope version")
	flag.Parse()

	hs := HttpSdk{}
	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(hs.Generate)
}

func (hs *HttpSdk) Generate(plugin *protogen.Plugin) error {
	glog.V(1).Infof("get protoc go http sdk cmd flag logtostderr: [%v]", flag.CommandLine.Lookup("logtostderr").Value)
	glog.V(1).Infof("get protoc go http sdk cmd flag env_file: [%v]", flag.CommandLine.Lookup("env_file").Value)

	viper.SetConfigFile(EnvFile)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
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
