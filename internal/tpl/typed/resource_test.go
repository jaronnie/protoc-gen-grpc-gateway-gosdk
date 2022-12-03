package typed

import (
	"fmt"
	"testing"

	"github.com/jaronnie/protoc-gen-go-httpsdk/internal/vars"
	"github.com/jaronnie/protoc-gen-go-httpsdk/utilx"
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
				HttpRequestBody: vars.HttpRequestBody{
					BodyName:   "name",
					GoBodyName: "Name",
				},
				HttpResponseBody: vars.HttpResponseBody{
					Name:         "HelloReply",
					GoImportPath: "git.hyperchain.cn/bfsdk/pb/corev1",
					RootPath:     "corev1",
				},
				IsStreamClient:   true,
				ProtoServiceName: "Core",
				FuncName:         "SayHello",
				HttpMethod:       "post",
				Url:              "/api/v1/credential/{id}",
				PathParams: []*vars.PathParam{
					{
						Name:   "name",
						GoName: "Name",
					},
				},
				QueryParams: nil,
			},
		},
		IsWarpHttpResponse: true,
		GoModule:           "github.com/jaronnie/autosdk",
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
