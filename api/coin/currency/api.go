package currency

import (
	currency "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	currency.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	currency.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
