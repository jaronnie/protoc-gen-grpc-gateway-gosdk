package gateway

import (
	"regexp"
	"sort"
	"strings"

	"github.com/pkg/errors"

	"github.com/jaronnie/protoc-gen-go-httpsdk/internal/vars"
	"github.com/jaronnie/protoc-gen-go-httpsdk/utilx"
)

func ClassifyResource(scopeResourceGws vars.ScopeResourceGateway) error {
	for scope, services := range scopeResourceGws {
		if len(services) == 1 {
			var serviceName string
			// classified by url
			// NOTICE: regexp by /api/{version}/{resource}
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
			// if services count more than or equal 1
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
	resourceReg := regexp.MustCompile(`/api/v\d{1,3}(.\d+){0,2}/.*`)

	findString := resourceReg.FindString(uri)
	split := strings.Split(findString, "/")
	if len(split) < 4 {
		return "", errors.Errorf("can not find resource in uri [%s]", uri)
	}
	return vars.Resource(split[3]), nil
}

func GetAllUpResources(scopeResourceGws vars.ScopeResourceGateway) []string {
	var resources []string
	for _, resourceMap := range scopeResourceGws {
		for resource, _ := range resourceMap {
			resources = append(resources, utilx.FirstUpper(string(resource)))
		}
	}
	sort.Strings(resources)
	return resources
}
