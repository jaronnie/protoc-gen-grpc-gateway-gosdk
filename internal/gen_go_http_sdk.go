package internal

import (
	"fmt"
	"path/filepath"

	"github.com/golang/glog"
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/jaronnie/protoc-gen-go-httpsdk/internal/tpl"
	"github.com/jaronnie/protoc-gen-go-httpsdk/internal/tpl/fake"
	"github.com/jaronnie/protoc-gen-go-httpsdk/internal/tpl/rest"
	"github.com/jaronnie/protoc-gen-go-httpsdk/internal/tpl/typed"
	typedfake "github.com/jaronnie/protoc-gen-go-httpsdk/internal/tpl/typed/fake"
	"github.com/jaronnie/protoc-gen-go-httpsdk/internal/vars"
	"github.com/jaronnie/protoc-gen-go-httpsdk/utilx"
)

type GenHttpSdk struct {
	Plugin *protogen.Plugin
	Env    *PluginEnv
}

func Generate(plugin *protogen.Plugin) error {
	env, err := getPluginEnv()
	if err != nil {
		return err
	}
	marshalEnv, _ := utilx.BeautifyJson(env)
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
				gw, err := getMethodGateway(m, env)
				if err != nil || gw == nil {
					continue
				}
				serviceGws[vars.Resource(s.Desc.Name())] = append(serviceGws[vars.Resource(s.Desc.Name())], gw)
			}
			scopeResourceGws[vars.Scope(env.ScopeVersion)] = serviceGws
		}
	}

	marshalScopeResourceGws, err := utilx.BeautifyJson(scopeResourceGws)
	if err != nil {
		glog.Errorf("marshal scope resource gateways meet error. Err: [%v]", err)
	}
	glog.V(1).Infof("get scope service gateway: [%s]", marshalScopeResourceGws)

	if err = classifyResource(scopeResourceGws); err != nil {
		return err
	}

	marshalScopeResourceGws, _ = utilx.BeautifyJson(scopeResourceGws)
	glog.V(1).Infof("after classify scope service gateway get scope service gateway: [%s]", marshalScopeResourceGws)

	glog.V(1).Infof("generate client set successfully")

	ghs := GenHttpSdk{plugin, env}

	// generate sdk go mod file
	if err = ghs.GenGoMod(); err != nil {
		return err
	}

	// generate sdk rest frame
	if err = ghs.GenRestFrame(); err != nil {
		return err
	}

	// gen scope_client file
	if err = ghs.GenScopeClient(scopeResourceGws); err != nil {
		return err
	}

	if err = ghs.GenResource(scopeResourceGws); err != nil {
		return err
	}

	// gen clientset file
	if err = ghs.GenClientSet(scopeResourceGws); err != nil {
		return err
	}

	return nil
}

func (x *GenHttpSdk) GenGoMod() error {
	// init go mod
	var goModFile *protogen.GeneratedFile
	goModFile = x.Plugin.NewGeneratedFile("go.mod", "")
	template, err := utilx.ParseTemplate(tpl.GoModData{
		GoModule:  x.Env.GoModule,
		GoVersion: x.Env.GoVersion,
	}, []byte(tpl.GoModTpl))
	if err != nil {
		glog.Errorf("generate clientset meet error. Err: [%v]", err)
		return err
	}
	if _, err = goModFile.Write(template); err != nil {
		return err
	}
	return nil
}

