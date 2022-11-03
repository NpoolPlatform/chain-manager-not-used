package fee

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

	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fee"
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

var fee = &npool.Fee{
	ID:                     uuid.NewString(),
	CoinTypeID:             uuid.NewString(),
	FeeCoinTypeID:          uuid.NewString(),
	WithdrawFeeByStableUSD: true,
	WithdrawFeeAmount:      "2.01",
	CollectFeeAmount:       "2.01",
	HotWalletFeeAmount:     "2.01",
	LowFeeAmount:           "0.201",
}

var feeReq = &npool.FeeReq{
	ID:                     &fee.ID,
	CoinTypeID:             &fee.CoinTypeID,
	FeeCoinTypeID:          &fee.FeeCoinTypeID,
	WithdrawFeeByStableUSD: &fee.WithdrawFeeByStableUSD,
	WithdrawFeeAmount:      &fee.WithdrawFeeAmount,
	CollectFeeAmount:       &fee.CollectFeeAmount,
	HotWalletFeeAmount:     &fee.HotWalletFeeAmount,
	LowFeeAmount:           &fee.LowFeeAmount,
}

func createFee(t *testing.T) {
	info, err := CreateFee(context.Background(), feeReq)
	if assert.Nil(t, err) {
		fee.CreatedAt = info.CreatedAt
		fee.UpdatedAt = info.UpdatedAt
		assert.Equal(t, fee, info)
	}
}

func updateFee(t *testing.T) {
	withdrawFeeByStableUSD := false
	withdrawFeeAmount := "22.01"
	collectFeeAmount := "23.01"
	hotWalletFeeAmount := "24.01"
	lowFeeAmount := "2.01"

	feeReq.WithdrawFeeByStableUSD = &withdrawFeeByStableUSD
	feeReq.WithdrawFeeAmount = &withdrawFeeAmount
	feeReq.CollectFeeAmount = &collectFeeAmount
	feeReq.HotWalletFeeAmount = &hotWalletFeeAmount
	feeReq.LowFeeAmount = &lowFeeAmount

	fee.WithdrawFeeByStableUSD = withdrawFeeByStableUSD
	fee.WithdrawFeeAmount = withdrawFeeAmount
	fee.CollectFeeAmount = collectFeeAmount
	fee.HotWalletFeeAmount = hotWalletFeeAmount
	fee.LowFeeAmount = lowFeeAmount

	info, err := UpdateFee(context.Background(), feeReq)
	if assert.Nil(t, err) {
		fee.UpdatedAt = info.UpdatedAt
		assert.Equal(t, fee, info)
	}
}

func createFees(t *testing.T) {
}

func getFee(t *testing.T) {
	info, err := GetFee(context.Background(), fee.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, fee, info)
	}
}

func deleteFee(t *testing.T) {
	info, err := DeleteFee(context.Background(), fee.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, fee, info)
	}

	_, err = GetFee(context.Background(), info.ID)
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

	t.Run("createFee", createFee)
	t.Run("updateFee", updateFee)
	t.Run("createFees", createFees)
	t.Run("getFee", getFee)
	t.Run("deleteFee", deleteFee)
}
