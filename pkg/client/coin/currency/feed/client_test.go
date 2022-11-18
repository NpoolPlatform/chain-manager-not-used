package currencyfeed

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

	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/feed"
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

var currencyFeed = &npool.CurrencyFeed{
	ID:         uuid.NewString(),
	CoinTypeID: uuid.NewString(),
	FeedSource: uuid.NewString(),
	FeedType:   npool.FeedType_CoinBase,
}

var currencyFeedReq = &npool.CurrencyFeedReq{
	ID:         &currencyFeed.ID,
	CoinTypeID: &currencyFeed.CoinTypeID,
	FeedSource: &currencyFeed.FeedSource,
	FeedType:   &currencyFeed.FeedType,
}

func createCurrencyFeed(t *testing.T) {
	info, err := CreateCurrencyFeed(context.Background(), currencyFeedReq)
	if assert.Nil(t, err) {
		currencyFeed.CreatedAt = info.CreatedAt
		currencyFeed.UpdatedAt = info.UpdatedAt
		assert.Equal(t, currencyFeed, info)
	}
}

func updateCurrencyFeed(t *testing.T) {
	feedSource := uuid.NewString()
	currencyFeedReq.FeedSource = &feedSource
	currencyFeed.FeedSource = feedSource

	info, err := UpdateCurrencyFeed(context.Background(), currencyFeedReq)
	if assert.Nil(t, err) {
		currencyFeed.UpdatedAt = info.UpdatedAt
		assert.Equal(t, currencyFeed, info)
	}
}

func createCurrencyFeeds(t *testing.T) {
}

func getCurrencyFeed(t *testing.T) {
	info, err := GetCurrencyFeed(context.Background(), currencyFeed.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, currencyFeed, info)
	}
}

func deleteCurrencyFeed(t *testing.T) {
	info, err := DeleteCurrencyFeed(context.Background(), currencyFeed.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, currencyFeed, info)
	}

	_, err = GetCurrencyFeed(context.Background(), info.ID)
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

	t.Run("createCurrencyFeed", createCurrencyFeed)
	t.Run("updateCurrencyFeed", updateCurrencyFeed)
	t.Run("createCurrencyFeeds", createCurrencyFeeds)
	t.Run("getCurrencyFeed", getCurrencyFeed)
	t.Run("deleteCurrencyFeed", deleteCurrencyFeed)
}
