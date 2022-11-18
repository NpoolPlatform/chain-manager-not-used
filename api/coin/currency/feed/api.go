package currencyfeed

import (
	currencyfeed "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/feed"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	currencyfeed.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	currencyfeed.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
