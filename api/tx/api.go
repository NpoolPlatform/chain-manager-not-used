package tx

import (
	"github.com/NpoolPlatform/message/npool/chain/mgr/v1/tx"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	tx.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	tx.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
