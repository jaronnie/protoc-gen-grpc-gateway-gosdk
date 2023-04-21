package typed

import (
	"fmt"
	"testing"

	"github.com/jaronnie/protoc-gen-grpc-gateway-gosdk/internal/vars"
	"github.com/jaronnie/protoc-gen-grpc-gateway-gosdk/utilx"
)

func TestGenerateResourceFile(t *testing.T) {
	resourceData := ResourceData{
		Gateways: []*vars.Gateway{
			{
				ProtoRequestBody: vars.ProtoRequestBody{
					Name:         "HelloRequest",
					GoImportPath: "git.hyperchain.cn/bfsdk/pb/corev1",
					RootPath:     "corev1",
				},
				HTTPRequestBody: vars.HTTPRequestBody{
					BodyName:   "name",
					GoBodyName: "Name",
				},
				HTTPResponseBody: vars.HTTPResponseBody{
					Name:         "HelloReply",
					GoImportPath: "git.hyperchain.cn/bfsdk/pb/corev1",
					RootPath:     "corev1",
				},
				IsSpecified: true,
				// IsStreamClient:   true,
				ProtoServiceName: "Core",
				FuncName:         "SayHello",
				HTTPMethod:       "post",
				URL:              "/api/v1/credential/{id}",
				PathParams: []*vars.PathParam{
					{
						Name:   "name",
						GoName: "Name",
					},
				},
				QueryParams: nil,
			},
		},
		IsWarpHTTPResponse: true,
		GoImportPaths:      []string{"github.com/jaronnie/autosdk/pb/corev1"},
		ScopeVersion:       "corev1",
		UpScopeVersion:     "Corev1",
		Resource:           "credential",
		UpResource:         "Credential",
	}
	template, err := utilx.ParseTemplate(resourceData, []byte(ResourceTpl))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(template))
}
