package gen

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/golang/glog"
	"github.com/rinchsan/gosimports"
	"github.com/samber/lo"
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/jaronnie/protoc-gen-grpc-gateway-go/internal/env"
	"github.com/jaronnie/protoc-gen-grpc-gateway-go/internal/gateway"
	"github.com/jaronnie/protoc-gen-grpc-gateway-go/internal/tpl"
	"github.com/jaronnie/protoc-gen-grpc-gateway-go/internal/tpl/fake"
	"github.com/jaronnie/protoc-gen-grpc-gateway-go/internal/tpl/rest"
	"github.com/jaronnie/protoc-gen-grpc-gateway-go/internal/tpl/typed"
	typedfake "github.com/jaronnie/protoc-gen-grpc-gateway-go/internal/tpl/typed/fake"
	"github.com/jaronnie/protoc-gen-grpc-gateway-go/internal/vars"
	"github.com/jaronnie/protoc-gen-grpc-gateway-go/utilx"
)

type GenHttpSdk struct {
	Plugin *protogen.Plugin
	Env    *env.PluginEnv
}

func (x *GenHttpSdk) GenGoMod() error {
	// init go mod
	goModFile := x.Plugin.NewGeneratedFile("go.mod", "")
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
		rf := x.Plugin.NewGeneratedFile(path.Join("rest", v), "")
		switch v {
		case "client.go":
			template, err := utilx.ParseTemplate(nil, []byte(rest.ClientTpl))
			if err != nil {
				glog.Errorf("generate rest frame client meet error. Err: [%v]", err)
				return err
			}
			// format template
			templateFormat, err := gosimports.Process("", template, nil)
			if err != nil {
				return err
			}
			if _, err := rf.Write(templateFormat); err != nil {
				return err
			}
		case "option.go":
			template, err := utilx.ParseTemplate(nil, []byte(rest.OptionTpl))
			if err != nil {
				glog.Errorf("generate rest frame option meet error. Err: [%v]", err)
				return err
			}
			// format template
			templateFormat, err := gosimports.Process("", template, nil)
			if err != nil {
				return err
			}
			if _, err := rf.Write(templateFormat); err != nil {
				return err
			}
		case "request.go":
			template, err := utilx.ParseTemplate(nil, []byte(rest.RequestTpl))
			if err != nil {
				glog.Errorf("generate rest frame request meet error. Err: [%v]", err)
				return err
			}
			// format template
			templateFormat, err := gosimports.Process("", template, nil)
			if err != nil {
				return err
			}
			if _, err := rf.Write(templateFormat); err != nil {
				return err
			}
		}
	}
	return nil
}

func (x *GenHttpSdk) GenClientSet(scopeResourceGws vars.ScopeResourceGateway) error {
	scopeVersionsMap := make(map[string]string, 0)
	if len(x.Env.ScopeVersions) == 0 {
		scopeVersionsMap = gateway.GetAllScopeVersionsMap(scopeResourceGws)
	} else {
		for _, v := range x.Env.ScopeVersions {
			scopeVersionsMap[v] = utilx.FirstUpper(v)
		}
	}

	// 适配仓如 core-go
	rootModule := filepath.Base(x.Env.GoModule)
	if strings.Contains(rootModule, "-") {
		s := strings.Split(rootModule, "-")
		rootModule = s[0] + "sdk"
	}

	// gen clientset
	clientSetFile := x.Plugin.NewGeneratedFile("clientset.go", "")
	template, err := utilx.ParseTemplate(tpl.ClientSetData{
		GoModule:      x.Env.GoModule,
		RootModule:    rootModule,
		ScopeVersions: scopeVersionsMap,
	}, []byte(tpl.ClientSetTpl))
	if err != nil {
		glog.Errorf("generate clientset meet error. Err: [%v]", err)
		return err
	}

	// format template
	templateFormat, err := gosimports.Process("", template, nil)
	if err != nil {
		return err
	}

	if _, err := clientSetFile.Write(templateFormat); err != nil {
		return err
	}

	// gen fake clientset
	fakeClientSetFile := x.Plugin.NewGeneratedFile(path.Join("fake", "fake_clientset.go"), "")
	template, err = utilx.ParseTemplate(fake.FakeClientSetData{
		GoModule:      x.Env.GoModule,
		ScopeVersions: scopeVersionsMap,
	}, []byte(fake.FakeClientSetTpl))
	if err != nil {
		glog.Errorf("generate clientset meet error. Err: [%v]", err)
		return err
	}

	// format template
	templateFormat, err = gosimports.Process("", template, nil)
	if err != nil {
		return err
	}

	if _, err := fakeClientSetFile.Write(templateFormat); err != nil {
		return err
	}
	return nil
}

