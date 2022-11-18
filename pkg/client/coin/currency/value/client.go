//nolint:dupl
package currencyvalue

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/value"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get currencyvalue connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateCurrencyValue(ctx context.Context, in *npool.CurrencyValueReq) (*npool.CurrencyValue, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateCurrencyValue(ctx, &npool.CreateCurrencyValueRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create currencyvalue: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create currencyvalue: %v", err)
	}
	return info.(*npool.CurrencyValue), nil
}

func CreateCurrencyValues(ctx context.Context, in []*npool.CurrencyValueReq) ([]*npool.CurrencyValue, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateCurrencyValues(ctx, &npool.CreateCurrencyValuesRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create currencyvalues: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create currencyvalues: %v", err)
	}
	return infos.([]*npool.CurrencyValue), nil
}

func UpdateCurrencyValue(ctx context.Context, in *npool.CurrencyValueReq) (*npool.CurrencyValue, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateCurrencyValue(ctx, &npool.UpdateCurrencyValueRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add currencyvalue: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update currencyvalue: %v", err)
	}
	return info.(*npool.CurrencyValue), nil
}

func GetCurrencyValue(ctx context.Context, id string) (*npool.CurrencyValue, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetCurrencyValue(ctx, &npool.GetCurrencyValueRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get currencyvalue: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get currencyvalue: %v", err)
	}
	return info.(*npool.CurrencyValue), nil
}

func GetCurrencyValueOnly(ctx context.Context, conds *npool.Conds) (*npool.CurrencyValue, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetCurrencyValueOnly(ctx, &npool.GetCurrencyValueOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get currencyvalue: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get currencyvalue: %v", err)
	}
	return info.(*npool.CurrencyValue), nil
}

func GetCurrencyValues(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.CurrencyValue, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetCurrencyValues(ctx, &npool.GetCurrencyValuesRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get currencyvalues: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get currencyvalues: %v", err)
	}
	return infos.([]*npool.CurrencyValue), total, nil
}

func ExistCurrencyValue(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistCurrencyValue(ctx, &npool.ExistCurrencyValueRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get currencyvalue: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get currencyvalue: %v", err)
	}
	return infos.(bool), nil
}

func ExistCurrencyValueConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistCurrencyValueConds(ctx, &npool.ExistCurrencyValueCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get currencyvalue: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get currencyvalue: %v", err)
	}
	return infos.(bool), nil
}

func CountCurrencyValues(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountCurrencyValues(ctx, &npool.CountCurrencyValuesRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count currencyvalue: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count currencyvalue: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteCurrencyValue(ctx context.Context, id string) (*npool.CurrencyValue, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteCurrencyValue(ctx, &npool.DeleteCurrencyValueRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete currencyvalue: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete currencyvalue: %v", err)
	}
	return infos.(*npool.CurrencyValue), nil
}
