package api

import (
	chainmgr "github.com/NpoolPlatform/message/npool/chain/mgr/v1"

	"github.com/NpoolPlatform/chain-manager/api/tx"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	chainmgr.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	chainmgr.RegisterManagerServer(server, &Server{})
	tx.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
