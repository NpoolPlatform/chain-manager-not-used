package setting

import (
	setting "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/setting"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	setting.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	setting.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
