package fiatcurrencytype

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

	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fiatcurrencytype"

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

var fiatcurrencyValue = &npool.FiatCurrencyType{
	ID:   uuid.NewString(),
	Name: uuid.NewString(),
}

var fiatcurrencyValueReq = &npool.FiatCurrencyTypeReq{
	ID:   &fiatcurrencyValue.ID,
	Name: &fiatcurrencyValue.Name,
}

func createFiatCurrencyType(t *testing.T) {
	info, err := CreateFiatCurrencyType(context.Background(), fiatcurrencyValueReq)
	if assert.Nil(t, err) {
		fiatcurrencyValue.CreatedAt = info.CreatedAt
		fiatcurrencyValue.UpdatedAt = info.UpdatedAt
		assert.Equal(t, fiatcurrencyValue, info)
	}
}

func updateFiatCurrencyType(t *testing.T) {
}

func createCurrencies(t *testing.T) {
}

func getFiatCurrencyType(t *testing.T) {
	info, err := GetFiatCurrencyType(context.Background(), fiatcurrencyValue.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, fiatcurrencyValue, info)
	}
}

func deleteFiatCurrencyType(t *testing.T) {
	info, err := DeleteFiatCurrencyType(context.Background(), fiatcurrencyValue.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, fiatcurrencyValue, info)
	}

	_, err = GetFiatCurrencyType(context.Background(), info.ID)
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

	t.Run("createFiatCurrencyType", createFiatCurrencyType)
	t.Run("updateFiatCurrencyType", updateFiatCurrencyType)
	t.Run("createCurrencies", createCurrencies)
	t.Run("getFiatCurrencyType", getFiatCurrencyType)
	t.Run("deleteFiatCurrencyType", deleteFiatCurrencyType)
}
