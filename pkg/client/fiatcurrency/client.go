//nolint:dupl
package fiatcurrency

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fiatcurrency"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get fiatcurrency connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateFiatCurrency(ctx context.Context, in *npool.FiatCurrencyReq) (*npool.FiatCurrency, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateFiatCurrency(ctx, &npool.CreateFiatCurrencyRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create fiatcurrency: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create fiatcurrency: %v", err)
	}
	return info.(*npool.FiatCurrency), nil
}

func CreateFiatCurrencies(ctx context.Context, in []*npool.FiatCurrencyReq) ([]*npool.FiatCurrency, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateFiatCurrencies(ctx, &npool.CreateFiatCurrenciesRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create fiatcurrencys: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create fiatcurrencys: %v", err)
	}
	return infos.([]*npool.FiatCurrency), nil
}

func UpdateFiatCurrency(ctx context.Context, in *npool.FiatCurrencyReq) (*npool.FiatCurrency, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateFiatCurrency(ctx, &npool.UpdateFiatCurrencyRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add fiatcurrency: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update fiatcurrency: %v", err)
	}
	return info.(*npool.FiatCurrency), nil
}

func GetFiatCurrency(ctx context.Context, id string) (*npool.FiatCurrency, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetFiatCurrency(ctx, &npool.GetFiatCurrencyRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get fiatcurrency: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get fiatcurrency: %v", err)
	}
	return info.(*npool.FiatCurrency), nil
}

func GetFiatCurrencyOnly(ctx context.Context, conds *npool.Conds) (*npool.FiatCurrency, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetFiatCurrencyOnly(ctx, &npool.GetFiatCurrencyOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get fiatcurrency: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get fiatcurrency: %v", err)
	}
	return info.(*npool.FiatCurrency), nil
}

func GetFiatCurrencies(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.FiatCurrency, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetFiatCurrencies(ctx, &npool.GetFiatCurrenciesRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get fiatcurrencys: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get fiatcurrencys: %v", err)
	}
	return infos.([]*npool.FiatCurrency), total, nil
}

func ExistFiatCurrency(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistFiatCurrency(ctx, &npool.ExistFiatCurrencyRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get fiatcurrency: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get fiatcurrency: %v", err)
	}
	return infos.(bool), nil
}

func ExistFiatCurrencyConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistFiatCurrencyConds(ctx, &npool.ExistFiatCurrencyCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get fiatcurrency: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get fiatcurrency: %v", err)
	}
	return infos.(bool), nil
}

func CountFiatCurrencies(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountFiatCurrencies(ctx, &npool.CountFiatCurrenciesRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count fiatcurrency: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count fiatcurrency: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteFiatCurrency(ctx context.Context, id string) (*npool.FiatCurrency, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteFiatCurrency(ctx, &npool.DeleteFiatCurrencyRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete fiatcurrency: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete fiatcurrency: %v", err)
	}
	return infos.(*npool.FiatCurrency), nil
}