func (x *GenHttpSdk) GenScopeClient(scopeResourceGws vars.ScopeResourceGateway) error {
	scopeVersionsMap := gateway.GetAllScopeVersionsMap(scopeResourceGws)
	resources := gateway.GetAllUpResources(scopeResourceGws)

	// gen direct client file
	directClientFile := x.Plugin.NewGeneratedFile(path.Join("typed", "direct_client.go"), "")
	template, err := utilx.ParseTemplate(typed.DirectClientData{
		GoModule: x.Env.GoModule,
	}, []byte(typed.DirectClientTpl))
	if err != nil {
		glog.Errorf("generate direct_client meet error. Err: [%v]", err)
		return err
	}

	// format template
	templateFormat, err := gosimports.Process("", template, nil)
	if err != nil {
		return err
	}

	if _, err = directClientFile.Write(templateFormat); err != nil {
		return err
	}

	// gen fake direct client file
	fakeDirectClientFile := x.Plugin.NewGeneratedFile(path.Join("typed", "fake", "fake_direct_client.go"), "")
	template, err = utilx.ParseTemplate(typed.DirectClientData{
		GoModule: x.Env.GoModule,
	}, []byte(typedfake.FakeDirectClientTpl))
	if err != nil {
		glog.Errorf("generate fake direct_client meet error. Err: [%v]", err)
		return err
	}

	// format template
	templateFormat, err = gosimports.Process("", template, nil)
	if err != nil {
		return err
	}

	if _, err = fakeDirectClientFile.Write(templateFormat); err != nil {
		return err
	}

	// gen scope client file
	for k, v := range scopeVersionsMap {
		scopeClientFile := x.Plugin.NewGeneratedFile(path.Join("typed", k, k+"_client.go"), "")
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

		// format template
		templateFormat, err := gosimports.Process("", template, nil)
		if err != nil {
			return err
		}

		if _, err = scopeClientFile.Write(templateFormat); err != nil {
			return err
		}

	}

	// gen fake scope client file
	for k, v := range scopeVersionsMap {
		fakeScopeClientFile := x.Plugin.NewGeneratedFile(path.Join("typed", k, "fake", "fake_"+k+"_client.go"), "")
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

		// format template
		templateFormat, err := gosimports.Process("", template, nil)
		if err != nil {
			return err
		}

		if _, err = fakeScopeClientFile.Write(templateFormat); err != nil {
			return err
		}

	}
	return nil
}

func (x *GenHttpSdk) GenResource(scopeResourceGws vars.ScopeResourceGateway) error {
	if err := genScopeResource(x.Plugin, x.Env, scopeResourceGws); err != nil {
		return err
	}

	if err := genScopeResourceExpansion(x.Plugin, x.Env, scopeResourceGws); err != nil {
		return err
	}

	if err := genScopeFakeResource(x.Plugin, x.Env, scopeResourceGws); err != nil {
		return err
	}

	if err := genScopeFakeResourceExpansion(x.Plugin, x.Env, scopeResourceGws); err != nil {
		return err
	}

	return nil
}

