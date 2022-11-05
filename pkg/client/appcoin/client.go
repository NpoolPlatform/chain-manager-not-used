//nolint:dupl
package appcoin

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get appcoin connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateAppCoin(ctx context.Context, in *npool.AppCoinReq) (*npool.AppCoin, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateAppCoin(ctx, &npool.CreateAppCoinRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create appcoin: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create appcoin: %v", err)
	}
	return info.(*npool.AppCoin), nil
}

func CreateAppCoins(ctx context.Context, in []*npool.AppCoinReq) ([]*npool.AppCoin, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateAppCoins(ctx, &npool.CreateAppCoinsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create appcoins: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create appcoins: %v", err)
	}
	return infos.([]*npool.AppCoin), nil
}

func UpdateAppCoin(ctx context.Context, in *npool.AppCoinReq) (*npool.AppCoin, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateAppCoin(ctx, &npool.UpdateAppCoinRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add appcoin: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update appcoin: %v", err)
	}
	return info.(*npool.AppCoin), nil
}

func GetAppCoin(ctx context.Context, id string) (*npool.AppCoin, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetAppCoin(ctx, &npool.GetAppCoinRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get appcoin: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get appcoin: %v", err)
	}
	return info.(*npool.AppCoin), nil
}

func GetAppCoinOnly(ctx context.Context, conds *npool.Conds) (*npool.AppCoin, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetAppCoinOnly(ctx, &npool.GetAppCoinOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get appcoin: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get appcoin: %v", err)
	}
	return info.(*npool.AppCoin), nil
}

func GetAppCoins(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.AppCoin, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetAppCoins(ctx, &npool.GetAppCoinsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get appcoins: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get appcoins: %v", err)
	}
	return infos.([]*npool.AppCoin), total, nil
}

func ExistAppCoin(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistAppCoin(ctx, &npool.ExistAppCoinRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get appcoin: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get appcoin: %v", err)
	}
	return infos.(bool), nil
}

func ExistAppCoinConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistAppCoinConds(ctx, &npool.ExistAppCoinCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get appcoin: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get appcoin: %v", err)
	}
	return infos.(bool), nil
}

func CountAppCoins(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountAppCoins(ctx, &npool.CountAppCoinsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count appcoin: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count appcoin: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteAppCoin(ctx context.Context, id string) (*npool.AppCoin, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteAppCoin(ctx, &npool.DeleteAppCoinRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete appcoin: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete appcoin: %v", err)
	}
	return infos.(*npool.AppCoin), nil
}
