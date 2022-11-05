package api

import (
	chainmgr "github.com/NpoolPlatform/message/npool/chain/mgr/v1"

	appcoin "github.com/NpoolPlatform/chain-manager/api/appcoin"
	description "github.com/NpoolPlatform/chain-manager/api/appcoin/description"
	exrate "github.com/NpoolPlatform/chain-manager/api/appcoin/exrate"
	coinbase "github.com/NpoolPlatform/chain-manager/api/coin/base"
	coinextra "github.com/NpoolPlatform/chain-manager/api/coin/extra"
	setting "github.com/NpoolPlatform/chain-manager/api/coin/setting"

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
	description.Register(server)
	setting.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
