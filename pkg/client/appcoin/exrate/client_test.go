package exrate

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

	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/exrate"
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

var exchangeRate = &npool.ExchangeRate{
	ID:            uuid.NewString(),
	AppID:         uuid.NewString(),
	CoinTypeID:    uuid.NewString(),
	MarketValue:   "89.123",
	SettlePercent: 80,
	Setter:        uuid.NewString(),
}

var exchangeRateReq = &npool.ExchangeRateReq{
	ID:            &exchangeRate.ID,
	AppID:         &exchangeRate.AppID,
	CoinTypeID:    &exchangeRate.CoinTypeID,
	MarketValue:   &exchangeRate.MarketValue,
	SettlePercent: &exchangeRate.SettlePercent,
	Setter:        &exchangeRate.Setter,
}

func createExchangeRate(t *testing.T) {
	info, err := CreateExchangeRate(context.Background(), exchangeRateReq)
	if assert.Nil(t, err) {
		exchangeRate.CreatedAt = info.CreatedAt
		exchangeRate.UpdatedAt = info.UpdatedAt
		exchangeRate.SettleValue = info.SettleValue
		assert.Equal(t, exchangeRate, info)
	}
}

func updateExchangeRate(t *testing.T) {
	marketValue := "1000.001"
	settlePercent := uint32(90)

	exchangeRateReq.MarketValue = &marketValue
	exchangeRateReq.SettlePercent = &settlePercent

	exchangeRate.MarketValue = marketValue
	exchangeRate.SettlePercent = settlePercent

	info, err := UpdateExchangeRate(context.Background(), exchangeRateReq)
	if assert.Nil(t, err) {
		exchangeRate.SettleValue = info.SettleValue
		exchangeRate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, exchangeRate, info)
	}
}

func createExchangeRates(t *testing.T) {
}

func getExchangeRate(t *testing.T) {
	info, err := GetExchangeRate(context.Background(), exchangeRate.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exchangeRate, info)
	}
}

func deleteExchangeRate(t *testing.T) {
	info, err := DeleteExchangeRate(context.Background(), exchangeRate.ID)
	if assert.Nil(t, err) {
		exchangeRate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, exchangeRate, info)
	}

	_, err = GetExchangeRate(context.Background(), info.ID)
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

	t.Run("createExchangeRate", createExchangeRate)
	t.Run("updateExchangeRate", updateExchangeRate)
	t.Run("createExchangeRates", createExchangeRates)
	t.Run("getExchangeRate", getExchangeRate)
	t.Run("deleteExchangeRate", deleteExchangeRate)
}
