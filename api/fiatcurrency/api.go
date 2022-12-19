package fiatcurrency

import (
	"github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fiatcurrency"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	fiatcurrency.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	fiatcurrency.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