func genScopeResource(plugin *protogen.Plugin, env *env.PluginEnv, scopeResourceGws vars.ScopeResourceGateway) error {
	// gen scope resource
	for scope, resources := range scopeResourceGws {
		for resource, gws := range resources {
			var scopeResourceFile *protogen.GeneratedFile
			var goImportPaths []string

			for _, gw := range gws {
				if !gw.IsStreamClient && !gw.IsStreamServer {
					// import context
					if !lo.Contains(goImportPaths, "context") {
						goImportPaths = append(goImportPaths, "context")
					}
				}
				if !lo.Contains(goImportPaths, gw.ProtoRequestBody.GoImportPath) {
					goImportPaths = append(goImportPaths, gw.ProtoRequestBody.GoImportPath)
				}
				if !lo.Contains(goImportPaths, gw.HttpResponseBody.GoImportPath) {
					goImportPaths = append(goImportPaths, gw.HttpResponseBody.GoImportPath)
				}
			}

			// import rest frame
			goImportPaths = append(goImportPaths, fmt.Sprintf("%s/rest", env.GoModule))

			scopeResourceFile = plugin.NewGeneratedFile(path.Join("typed", string(scope), string(resource)+".go"), "")
			template, err := utilx.ParseTemplate(typed.ResourceData{
				Gateways:           gws,
				IsWarpHttpResponse: env.IsWarpHttpResponse,
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

			// format template
			templateFormat, err := gosimports.Process("", template, nil)
			if err != nil {
				return err
			}

			if _, err := scopeResourceFile.Write(templateFormat); err != nil {
				return err
			}

		}
	}
	return nil
}

func genScopeResourceExpansion(plugin *protogen.Plugin, env *env.PluginEnv, scopeResourceGws vars.ScopeResourceGateway) error {
	// gen scope expansion resource
	for scope, resources := range scopeResourceGws {
		for resource := range resources {
			file := path.Join("typed", string(scope), string(resource)+"_expansion.go")
			b, _ := os.ReadFile(path.Join(env.PluginOutputPath, file))
			if string(b) != "" {
				continue
			}
			scopeResourceExpansionFile := plugin.NewGeneratedFile(file, "")

			template, err := utilx.ParseTemplate(typed.ResourceExpansionData{
				ScopeVersion: string(scope),
				UpResource:   utilx.FirstUpper(string(resource)),
			}, []byte(typed.ResourceExpansionTpl))
			if err != nil {
				glog.Errorf("generate resource expansion meet error. Err: [%v]", err)
				return err
			}

			// format template
			templateFormat, err := gosimports.Process("", template, nil)
			if err != nil {
				return err
			}

			if _, err := scopeResourceExpansionFile.Write(templateFormat); err != nil {
				return err
			}
		}
	}
	return nil
}

func genScopeFakeResource(plugin *protogen.Plugin, env *env.PluginEnv, scopeResourceGws vars.ScopeResourceGateway) error {
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

			fakeScopeResourceFile = plugin.NewGeneratedFile(path.Join("typed", string(scope), "fake", "fake_"+string(resource)+".go"), "")
			template, err := utilx.ParseTemplate(typedfake.FakeResourceData{
				Gateways:           gws,
				IsWarpHttpResponse: env.IsWarpHttpResponse,
				GoModule:           env.GoModule,
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

			// format template
			templateFormat, err := gosimports.Process("", template, nil)
			if err != nil {
				return err
			}

			if _, err := fakeScopeResourceFile.Write(templateFormat); err != nil {
				return err
			}

		}
	}
	return nil
}

func genScopeFakeResourceExpansion(plugin *protogen.Plugin, env *env.PluginEnv, scopeResourceGws vars.ScopeResourceGateway) error {
	// gen fake scope resource expansion
	for scope, resources := range scopeResourceGws {
		for resource := range resources {
			file := path.Join("typed", string(scope), "fake", "fake_"+string(resource)+"_expansion.go")
			b, _ := os.ReadFile(path.Join(env.PluginOutputPath, file))
			if string(b) != "" {
				continue
			}

			fakeScopeResourceExpansionFile := plugin.NewGeneratedFile(file, "")
			template, err := utilx.ParseTemplate(typedfake.FakeResourceExpansionData{
				UpResource: utilx.FirstUpper(string(resource)),
			}, []byte(typedfake.FakeResourceExpansionTpl))
			if err != nil {
				glog.Errorf("generate fake resource expansion meet error. Err: [%v]", err)
				return err
			}

			// format template
			templateFormat, err := gosimports.Process("", template, nil)
			if err != nil {
				return err
			}

			if _, err := fakeScopeResourceExpansionFile.Write(templateFormat); err != nil {
				return err
			}

		}
	}
	return nil
}
