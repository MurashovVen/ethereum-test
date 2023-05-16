package grpc

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"

	"ethereum-test/internal/domain/service"
	ethgrpc "ethereum-test/pkg/grpc"
)

func New(svc *service.Service) *grpc.Server {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(recovery.UnaryServerInterceptor()),
		grpc.StreamInterceptor(recovery.StreamServerInterceptor()),
	)

	ethgrpc.RegisterEthereumServer(server, &ethereumService{svc: svc})

	return server
}
