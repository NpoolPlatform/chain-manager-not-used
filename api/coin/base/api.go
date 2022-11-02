package coinbase

import (
	coinbase "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/base"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	coinbase.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	coinbase.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
