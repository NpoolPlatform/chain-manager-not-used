package currencyvalue

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

	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/value"
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

var currencyValue = &npool.CurrencyValue{
	ID:              uuid.NewString(),
	CoinTypeID:      uuid.NewString(),
	FeedSourceID:    uuid.NewString(),
	MarketValueHigh: "99.123",
	MarketValueLow:  "97.123",
}

var currencyValueReq = &npool.CurrencyValueReq{
	ID:              &currencyValue.ID,
	CoinTypeID:      &currencyValue.CoinTypeID,
	FeedSourceID:    &currencyValue.FeedSourceID,
	MarketValueHigh: &currencyValue.MarketValueHigh,
	MarketValueLow:  &currencyValue.MarketValueLow,
}

func createCurrencyValue(t *testing.T) {
	info, err := CreateCurrencyValue(context.Background(), currencyValueReq)
	if assert.Nil(t, err) {
		currencyValue.CreatedAt = info.CreatedAt
		currencyValue.UpdatedAt = info.UpdatedAt
		assert.Equal(t, currencyValue, info)
	}
}

func updateCurrencyValue(t *testing.T) {
}

func createCurrencyValues(t *testing.T) {
}

func getCurrencyValue(t *testing.T) {
	info, err := GetCurrencyValue(context.Background(), currencyValue.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, currencyValue, info)
	}
}

func deleteCurrencyValue(t *testing.T) {
	info, err := DeleteCurrencyValue(context.Background(), currencyValue.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, currencyValue, info)
	}

	_, err = GetCurrencyValue(context.Background(), info.ID)
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

	t.Run("createCurrencyValue", createCurrencyValue)
	t.Run("updateCurrencyValue", updateCurrencyValue)
	t.Run("createCurrencyValues", createCurrencyValues)
	t.Run("getCurrencyValue", getCurrencyValue)
	t.Run("deleteCurrencyValue", deleteCurrencyValue)
}
