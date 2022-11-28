package internal

import (
	"regexp"
	"sort"
	"strings"

	"github.com/pkg/errors"

	"github.com/jaronnie/protoc-gen-go-httpsdk/internal/vars"
	"github.com/jaronnie/protoc-gen-go-httpsdk/utilx"
)

func classifyResource(scopeResourceGws vars.ScopeResourceGateway) error {
	// Rules:
	//	如果 scope 中只有一个 service, 那么就通过路由对资源进行分类, 路径匹配格式为 /api/v1/credential. 示例的资源对象即为 credential
	// 	如果已经有多个 service, 那么不进行更改, 以 service 的名称作为资源对象, 并将首字母转为小写
	for scope, services := range scopeResourceGws {
		if len(services) == 1 {
			var serviceName string
			// 根据路由进行分类
			// NOTICE: 仅支持分类 /api/{version}/{resource} 的路由
			for _, gws := range services {
				serviceGws := make(vars.ServiceGateway, 0)
				for _, gw := range gws {
					serviceName = gw.ProtoServiceName
					resource, err := getResourceByUri(gw.Url)
					if err != nil {
						return err
					}
					serviceGws[resource] = append(serviceGws[resource], gw)
				}
				scopeResourceGws[scope] = serviceGws
			}
			// delete map
			delete(services, vars.Resource(serviceName))

		} else {
			// 超过 1 的不用分类
			// 但是保证 service 的首字母是小写的
			for service, gws := range services {
				delete(scopeResourceGws[scope], service)
				scopeResourceGws[scope][vars.Resource(utilx.FirstLower(string(service)))] = gws
			}
			continue
		}
	}

	return nil
}

func getResourceByUri(uri string) (vars.Resource, error) {
	// regexp.MustCompile(`\d{1,3}\.\d{1,3}\.?\d{0,5}`)
	resourceReg := regexp.MustCompile(`/api/v\d{1,3}(.\d+){0,2}/.*`)

	findString := resourceReg.FindString(uri)
	split := strings.Split(findString, "/")
	if len(split) < 4 {
		return "", errors.Errorf("can not find resource in uri [%s]", uri)
	}
	return vars.Resource(split[3]), nil
}

func getAllUpResources(scopeResourceGws vars.ScopeResourceGateway) []string {
	var resources []string
	for _, resourceMap := range scopeResourceGws {
		for resource, _ := range resourceMap {
			resources = append(resources, utilx.FirstUpper(string(resource)))
		}
	}
	sort.Strings(resources)
	return resources
}
