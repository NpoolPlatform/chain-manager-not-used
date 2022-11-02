package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-manager/pkg/db/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CoinBase holds the schema definition for the CoinBase entity.
type CoinBase struct {
	ent.Schema
}

func (CoinBase) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the CoinBase.
func (CoinBase) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			String("name").
			Optional().
			Default(""),
		field.
			String("logo").
			Optional().
			Default(""),
		field.
			Bool("presale").
			Optional().
			Default(false),
		field.
			String("unit").
			Optional().
			Default(""),
		field.
			String("env").
			Optional().
			Default(""),
		field.
			Other("reserved_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Bool("for_pay").
			Optional().
			Default(false),
	}
}

// Edges of the CoinBase.
func (CoinBase) Edges() []ent.Edge {
	return nil
}