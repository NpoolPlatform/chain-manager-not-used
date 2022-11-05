//nolint:dupl
package description

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/description"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get description connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateCoinDescription(ctx context.Context, in *npool.CoinDescriptionReq) (*npool.CoinDescription, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateCoinDescription(ctx, &npool.CreateCoinDescriptionRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create description: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create description: %v", err)
	}
	return info.(*npool.CoinDescription), nil
}

func CreateCoinDescriptions(ctx context.Context, in []*npool.CoinDescriptionReq) ([]*npool.CoinDescription, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateCoinDescriptions(ctx, &npool.CreateCoinDescriptionsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create descriptions: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create descriptions: %v", err)
	}
	return infos.([]*npool.CoinDescription), nil
}

func UpdateCoinDescription(ctx context.Context, in *npool.CoinDescriptionReq) (*npool.CoinDescription, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateCoinDescription(ctx, &npool.UpdateCoinDescriptionRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add description: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update description: %v", err)
	}
	return info.(*npool.CoinDescription), nil
}

func GetCoinDescription(ctx context.Context, id string) (*npool.CoinDescription, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetCoinDescription(ctx, &npool.GetCoinDescriptionRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get description: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get description: %v", err)
	}
	return info.(*npool.CoinDescription), nil
}

func GetCoinDescriptionOnly(ctx context.Context, conds *npool.Conds) (*npool.CoinDescription, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetCoinDescriptionOnly(ctx, &npool.GetCoinDescriptionOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get description: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get description: %v", err)
	}
	return info.(*npool.CoinDescription), nil
}

func GetCoinDescriptions(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.CoinDescription, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetCoinDescriptions(ctx, &npool.GetCoinDescriptionsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get descriptions: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get descriptions: %v", err)
	}
	return infos.([]*npool.CoinDescription), total, nil
}

func ExistCoinDescription(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistCoinDescription(ctx, &npool.ExistCoinDescriptionRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get description: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get description: %v", err)
	}
	return infos.(bool), nil
}

func ExistCoinDescriptionConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistCoinDescriptionConds(ctx, &npool.ExistCoinDescriptionCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get description: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get description: %v", err)
	}
	return infos.(bool), nil
}

func CountCoinDescriptions(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountCoinDescriptions(ctx, &npool.CountCoinDescriptionsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count description: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count description: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteCoinDescription(ctx context.Context, id string) (*npool.CoinDescription, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteCoinDescription(ctx, &npool.DeleteCoinDescriptionRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete description: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete description: %v", err)
	}
	return infos.(*npool.CoinDescription), nil
}
