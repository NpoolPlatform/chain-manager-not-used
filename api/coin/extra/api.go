package coinextra

import (
	coinextra "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/extra"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	coinextra.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	coinextra.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
