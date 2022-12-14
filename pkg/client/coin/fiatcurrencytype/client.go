//nolint:dupl
package fiatcurrencytype

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fiatcurrencytype"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get fiatcurrencytype connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateFiatCurrencyType(ctx context.Context, in *npool.FiatCurrencyTypeReq) (*npool.FiatCurrencyType, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateFiatCurrencyType(ctx, &npool.CreateFiatCurrencyTypeRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create fiatcurrencytype: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create fiatcurrencytype: %v", err)
	}
	return info.(*npool.FiatCurrencyType), nil
}

func CreateFiatCurrencyTypes(ctx context.Context, in []*npool.FiatCurrencyTypeReq) ([]*npool.FiatCurrencyType, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateFiatCurrencyTypes(ctx, &npool.CreateFiatCurrencyTypesRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create fiatcurrencytypes: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create fiatcurrencytypes: %v", err)
	}
	return infos.([]*npool.FiatCurrencyType), nil
}

func UpdateFiatCurrencyType(ctx context.Context, in *npool.FiatCurrencyTypeReq) (*npool.FiatCurrencyType, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateFiatCurrencyType(ctx, &npool.UpdateFiatCurrencyTypeRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add fiatcurrencytype: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update fiatcurrencytype: %v", err)
	}
	return info.(*npool.FiatCurrencyType), nil
}

func GetFiatCurrencyType(ctx context.Context, id string) (*npool.FiatCurrencyType, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetFiatCurrencyType(ctx, &npool.GetFiatCurrencyTypeRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get fiatcurrencytype: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get fiatcurrencytype: %v", err)
	}
	return info.(*npool.FiatCurrencyType), nil
}

func GetFiatCurrencyTypeOnly(ctx context.Context, conds *npool.Conds) (*npool.FiatCurrencyType, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetFiatCurrencyTypeOnly(ctx, &npool.GetFiatCurrencyTypeOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get fiatcurrencytype: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get fiatcurrencytype: %v", err)
	}
	return info.(*npool.FiatCurrencyType), nil
}

func GetFiatCurrencyTypes(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.FiatCurrencyType, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetFiatCurrencyTypes(ctx, &npool.GetFiatCurrencyTypesRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get fiatcurrencytypes: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get fiatcurrencytypes: %v", err)
	}
	return infos.([]*npool.FiatCurrencyType), total, nil
}

func ExistFiatCurrencyType(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistFiatCurrencyType(ctx, &npool.ExistFiatCurrencyTypeRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get fiatcurrencytype: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get fiatcurrencytype: %v", err)
	}
	return infos.(bool), nil
}

func ExistFiatCurrencyTypeConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistFiatCurrencyTypeConds(ctx, &npool.ExistFiatCurrencyTypeCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get fiatcurrencytype: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get fiatcurrencytype: %v", err)
	}
	return infos.(bool), nil
}

func CountFiatCurrencyTypes(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountFiatCurrencyTypes(ctx, &npool.CountFiatCurrencyTypesRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count fiatcurrencytype: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count fiatcurrencytype: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteFiatCurrencyType(ctx context.Context, id string) (*npool.FiatCurrencyType, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteFiatCurrencyType(ctx, &npool.DeleteFiatCurrencyTypeRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete fiatcurrencytype: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete fiatcurrencytype: %v", err)
	}
	return infos.(*npool.FiatCurrencyType), nil
}
