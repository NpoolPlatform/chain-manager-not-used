//nolint:dupl
package setting

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/setting"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get setting connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateSetting(ctx context.Context, in *npool.SettingReq) (*npool.Setting, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateSetting(ctx, &npool.CreateSettingRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create setting: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create setting: %v", err)
	}
	return info.(*npool.Setting), nil
}

func CreateSettings(ctx context.Context, in []*npool.SettingReq) ([]*npool.Setting, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateSettings(ctx, &npool.CreateSettingsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create settings: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create settings: %v", err)
	}
	return infos.([]*npool.Setting), nil
}

func UpdateSetting(ctx context.Context, in *npool.SettingReq) (*npool.Setting, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateSetting(ctx, &npool.UpdateSettingRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add setting: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update setting: %v", err)
	}
	return info.(*npool.Setting), nil
}

func GetSetting(ctx context.Context, id string) (*npool.Setting, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetSetting(ctx, &npool.GetSettingRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get setting: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get setting: %v", err)
	}
	return info.(*npool.Setting), nil
}

func GetSettingOnly(ctx context.Context, conds *npool.Conds) (*npool.Setting, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetSettingOnly(ctx, &npool.GetSettingOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get setting: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get setting: %v", err)
	}
	return info.(*npool.Setting), nil
}

func GetSettings(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.Setting, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetSettings(ctx, &npool.GetSettingsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get settings: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get settings: %v", err)
	}
	return infos.([]*npool.Setting), total, nil
}

func ExistSetting(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistSetting(ctx, &npool.ExistSettingRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get setting: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get setting: %v", err)
	}
	return infos.(bool), nil
}

func ExistSettingConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistSettingConds(ctx, &npool.ExistSettingCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get setting: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get setting: %v", err)
	}
	return infos.(bool), nil
}

func CountSettings(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountSettings(ctx, &npool.CountSettingsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count setting: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count setting: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteSetting(ctx context.Context, id string) (*npool.Setting, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteSetting(ctx, &npool.DeleteSettingRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete setting: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete setting: %v", err)
	}
	return infos.(*npool.Setting), nil
}
