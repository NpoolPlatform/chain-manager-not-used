package exrate

import (
	exrate "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/exrate"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	exrate.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	exrate.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
