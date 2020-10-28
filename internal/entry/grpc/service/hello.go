package service

import (
	"context"

	"github.com/donech/nirvana/internal/entry/grpc/proto"
	"github.com/donech/tool/xlog"
)

type HelloService struct {
}

func (h HelloService) SayHello(ctx context.Context, req *proto.HelloReq) (*proto.HelloResp, error) {
	xlog.S(ctx).Infof("recieve request %#v", req)
	return &proto.HelloResp{
		Message: "hello " + req.Name,
	}, nil
}
