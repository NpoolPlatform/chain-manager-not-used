package currencytype

import (
	"github.com/NpoolPlatform/message/npool/chain/mgr/v1/fiat/currencytype"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	currencytype.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	currencytype.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
