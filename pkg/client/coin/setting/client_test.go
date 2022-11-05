package setting

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/chain-manager/pkg/testinit"

	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/setting"
	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var setting = &npool.Setting{
	ID:                          uuid.NewString(),
	CoinTypeID:                  uuid.NewString(),
	FeeCoinTypeID:               uuid.NewString(),
	WithdrawFeeByStableUSD:      true,
	WithdrawFeeAmount:           "2.01",
	CollectFeeAmount:            "2.01",
	HotWalletFeeAmount:          "2.01",
	LowFeeAmount:                "0.201",
	HotWalletAccountAmount:      "9000.123",
	PaymentAccountCollectAmount: "9000.123",
}

var settingReq = &npool.SettingReq{
	ID:                          &setting.ID,
	CoinTypeID:                  &setting.CoinTypeID,
	FeeCoinTypeID:               &setting.FeeCoinTypeID,
	WithdrawFeeByStableUSD:      &setting.WithdrawFeeByStableUSD,
	WithdrawFeeAmount:           &setting.WithdrawFeeAmount,
	CollectFeeAmount:            &setting.CollectFeeAmount,
	HotWalletFeeAmount:          &setting.HotWalletFeeAmount,
	LowFeeAmount:                &setting.LowFeeAmount,
	HotWalletAccountAmount:      &setting.HotWalletAccountAmount,
	PaymentAccountCollectAmount: &setting.PaymentAccountCollectAmount,
}

func createSetting(t *testing.T) {
	info, err := CreateSetting(context.Background(), settingReq)
	if assert.Nil(t, err) {
		setting.CreatedAt = info.CreatedAt
		setting.UpdatedAt = info.UpdatedAt
		assert.Equal(t, setting, info)
	}
}

func updateSetting(t *testing.T) {
	withdrawFeeByStableUSD := false
	withdrawFeeAmount := "22.01"
	collectFeeAmount := "23.01"
	hotWalletFeeAmount := "24.01"
	lowFeeAmount := "2.01"

	settingReq.WithdrawFeeByStableUSD = &withdrawFeeByStableUSD
	settingReq.WithdrawFeeAmount = &withdrawFeeAmount
	settingReq.CollectFeeAmount = &collectFeeAmount
	settingReq.HotWalletFeeAmount = &hotWalletFeeAmount
	settingReq.LowFeeAmount = &lowFeeAmount

	setting.WithdrawFeeByStableUSD = withdrawFeeByStableUSD
	setting.WithdrawFeeAmount = withdrawFeeAmount
	setting.CollectFeeAmount = collectFeeAmount
	setting.HotWalletFeeAmount = hotWalletFeeAmount
	setting.LowFeeAmount = lowFeeAmount

	info, err := UpdateSetting(context.Background(), settingReq)
	if assert.Nil(t, err) {
		setting.UpdatedAt = info.UpdatedAt
		assert.Equal(t, setting, info)
	}
}

func createSettings(t *testing.T) {
}

func getSetting(t *testing.T) {
	info, err := GetSetting(context.Background(), setting.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, setting, info)
	}
}

func deleteSetting(t *testing.T) {
	info, err := DeleteSetting(context.Background(), setting.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, setting, info)
	}

	_, err = GetSetting(context.Background(), info.ID)
	assert.NotNil(t, err)
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createSetting", createSetting)
	t.Run("updateSetting", updateSetting)
	t.Run("createSettings", createSettings)
	t.Run("getSetting", getSetting)
	t.Run("deleteSetting", deleteSetting)
}
