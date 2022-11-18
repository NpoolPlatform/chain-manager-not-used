package currencyvalue

import (
	currencyvalue "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/value"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	currencyvalue.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	currencyvalue.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
