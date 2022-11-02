//nolint:dupl
package coinextra

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/extra"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get coinextra connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateCoinExtra(ctx context.Context, in *npool.CoinExtraReq) (*npool.CoinExtra, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateCoinExtra(ctx, &npool.CreateCoinExtraRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create coinextra: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create coinextra: %v", err)
	}
	return info.(*npool.CoinExtra), nil
}

func CreateCoinExtras(ctx context.Context, in []*npool.CoinExtraReq) ([]*npool.CoinExtra, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateCoinExtras(ctx, &npool.CreateCoinExtrasRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create coinextras: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create coinextras: %v", err)
	}
	return infos.([]*npool.CoinExtra), nil
}

func UpdateCoinExtra(ctx context.Context, in *npool.CoinExtraReq) (*npool.CoinExtra, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateCoinExtra(ctx, &npool.UpdateCoinExtraRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add coinextra: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update coinextra: %v", err)
	}
	return info.(*npool.CoinExtra), nil
}

func GetCoinExtra(ctx context.Context, id string) (*npool.CoinExtra, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetCoinExtra(ctx, &npool.GetCoinExtraRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get coinextra: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get coinextra: %v", err)
	}
	return info.(*npool.CoinExtra), nil
}

func GetCoinExtraOnly(ctx context.Context, conds *npool.Conds) (*npool.CoinExtra, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetCoinExtraOnly(ctx, &npool.GetCoinExtraOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get coinextra: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get coinextra: %v", err)
	}
	return info.(*npool.CoinExtra), nil
}

func GetCoinExtras(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.CoinExtra, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetCoinExtras(ctx, &npool.GetCoinExtrasRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get coinextras: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get coinextras: %v", err)
	}
	return infos.([]*npool.CoinExtra), total, nil
}

func ExistCoinExtra(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistCoinExtra(ctx, &npool.ExistCoinExtraRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get coinextra: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get coinextra: %v", err)
	}
	return infos.(bool), nil
}

func ExistCoinExtraConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistCoinExtraConds(ctx, &npool.ExistCoinExtraCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get coinextra: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get coinextra: %v", err)
	}
	return infos.(bool), nil
}

func CountCoinExtras(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountCoinExtras(ctx, &npool.CountCoinExtrasRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count coinextra: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count coinextra: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteCoinExtra(ctx context.Context, id string) (*npool.CoinExtra, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteCoinExtra(ctx, &npool.DeleteCoinExtraRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete coinextra: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete coinextra: %v", err)
	}
	return infos.(*npool.CoinExtra), nil
}
