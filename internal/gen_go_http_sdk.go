package internal

import (
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/jaronnie/protoc-gen-grpc-gateway-go/internal/env"
	"github.com/jaronnie/protoc-gen-grpc-gateway-go/internal/gateway"
	"github.com/jaronnie/protoc-gen-grpc-gateway-go/internal/gen"
	"github.com/jaronnie/protoc-gen-grpc-gateway-go/internal/vars"
	"github.com/jaronnie/protoc-gen-grpc-gateway-go/utilx"
)

func Generate(plugin *protogen.Plugin) error {
	pluginEnv, err := env.GetPluginEnv()
	if err != nil {
		return err
	}
	marshalEnv, _ := utilx.BeautifyJSON(pluginEnv)
	glog.V(1).Infof("get plugin env [%v]", marshalEnv)

	// get scope service gateway
	scopeResourceGws := make(vars.ScopeResourceGateway, 0)
	serviceGws := make(vars.ServiceGateway, 0)

	for _, f := range plugin.Files {
		if len(f.Services) == 0 {
			// just parse proto service
			continue
		}
		glog.V(1).Infof("generated file prefix: %v", f.GeneratedFilenamePrefix)
		glog.V(1).Infof("go import path: %v", f.GoImportPath)

		for _, s := range f.Services {
			// for each service methods
			for _, m := range s.Methods {
				gw, err := gateway.GetMethodGateway(m, pluginEnv)
				if err != nil || gw == nil {
					continue
				}
				serviceGws[vars.Resource(s.Desc.Name())] = append(serviceGws[vars.Resource(s.Desc.Name())], gw)
			}
			scopeResourceGws[vars.Scope(pluginEnv.ScopeVersion)] = serviceGws
		}
	}

	marshalScopeResourceGws, err := utilx.BeautifyJSON(scopeResourceGws)
	if err != nil {
		glog.Errorf("marshal scope resource gateways meet error. Err: [%v]", err)
	}
	glog.V(1).Infof("get scope service gateway: [%s]", marshalScopeResourceGws)

	if err = gateway.ClassifyResource(scopeResourceGws); err != nil {
		return err
	}

	marshalScopeResourceGws, _ = utilx.BeautifyJSON(scopeResourceGws)
	glog.V(1).Infof("after classify scope service gateway get scope service gateway: [%s]", marshalScopeResourceGws)

	ghs := gen.HTTPSdk{Plugin: plugin, Env: pluginEnv}

	if pluginEnv.IsNeedGenerateGoMod() {
		// generate sdk go mod file
		if err = ghs.GenGoMod(); err != nil {
			return err
		}
		glog.V(1).Infof("generate go mod successfully")
	}

	// generate sdk rest frame
	if err = ghs.GenRestFrame(); err != nil {
		return errors.Wrap(err, "generate sdk rest frame")
	}
	glog.V(1).Infof("generate sdk rest frame successfully")

	// gen scope_client file
	if err = ghs.GenScopeClient(scopeResourceGws); err != nil {
		return errors.Wrap(err, "generate scope client")
	}
	glog.V(1).Infof("generate scope client successfully")

	if err = ghs.GenResource(scopeResourceGws); err != nil {
		return errors.Wrap(err, "generate resource")
	}
	glog.V(1).Infof("generate scope resource successfully")

	// gen clientset file
	if err = ghs.GenClientSet(scopeResourceGws); err != nil {
		return errors.Wrap(err, "generate client set")
	}
	glog.V(1).Infof("generate client set successfully")

	return nil
}
