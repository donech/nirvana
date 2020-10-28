package grpc

import (
	"context"

	"github.com/donech/nirvana/internal/config"
	"github.com/donech/nirvana/internal/entry/grpc/proto"
	"github.com/donech/nirvana/internal/entry/grpc/service"
	"github.com/donech/tool/entry/grpc"
	"github.com/donech/tool/xlog"
	"go.uber.org/zap"
	grpc2 "google.golang.org/grpc"
)

func NewRegisteServer() grpc.RegisteServer {
	return func(server *grpc2.Server) {
		srv := service.HelloService{}
		proto.RegisterGreeterServer(server, srv)
	}
}

func NewEntry(config *config.Config, logger *zap.Logger, server grpc.RegisteServer) *grpc.Entry {
	xlog.S(context.Background()).Infof("config %#v", config)
	return grpc.New(config.Grpc, logger, server)
}
