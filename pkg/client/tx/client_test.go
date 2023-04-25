package tx

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

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/tx"

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

var ret = &npool.Tx{
	ID:            uuid.NewString(),
	CoinTypeID:    uuid.NewString(),
	FromAccountID: uuid.NewString(),
	ToAccountID:   uuid.NewString(),
	Amount:        "89.5678",
	FeeAmount:     "2.5678",
	ChainTxID:     uuid.UUID{}.String(),
	State:         npool.TxState_StateCreated,
	Extra:         uuid.NewString(),
	Type:          basetypes.TxType_TxPaymentCollect,
}

var req = &npool.TxReq{
	ID:            &ret.ID,
	CoinTypeID:    &ret.CoinTypeID,
	FromAccountID: &ret.FromAccountID,
	ToAccountID:   &ret.ToAccountID,
	Amount:        &ret.Amount,
	FeeAmount:     &ret.FeeAmount,
	ChainTxID:     &ret.ChainTxID,
	State:         &ret.State,
	Extra:         &ret.Extra,
	Type:          &ret.Type,
}

func createTx(t *testing.T) {
	info, err := CreateTx(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func updateTx(t *testing.T) {
	state := npool.TxState_StateWait

	req.State = &state
	ret.State = state

	info, err := UpdateTx(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func createTxs(t *testing.T) {
}

func getTx(t *testing.T) {
	info, err := GetTx(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deleteTx(t *testing.T) {
	info, err := DeleteTx(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}

	_, err = GetTx(context.Background(), info.ID)
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

	t.Run("createTx", createTx)
	t.Run("updateTx", updateTx)
	t.Run("createTxs", createTxs)
	t.Run("getTx", getTx)
	t.Run("deleteTx", deleteTx)
}
