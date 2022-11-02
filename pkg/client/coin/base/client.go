//nolint:dupl
package coinbase

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/base"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get coinbase connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateCoinBase(ctx context.Context, in *npool.CoinBaseReq) (*npool.CoinBase, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateCoinBase(ctx, &npool.CreateCoinBaseRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create coinbase: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create coinbase: %v", err)
	}
	return info.(*npool.CoinBase), nil
}

func CreateCoinBases(ctx context.Context, in []*npool.CoinBaseReq) ([]*npool.CoinBase, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateCoinBases(ctx, &npool.CreateCoinBasesRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create coinbases: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create coinbases: %v", err)
	}
	return infos.([]*npool.CoinBase), nil
}

func UpdateCoinBase(ctx context.Context, in *npool.CoinBaseReq) (*npool.CoinBase, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateCoinBase(ctx, &npool.UpdateCoinBaseRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add coinbase: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update coinbase: %v", err)
	}
	return info.(*npool.CoinBase), nil
}

func GetCoinBase(ctx context.Context, id string) (*npool.CoinBase, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetCoinBase(ctx, &npool.GetCoinBaseRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get coinbase: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get coinbase: %v", err)
	}
	return info.(*npool.CoinBase), nil
}

func GetCoinBaseOnly(ctx context.Context, conds *npool.Conds) (*npool.CoinBase, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetCoinBaseOnly(ctx, &npool.GetCoinBaseOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get coinbase: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get coinbase: %v", err)
	}
	return info.(*npool.CoinBase), nil
}

func GetCoinBases(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.CoinBase, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetCoinBases(ctx, &npool.GetCoinBasesRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get coinbases: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get coinbases: %v", err)
	}
	return infos.([]*npool.CoinBase), total, nil
}

func ExistCoinBase(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistCoinBase(ctx, &npool.ExistCoinBaseRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get coinbase: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get coinbase: %v", err)
	}
	return infos.(bool), nil
}

func ExistCoinBaseConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistCoinBaseConds(ctx, &npool.ExistCoinBaseCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get coinbase: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get coinbase: %v", err)
	}
	return infos.(bool), nil
}

func CountCoinBases(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountCoinBases(ctx, &npool.CountCoinBasesRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count coinbase: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count coinbase: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteCoinBase(ctx context.Context, id string) (*npool.CoinBase, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteCoinBase(ctx, &npool.DeleteCoinBaseRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete coinbase: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete coinbase: %v", err)
	}
	return infos.(*npool.CoinBase), nil
}
