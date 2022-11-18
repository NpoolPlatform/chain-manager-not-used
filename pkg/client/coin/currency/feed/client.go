//nolint:dupl
package currencyfeed

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/feed"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get currencyfeed connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateCurrencyFeed(ctx context.Context, in *npool.CurrencyFeedReq) (*npool.CurrencyFeed, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateCurrencyFeed(ctx, &npool.CreateCurrencyFeedRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create currencyfeed: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create currencyfeed: %v", err)
	}
	return info.(*npool.CurrencyFeed), nil
}

func CreateCurrencyFeeds(ctx context.Context, in []*npool.CurrencyFeedReq) ([]*npool.CurrencyFeed, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateCurrencyFeeds(ctx, &npool.CreateCurrencyFeedsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create currencyfeeds: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create currencyfeeds: %v", err)
	}
	return infos.([]*npool.CurrencyFeed), nil
}

func UpdateCurrencyFeed(ctx context.Context, in *npool.CurrencyFeedReq) (*npool.CurrencyFeed, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateCurrencyFeed(ctx, &npool.UpdateCurrencyFeedRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add currencyfeed: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update currencyfeed: %v", err)
	}
	return info.(*npool.CurrencyFeed), nil
}

func GetCurrencyFeed(ctx context.Context, id string) (*npool.CurrencyFeed, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetCurrencyFeed(ctx, &npool.GetCurrencyFeedRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get currencyfeed: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get currencyfeed: %v", err)
	}
	return info.(*npool.CurrencyFeed), nil
}

func GetCurrencyFeedOnly(ctx context.Context, conds *npool.Conds) (*npool.CurrencyFeed, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetCurrencyFeedOnly(ctx, &npool.GetCurrencyFeedOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get currencyfeed: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get currencyfeed: %v", err)
	}
	return info.(*npool.CurrencyFeed), nil
}

func GetCurrencyFeeds(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.CurrencyFeed, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetCurrencyFeeds(ctx, &npool.GetCurrencyFeedsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get currencyfeeds: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get currencyfeeds: %v", err)
	}
	return infos.([]*npool.CurrencyFeed), total, nil
}

func ExistCurrencyFeed(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistCurrencyFeed(ctx, &npool.ExistCurrencyFeedRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get currencyfeed: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get currencyfeed: %v", err)
	}
	return infos.(bool), nil
}

func ExistCurrencyFeedConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistCurrencyFeedConds(ctx, &npool.ExistCurrencyFeedCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get currencyfeed: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get currencyfeed: %v", err)
	}
	return infos.(bool), nil
}

func CountCurrencyFeeds(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountCurrencyFeeds(ctx, &npool.CountCurrencyFeedsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count currencyfeed: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count currencyfeed: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteCurrencyFeed(ctx context.Context, id string) (*npool.CurrencyFeed, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteCurrencyFeed(ctx, &npool.DeleteCurrencyFeedRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete currencyfeed: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete currencyfeed: %v", err)
	}
	return infos.(*npool.CurrencyFeed), nil
}
