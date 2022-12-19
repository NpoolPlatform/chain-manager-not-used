package fiatcurrency

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

	"github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/fiatcurrency"

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

var fiatcurrencyValue = &npool.FiatCurrency{
	ID:                 uuid.NewString(),
	FiatCurrencyTypeID: uuid.NewString(),
	FeedType:           currency.FeedType_CoinBase,
	MarketValueHigh:    "99.123",
	MarketValueLow:     "97.123",
}

var fiatcurrencyValueReq = &npool.FiatCurrencyReq{
	ID:                 &fiatcurrencyValue.ID,
	FiatCurrencyTypeID: &fiatcurrencyValue.FiatCurrencyTypeID,
	FeedType:           &fiatcurrencyValue.FeedType,
	MarketValueHigh:    &fiatcurrencyValue.MarketValueHigh,
	MarketValueLow:     &fiatcurrencyValue.MarketValueLow,
}

func createFiatCurrency(t *testing.T) {
	info, err := CreateFiatCurrency(context.Background(), fiatcurrencyValueReq)
	if assert.Nil(t, err) {
		fiatcurrencyValue.CreatedAt = info.CreatedAt
		fiatcurrencyValue.UpdatedAt = info.UpdatedAt
		assert.Equal(t, fiatcurrencyValue, info)
	}
}

func updateFiatCurrency(t *testing.T) {
}

func createCurrencies(t *testing.T) {
}

func getFiatCurrency(t *testing.T) {
	info, err := GetFiatCurrency(context.Background(), fiatcurrencyValue.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, fiatcurrencyValue, info)
	}
}

func deleteFiatCurrency(t *testing.T) {
	info, err := DeleteFiatCurrency(context.Background(), fiatcurrencyValue.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, fiatcurrencyValue, info)
	}

	_, err = GetFiatCurrency(context.Background(), info.ID)
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

	t.Run("createFiatCurrency", createFiatCurrency)
	t.Run("updateFiatCurrency", updateFiatCurrency)
	t.Run("createCurrencies", createCurrencies)
	t.Run("getFiatCurrency", getFiatCurrency)
	t.Run("deleteFiatCurrency", deleteFiatCurrency)
}
