package description

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

	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/description"
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

var coinDescription = &npool.CoinDescription{
	ID:         uuid.NewString(),
	AppID:      uuid.NewString(),
	CoinTypeID: uuid.NewString(),
	UsedFor:    npool.UsedFor_ProductPage,
	Title:      uuid.NewString(),
	Message:    uuid.NewString(),
}

var coinDescriptionReq = &npool.CoinDescriptionReq{
	ID:         &coinDescription.ID,
	AppID:      &coinDescription.AppID,
	CoinTypeID: &coinDescription.CoinTypeID,
	UsedFor:    &coinDescription.UsedFor,
	Title:      &coinDescription.Title,
	Message:    &coinDescription.Message,
}

func createCoinDescription(t *testing.T) {
	info, err := CreateCoinDescription(context.Background(), coinDescriptionReq)
	if assert.Nil(t, err) {
		coinDescription.CreatedAt = info.CreatedAt
		coinDescription.UpdatedAt = info.UpdatedAt
		assert.Equal(t, coinDescription, info)
	}
}

func updateCoinDescription(t *testing.T) {
	_title := uuid.NewString()
	_message := uuid.NewString()

	coinDescriptionReq.Title = &_title
	coinDescriptionReq.Message = &_message

	coinDescription.Title = _title
	coinDescription.Message = _message

	info, err := UpdateCoinDescription(context.Background(), coinDescriptionReq)
	if assert.Nil(t, err) {
		coinDescription.UpdatedAt = info.UpdatedAt
		assert.Equal(t, coinDescription, info)
	}
}

func createCoinDescriptions(t *testing.T) {
}

func getCoinDescription(t *testing.T) {
	info, err := GetCoinDescription(context.Background(), coinDescription.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, coinDescription, info)
	}
}

func deleteCoinDescription(t *testing.T) {
	info, err := DeleteCoinDescription(context.Background(), coinDescription.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, coinDescription, info)
	}

	_, err = GetCoinDescription(context.Background(), info.ID)
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

	t.Run("createCoinDescription", createCoinDescription)
	t.Run("updateCoinDescription", updateCoinDescription)
	t.Run("createCoinDescriptions", createCoinDescriptions)
	t.Run("getCoinDescription", getCoinDescription)
	t.Run("deleteCoinDescription", deleteCoinDescription)
}
