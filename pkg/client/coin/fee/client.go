//nolint:dupl
package fee

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fee"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get fee connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateFee(ctx context.Context, in *npool.FeeReq) (*npool.Fee, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateFee(ctx, &npool.CreateFeeRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create fee: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create fee: %v", err)
	}
	return info.(*npool.Fee), nil
}

func CreateFees(ctx context.Context, in []*npool.FeeReq) ([]*npool.Fee, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateFees(ctx, &npool.CreateFeesRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create fees: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create fees: %v", err)
	}
	return infos.([]*npool.Fee), nil
}

func UpdateFee(ctx context.Context, in *npool.FeeReq) (*npool.Fee, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateFee(ctx, &npool.UpdateFeeRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add fee: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update fee: %v", err)
	}
	return info.(*npool.Fee), nil
}

func GetFee(ctx context.Context, id string) (*npool.Fee, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetFee(ctx, &npool.GetFeeRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get fee: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get fee: %v", err)
	}
	return info.(*npool.Fee), nil
}

func GetFeeOnly(ctx context.Context, conds *npool.Conds) (*npool.Fee, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetFeeOnly(ctx, &npool.GetFeeOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get fee: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get fee: %v", err)
	}
	return info.(*npool.Fee), nil
}

func GetFees(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.Fee, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetFees(ctx, &npool.GetFeesRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get fees: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get fees: %v", err)
	}
	return infos.([]*npool.Fee), total, nil
}

func ExistFee(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistFee(ctx, &npool.ExistFeeRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get fee: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get fee: %v", err)
	}
	return infos.(bool), nil
}

func ExistFeeConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistFeeConds(ctx, &npool.ExistFeeCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get fee: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get fee: %v", err)
	}
	return infos.(bool), nil
}

func CountFees(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountFees(ctx, &npool.CountFeesRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count fee: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count fee: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteFee(ctx context.Context, id string) (*npool.Fee, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteFee(ctx, &npool.DeleteFeeRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete fee: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete fee: %v", err)
	}
	return infos.(*npool.Fee), nil
}
