package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jaronnie/protoc-gen-grpc-gateway-gosdk/grpc-restful/internal/svc"
	"github.com/jaronnie/protoc-gen-grpc-gateway-gosdk/grpc-restful/userpb"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *userpb.AddUserReq) (*userpb.AddUserResp, error) {
	// todo: add your logic here and delete this line

	return &userpb.AddUserResp{Id: 1}, nil
}
