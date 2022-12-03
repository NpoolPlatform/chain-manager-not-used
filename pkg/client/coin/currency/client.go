//nolint:dupl
package currency

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get currency connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateCurrency(ctx context.Context, in *npool.CurrencyReq) (*npool.Currency, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateCurrency(ctx, &npool.CreateCurrencyRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create currency: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create currency: %v", err)
	}
	return info.(*npool.Currency), nil
}

func CreateCurrencies(ctx context.Context, in []*npool.CurrencyReq) ([]*npool.Currency, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateCurrencies(ctx, &npool.CreateCurrenciesRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create currencys: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create currencys: %v", err)
	}
	return infos.([]*npool.Currency), nil
}

func UpdateCurrency(ctx context.Context, in *npool.CurrencyReq) (*npool.Currency, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateCurrency(ctx, &npool.UpdateCurrencyRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add currency: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update currency: %v", err)
	}
	return info.(*npool.Currency), nil
}

func GetCurrency(ctx context.Context, id string) (*npool.Currency, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetCurrency(ctx, &npool.GetCurrencyRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get currency: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get currency: %v", err)
	}
	return info.(*npool.Currency), nil
}

func GetCurrencyOnly(ctx context.Context, conds *npool.Conds) (*npool.Currency, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetCurrencyOnly(ctx, &npool.GetCurrencyOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get currency: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get currency: %v", err)
	}
	return info.(*npool.Currency), nil
}

func GetCurrencies(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.Currency, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetCurrencies(ctx, &npool.GetCurrenciesRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get currencys: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get currencys: %v", err)
	}
	return infos.([]*npool.Currency), total, nil
}

func ExistCurrency(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistCurrency(ctx, &npool.ExistCurrencyRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get currency: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get currency: %v", err)
	}
	return infos.(bool), nil
}

func ExistCurrencyConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistCurrencyConds(ctx, &npool.ExistCurrencyCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get currency: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get currency: %v", err)
	}
	return infos.(bool), nil
}

func CountCurrencies(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountCurrencies(ctx, &npool.CountCurrenciesRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count currency: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count currency: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteCurrency(ctx context.Context, id string) (*npool.Currency, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteCurrency(ctx, &npool.DeleteCurrencyRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete currency: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete currency: %v", err)
	}
	return infos.(*npool.Currency), nil
}
