package appcoin

import (
	appcoin "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/appcoin"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	appcoin.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	appcoin.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
