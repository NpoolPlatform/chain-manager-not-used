package description

import (
	description "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/description"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	description.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	description.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
