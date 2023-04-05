package gen

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/golang/glog"
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/jaronnie/protoc-gen-go-httpsdk/internal/env"
	"github.com/jaronnie/protoc-gen-go-httpsdk/internal/gateway"
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
	Env    *env.PluginEnv
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
		scopeVersionsMap = gateway.GetAllScopeVersionsMap(scopeResourceGws)
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
	scopeVersionsMap := gateway.GetAllScopeVersionsMap(scopeResourceGws)
	resources := gateway.GetAllUpResources(scopeResourceGws)

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

	// gen scope expansion resource
	for scope, resources := range scopeResourceGws {
		for resource, _ := range resources {
			filepath := path.Join("typed", string(scope), string(resource)+"_expansion.go")
			b, _ := os.ReadFile(path.Join(x.Env.PluginOutputPath, filepath))
			if string(b) != "" {
				continue
			}
			scopeResourceExpansionFile := x.Plugin.NewGeneratedFile(filepath, "")

			template, err := utilx.ParseTemplate(typed.ResourceExpansionData{
				ScopeVersion: string(scope),
				UpResource:   utilx.FirstUpper(string(resource)),
			}, []byte(typed.ResourceExpansionTpl))
			if err != nil {
				glog.Errorf("generate resource expansion meet error. Err: [%v]", err)
				return err
			}
			if _, err := scopeResourceExpansionFile.Write(template); err != nil {
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
