//nolint:dupl
package exrate

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/exrate"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get exrate connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateExchangeRate(ctx context.Context, in *npool.ExchangeRateReq) (*npool.ExchangeRate, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateExchangeRate(ctx, &npool.CreateExchangeRateRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create exrate: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create exrate: %v", err)
	}
	return info.(*npool.ExchangeRate), nil
}

func CreateExchangeRates(ctx context.Context, in []*npool.ExchangeRateReq) ([]*npool.ExchangeRate, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateExchangeRates(ctx, &npool.CreateExchangeRatesRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create exrates: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create exrates: %v", err)
	}
	return infos.([]*npool.ExchangeRate), nil
}

func UpdateExchangeRate(ctx context.Context, in *npool.ExchangeRateReq) (*npool.ExchangeRate, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateExchangeRate(ctx, &npool.UpdateExchangeRateRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add exrate: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update exrate: %v", err)
	}
	return info.(*npool.ExchangeRate), nil
}

func GetExchangeRate(ctx context.Context, id string) (*npool.ExchangeRate, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetExchangeRate(ctx, &npool.GetExchangeRateRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get exrate: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get exrate: %v", err)
	}
	return info.(*npool.ExchangeRate), nil
}

func GetExchangeRateOnly(ctx context.Context, conds *npool.Conds) (*npool.ExchangeRate, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetExchangeRateOnly(ctx, &npool.GetExchangeRateOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get exrate: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get exrate: %v", err)
	}
	return info.(*npool.ExchangeRate), nil
}

func GetExchangeRates(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.ExchangeRate, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetExchangeRates(ctx, &npool.GetExchangeRatesRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get exrates: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get exrates: %v", err)
	}
	return infos.([]*npool.ExchangeRate), total, nil
}

func ExistExchangeRate(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistExchangeRate(ctx, &npool.ExistExchangeRateRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get exrate: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get exrate: %v", err)
	}
	return infos.(bool), nil
}

func ExistExchangeRateConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistExchangeRateConds(ctx, &npool.ExistExchangeRateCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get exrate: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get exrate: %v", err)
	}
	return infos.(bool), nil
}

func CountExchangeRates(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountExchangeRates(ctx, &npool.CountExchangeRatesRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count exrate: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count exrate: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteExchangeRate(ctx context.Context, id string) (*npool.ExchangeRate, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteExchangeRate(ctx, &npool.DeleteExchangeRateRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete exrate: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete exrate: %v", err)
	}
	return infos.(*npool.ExchangeRate), nil
}
