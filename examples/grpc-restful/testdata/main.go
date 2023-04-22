package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaronnie/protoc-gen-grpc-gateway-gosdk/grpc-restful/pkgsdk"
	"github.com/jaronnie/protoc-gen-grpc-gateway-gosdk/grpc-restful/pkgsdk/pb/userpb"
	"github.com/jaronnie/protoc-gen-grpc-gateway-gosdk/grpc-restful/pkgsdk/rest"
)

func main() {
	cs, err := pkgsdk.NewClientWithOptions(
		rest.WithProtocol("http"),
		rest.WithAddr("127.0.0.1"),
		rest.WithPort("8081"),
		rest.WithHeaders(http.Header{"Content-Type": []string{"application/json"}}),
	)

	if err != nil {
		panic(err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	//defer cancel()

	data, err := cs.Userv1().User().Add(context.Background(), &userpb.AddUserReq{
		Name: "jaronnie",
		Age:  22,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
