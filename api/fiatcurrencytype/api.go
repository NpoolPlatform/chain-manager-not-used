package fiatcurrencytype

import (
	"github.com/NpoolPlatform/message/npool/chain/mgr/v1/fiatcurrencytype"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	fiatcurrencytype.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	fiatcurrencytype.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
