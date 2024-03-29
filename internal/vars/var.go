package vars

type Scope string

type Resource string

type ScopeResourceGateway map[Scope]map[Resource][]*Gateway

type ServiceGateway map[Resource][]*Gateway

type Gateway struct {
	ProtoRequestBody ProtoRequestBody
	HTTPRequestBody  HTTPRequestBody
	HTTPResponseBody HTTPResponseBody

	IsStreamClient   bool // is stream client
	IsStreamServer   bool // is stream server
	ProtoServiceName string
	FuncName         string
	Comments         string
	HTTPMethod       string
	URL              string
	PathParams       []*PathParam
	QueryParams      []*QueryParam

	IsSpecified bool
}

type ProtoRequestBody struct {
	Name         string
	GoImportPath string // github.com/autosdk/pb/corev1
	RootPath     string // corev1
}

type HTTPRequestBody struct {
	BodyName   string // * or others
	GoBodyName string
}

type HTTPResponseBody struct {
	Name         string
	GoImportPath string
	RootPath     string
}

type PathParam struct {
	Index  int
	Name   string
	GoName string
}

type QueryParam struct {
	GoName string
	Name   string
}
