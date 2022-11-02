package coinbase

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

	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/base"
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

var coinBase = &npool.CoinBase{
	ID:             uuid.NewString(),
	Name:           uuid.NewString(),
	Logo:           uuid.NewString(),
	Presale:        false,
	Unit:           "BTC",
	ENV:            "main",
	ReservedAmount: "89.123",
	ForPay:         false,
}

var coinBaseReq = &npool.CoinBaseReq{
	ID:             &coinBase.ID,
	Name:           &coinBase.Name,
	Logo:           &coinBase.Logo,
	Presale:        &coinBase.Presale,
	Unit:           &coinBase.Unit,
	ENV:            &coinBase.ENV,
	ReservedAmount: &coinBase.ReservedAmount,
	ForPay:         &coinBase.ForPay,
}

func createCoinBase(t *testing.T) {
	info, err := CreateCoinBase(context.Background(), coinBaseReq)
	if assert.Nil(t, err) {
		coinBase.CreatedAt = info.CreatedAt
		coinBase.UpdatedAt = info.UpdatedAt
		assert.Equal(t, coinBase, info)
	}
}

func updateCoinBase(t *testing.T) {
	presale := false
	reservedAmount := "81.0234"
	forPay := true

	coinBaseReq.Presale = &presale
	coinBaseReq.ReservedAmount = &reservedAmount
	coinBaseReq.ForPay = &forPay

	coinBase.Presale = presale
	coinBase.ReservedAmount = reservedAmount
	coinBase.ForPay = forPay

	info, err := UpdateCoinBase(context.Background(), coinBaseReq)
	if assert.Nil(t, err) {
		coinBase.UpdatedAt = info.UpdatedAt
		assert.Equal(t, coinBase, info)
	}
}

func createCoinBases(t *testing.T) {
}

func getCoinBase(t *testing.T) {
	info, err := GetCoinBase(context.Background(), coinBase.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, coinBase, info)
	}
}

func deleteCoinBase(t *testing.T) {
	info, err := DeleteCoinBase(context.Background(), coinBase.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, coinBase, info)
	}

	_, err = GetCoinBase(context.Background(), info.ID)
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

	t.Run("createCoinBase", createCoinBase)
	t.Run("updateCoinBase", updateCoinBase)
	t.Run("createCoinBases", createCoinBases)
	t.Run("getCoinBase", getCoinBase)
	t.Run("deleteCoinBase", deleteCoinBase)
}
