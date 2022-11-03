package fee

import (
	fee "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fee"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	fee.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	fee.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
