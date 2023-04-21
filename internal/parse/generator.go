package parse

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/jaronnie/protoc-gen-grpc-gateway-go/internal/vars"
)

var toCamelCaseRe = regexp.MustCompile(`(^[A-Za-z])|(_|\.)([A-Za-z])`)

func toCamelCase(str string) string {
	return toCamelCaseRe.ReplaceAllStringFunc(str, func(s string) string {
		return strings.ToUpper(strings.Replace(s, "_", "", -1))
	})
}

func PathParam(pattern string) ([]*vars.PathParam, error) {
	if !strings.HasPrefix(pattern, "/") {
		return nil, fmt.Errorf("no leading /")
	}
	tokens, _ := tokenize(pattern[1:])

	p := parser{tokens: tokens}
	segs, err := p.topLevelSegments()
	if err != nil {
		return nil, err
	}

	params := make([]*vars.PathParam, 0)
	for i, seg := range segs {
		if v, ok := seg.(variable); ok {
			params = append(params, &vars.PathParam{
				Index:  i + 1,
				Name:   v.path,
				GoName: toCamelCase(v.path),
			})
		}
	}

	sort.Slice(params, func(i, j int) bool {
		a := params[i]
		b := params[j]
		if len(strings.Split(a.Name, ".")) < len(strings.Split(b.Name, ".")) {
			return true
		}
		return params[i].Name < params[j].Name
	})

	return params, nil
}

func CreateQueryParams(method *protogen.Method) []*vars.QueryParam {
	queryParams := make([]*vars.QueryParam, 0)

	var f func(parent *vars.QueryParam, fields []*protogen.Field)

	f = func(parent *vars.QueryParam, fields []*protogen.Field) {
		for _, field := range fields {
			if field.Desc.Kind() == protoreflect.MessageKind {
				q := &vars.QueryParam{
					// Field:  field,
					GoName: fmt.Sprintf("%s.", field.GoName),
					Name:   fmt.Sprintf("%s.", field.Desc.Name()),
				}
				f(q, field.Message.Fields)
				continue
			}
			queryParams = append(queryParams, &vars.QueryParam{
				// Field:  field,
				GoName: fmt.Sprintf("%s%s", parent.GoName, field.GoName),
				Name:   fmt.Sprintf("%s%s", parent.Name, field.Desc.Name()),
			})
		}
	}

	f(&vars.QueryParam{GoName: "", Name: ""}, method.Input.Fields)

	return queryParams
}