func (x *GenHttpSdk) GenRestFrame() error {
	restFiles := []string{"client.go", "option.go", "request.go"}
	for _, v := range restFiles {
		rf := x.Plugin.NewGeneratedFile("rest/"+v, "")
		switch v {
		case "client.go":
			if _, err := rf.Write([]byte(rest.ClientTpl)); err != nil {
				return err
			}
		case "option.go":
			if _, err := rf.Write([]byte(rest.OptionTpl)); err != nil {
				return err
			}
		case "request.go":
			if _, err := rf.Write([]byte(rest.RequestTpl)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (x *GenHttpSdk) GenClientSet(scopeResourceGws vars.ScopeResourceGateway) error {
	scopeVersionsMap := make(map[string]string, 0)
	if len(x.Env.ScopeVersions) == 0 {
		scopeVersionsMap = getAllScopeVersionsMap(scopeResourceGws)
	} else {
		for _, v := range x.Env.ScopeVersions {
			scopeVersionsMap[v] = utilx.FirstUpper(v)
		}
	}

	// gen clientset
	var clientSetFile *protogen.GeneratedFile
	clientSetFile = x.Plugin.NewGeneratedFile("clientset.go", "")
	template, err := utilx.ParseTemplate(tpl.ClientSetData{
		GoModule:      x.Env.GoModule,
		RootModule:    filepath.Base(x.Env.GoModule),
		ScopeVersions: scopeVersionsMap,
	}, []byte(tpl.ClientSetTpl))
	if err != nil {
		glog.Errorf("generate clientset meet error. Err: [%v]", err)
		return err
	}
	if _, err := clientSetFile.Write(template); err != nil {
		return err
	}

	// gen fake clientset
	var fakeClientSetFile *protogen.GeneratedFile
	fakeClientSetFile = x.Plugin.NewGeneratedFile("fake/fake_clientset.go", "")
	template, err = utilx.ParseTemplate(fake.FakeClientSetData{
		GoModule:      x.Env.GoModule,
		ScopeVersions: scopeVersionsMap,
	}, []byte(fake.FakeClientSetTpl))
	if err != nil {
		glog.Errorf("generate clientset meet error. Err: [%v]", err)
		return err
	}
	if _, err := fakeClientSetFile.Write(template); err != nil {
		return err
	}
	return nil
}

func (x *GenHttpSdk) GenScopeClient(scopeResourceGws vars.ScopeResourceGateway) error {
	scopeVersionsMap := getAllScopeVersionsMap(scopeResourceGws)
	resources := getAllUpResources(scopeResourceGws)

	// gen direct client file
	var directClientFile *protogen.GeneratedFile
	directClientFile = x.Plugin.NewGeneratedFile("typed/direct_client.go", "")
	template, err := utilx.ParseTemplate(typed.DirectClientData{
		GoModule: x.Env.GoModule,
	}, []byte(typed.DirectClientTpl))
	if err != nil {
		glog.Errorf("generate direct_client meet error. Err: [%v]", err)
		return err
	}
	if _, err = directClientFile.Write(template); err != nil {
		return err
	}

	// gen fake direct client file
	var fakeDirectClientFile *protogen.GeneratedFile
	fakeDirectClientFile = x.Plugin.NewGeneratedFile("typed/fake/fake_direct_client.go", "")
	template, err = utilx.ParseTemplate(typed.DirectClientData{
		GoModule: x.Env.GoModule,
	}, []byte(typedfake.FakeDirectClientTpl))
	if err != nil {
		glog.Errorf("generate fake direct_client meet error. Err: [%v]", err)
		return err
	}
	if _, err = fakeDirectClientFile.Write(template); err != nil {
		return err
	}

	// gen scope client file
	for k, v := range scopeVersionsMap {
		var scopeClientFile *protogen.GeneratedFile
		scopeClientFile = x.Plugin.NewGeneratedFile(fmt.Sprintf("typed/%s/%s_client.go",
			k, k), "")
		template, err = utilx.ParseTemplate(typed.ScopeClientData{
			ScopeVersion:   k,
			UpScopeVersion: v,
			GoModule:       x.Env.GoModule,
			UpResources:    resources,
		}, []byte(typed.ScopeClientTpl))
		if err != nil {
			glog.Errorf("generate scope_client meet error. Err: [%v]", err)
			return err
		}
		if _, err = scopeClientFile.Write(template); err != nil {
			return err
		}

	}

	// gen fake scope client file
	for k, v := range scopeVersionsMap {
		var fakeScopeClientFile *protogen.GeneratedFile
		fakeScopeClientFile = x.Plugin.NewGeneratedFile(fmt.Sprintf("typed/%s/fake/fake_%s_client.go",
			k, k), "")
		template, err = utilx.ParseTemplate(typedfake.FakeScopeClientData{
			ScopeVersion:   k,
			UpScopeVersion: v,
			GoModule:       x.Env.GoModule,
			UpResources:    resources,
		}, []byte(typedfake.FakeScopeClientTpl))
		if err != nil {
			glog.Errorf("generate fake_scope_client meet error. Err: [%v]", err)
			return err
		}
		if _, err = fakeScopeClientFile.Write(template); err != nil {
			return err
		}

	}
	return nil
}

func (x *GenHttpSdk) GenResource(scopeResourceGws vars.ScopeResourceGateway) error {
	// gen scope resource
	for scope, resources := range scopeResourceGws {
		for resource, gws := range resources {
			var scopeResourceFile *protogen.GeneratedFile
			var goImportPaths []string
			for _, gw := range gws {
				goImportPaths = append(goImportPaths, gw.ProtoRequestBody.GoImportPath, gw.HttpResponseBody.GoImportPath)
			}
			// remove duplicate
			goImportPaths = utilx.RemoveDuplicateElement(goImportPaths)
			scopeResourceFile = x.Plugin.NewGeneratedFile(fmt.Sprintf("typed/%s/%s.go", scope, resource), "")
			template, err := utilx.ParseTemplate(typed.ResourceData{
				Gateways:           gws,
				IsWarpHttpResponse: x.Env.IsWarpHttpResponse,
				GoModule:           x.Env.GoModule,
				GoImportPaths:      goImportPaths,
				ScopeVersion:       string(scope),
				UpScopeVersion:     utilx.FirstUpper(string(scope)),
				Resource:           string(resource),
				UpResource:         utilx.FirstUpper(string(resource)),
			}, []byte(typed.ResourceTpl))
			if err != nil {
				glog.Errorf("generate resource meet error. Err: [%v]", err)
				return err
			}
			if _, err := scopeResourceFile.Write(template); err != nil {
				return err
			}

		}
	}

	// gen fake scope resource
	for scope, resources := range scopeResourceGws {
		for resource, gws := range resources {
			var fakeScopeResourceFile *protogen.GeneratedFile
			var goImportPaths []string
			for _, gw := range gws {
				goImportPaths = append(goImportPaths, gw.ProtoRequestBody.GoImportPath, gw.HttpResponseBody.GoImportPath)
			}
			// remove duplicate
			goImportPaths = utilx.RemoveDuplicateElement(goImportPaths)
			fakeScopeResourceFile = x.Plugin.NewGeneratedFile(fmt.Sprintf("typed/%s/fake/fake_%s.go", scope, resource), "")
			template, err := utilx.ParseTemplate(typedfake.FakeResourceData{
				Gateways:           gws,
				IsWarpHttpResponse: x.Env.IsWarpHttpResponse,
				GoModule:           x.Env.GoModule,
				GoImportPaths:      goImportPaths,
				ScopeVersion:       string(scope),
				UpScopeVersion:     utilx.FirstUpper(string(scope)),
				Resource:           string(resource),
				UpResource:         utilx.FirstUpper(string(resource)),
			}, []byte(typedfake.FakeResourceTpl))
			if err != nil {
				glog.Errorf("generate fake resource meet error. Err: [%v]", err)
				return err
			}
			if _, err := fakeScopeResourceFile.Write(template); err != nil {
				return err
			}

		}
	}
	return nil
}
