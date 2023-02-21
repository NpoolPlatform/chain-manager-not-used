package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-manager/pkg/db/mixin"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	tran "github.com/NpoolPlatform/message/npool/chain/mgr/v1/tx"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Tran holds the schema definition for the Tran entity.
type Tran struct {
	ent.Schema
}

func (Tran) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Tran.
func (Tran) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("from_account_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("to_account_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Other("amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("fee_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			String("chain_tx_id").
			Optional().
			Default(""),
		field.
			String("state").
			Optional().
			Default(tran.TxState_DefaultTxState.String()),
		field.
			String("extra").
			Optional().
			Default(""),
		field.
			String("type").
			Optional().
			Default(basetypes.TxType_DefaultTxType.String()),
	}
}

// Edges of the Tran.
func (Tran) Edges() []ent.Edge {
	return nil
}
