package api

import (
	chainmgr "github.com/NpoolPlatform/message/npool/chain/mgr/v1"

	appcoin "github.com/NpoolPlatform/chain-manager/api/coin/appcoin"
	coinbase "github.com/NpoolPlatform/chain-manager/api/coin/base"
	exrate "github.com/NpoolPlatform/chain-manager/api/coin/exrate"
	coinextra "github.com/NpoolPlatform/chain-manager/api/coin/extra"

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
	coinbase.Register(server)
	coinextra.Register(server)
	appcoin.Register(server)
	exrate.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
