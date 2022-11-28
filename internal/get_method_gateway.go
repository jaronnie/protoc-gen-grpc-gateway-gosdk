package internal

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/golang/glog"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/jaronnie/protoc-gen-go-httpsdk/internal/vars"
)

func getMethodGateway(m *protogen.Method, env *PluginEnv) (*vars.Gateway, error) {

	if m.Desc.IsStreamingClient() && m.Desc.IsStreamingServer() {
		return nil, nil
	}

	// 获取 grpc gateway 中的 路由
	options, ok := m.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return nil, nil
	}
	httpRule, ok := proto.GetExtension(options, annotations.E_Http).(*annotations.HttpRule)
	if !ok {
		return nil, nil
	}

	var (
		httpMethod string
		url        string
	)

	switch httpRule.GetPattern().(type) {
	case *annotations.HttpRule_Get:
		httpMethod = "GET"
		url = httpRule.GetGet()
	case *annotations.HttpRule_Put:
		httpMethod = "PUT"
		url = httpRule.GetPut()
	case *annotations.HttpRule_Post:
		httpMethod = "POST"
		url = httpRule.GetPost()
	case *annotations.HttpRule_Delete:
		httpMethod = "DELETE"
		url = httpRule.GetDelete()
	case *annotations.HttpRule_Patch:
		httpMethod = "PATCH"
		url = httpRule.GetPatch()
	default:
		return nil, nil
	}

	pathParams, err := parsePathParam(url)
	if err != nil {
		return nil, nil
	}

	queryParams := createQueryParams(m)

	if env.GatewayPrefix != "" {
		// 获取 scope
		index := strings.LastIndex(env.ScopeVersion, "v")
		scope := env.ScopeVersion[:index]
		url = fmt.Sprintf("%s/%s%s", env.GatewayPrefix, scope, url)
	}

	// 遍历 protoRequestBody 的 fields
	httpRuleBodyName := httpRule.Body
	var httpRuleBodyGoName string

	if httpRuleBodyName != "*" {
		for _, v := range m.Input.Fields {
			if httpRuleBodyName == string(v.Desc.Name()) {
				httpRuleBodyGoName = v.GoName
			}
		}
	}

	glog.V(1).Infof("rpc method name: [%v], http request method: [%v], url: [%v], http request body [%s], http response body [%s], path params: [%v], query params: [%v], body: [%v]", m.GoName, httpMethod, url, m.Input.GoIdent.GoName, m.Output.GoIdent.GoName, pathParams, queryParams, httpRule.Body)
	return &vars.Gateway{
		ProtoRequestBody: vars.ProtoRequestBody{
			Name:         m.Input.GoIdent.GoName,
			GoImportPath: env.GoModule + "/pb" + strings.TrimLeft(string(m.Input.GoIdent.GoImportPath), "."),
			RootPath:     filepath.Base(env.GoModule + strings.TrimLeft(string(m.Input.GoIdent.GoImportPath), ".")),
		},
		HttpRequestBody: vars.HttpRequestBody{
			BodyName:   httpRuleBodyName,
			GoBodyName: httpRuleBodyGoName,
		},
		HttpResponseBody: vars.HttpResponseBody{
			Name:         m.Output.GoIdent.GoName,
			GoImportPath: env.GoModule + "/pb" + strings.TrimLeft(string(m.Output.GoIdent.GoImportPath), "."),
			RootPath:     filepath.Base(env.GoModule + strings.TrimLeft(string(m.Output.GoIdent.GoImportPath), ".")),
		},
		IsStreamClient:   m.Desc.IsStreamingClient(),
		IsStreamServer:   m.Desc.IsStreamingServer(),
		ProtoServiceName: string(m.Parent.Desc.Name()),
		FuncName:         m.GoName,
		HttpMethod:       httpMethod,
		Url:              url,
		PathParams:       pathParams,
		QueryParams:      queryParams,
	}, nil
}
