package coinextra

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

	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/extra"
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

var coinExtra = &npool.CoinExtra{
	ID:         uuid.NewString(),
	CoinTypeID: uuid.NewString(),
	HomePage:   uuid.NewString(),
}

var coinExtraReq = &npool.CoinExtraReq{
	ID:         &coinExtra.ID,
	CoinTypeID: &coinExtra.CoinTypeID,
	HomePage:   &coinExtra.HomePage,
}

func createCoinExtra(t *testing.T) {
	info, err := CreateCoinExtra(context.Background(), coinExtraReq)
	if assert.Nil(t, err) {
		coinExtra.CreatedAt = info.CreatedAt
		coinExtra.UpdatedAt = info.UpdatedAt
		assert.Equal(t, coinExtra, info)
	}
}

func updateCoinExtra(t *testing.T) {
	homePage := uuid.NewString()

	coinExtraReq.HomePage = &homePage

	coinExtra.HomePage = homePage

	info, err := UpdateCoinExtra(context.Background(), coinExtraReq)
	if assert.Nil(t, err) {
		coinExtra.UpdatedAt = info.UpdatedAt
		assert.Equal(t, coinExtra, info)
	}
}

func createCoinExtras(t *testing.T) {
}

func getCoinExtra(t *testing.T) {
	info, err := GetCoinExtra(context.Background(), coinExtra.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, coinExtra, info)
	}
}

func deleteCoinExtra(t *testing.T) {
	info, err := DeleteCoinExtra(context.Background(), coinExtra.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, coinExtra, info)
	}

	_, err = GetCoinExtra(context.Background(), info.ID)
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

	t.Run("createCoinExtra", createCoinExtra)
	t.Run("updateCoinExtra", updateCoinExtra)
	t.Run("createCoinExtras", createCoinExtras)
	t.Run("getCoinExtra", getCoinExtra)
	t.Run("deleteCoinExtra", deleteCoinExtra)
}
