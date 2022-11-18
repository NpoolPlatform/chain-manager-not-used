package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-manager/pkg/db/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CurrencyValue holds the schema definition for the CurrencyValue entity.
type CurrencyValue struct {
	ent.Schema
}

func (CurrencyValue) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the CurrencyValue.
func (CurrencyValue) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			UUID("feed_source_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			Other("market_value_high", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("market_value_low", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
	}
}

// Edges of the CurrencyValue.
func (CurrencyValue) Edges() []ent.Edge {
	return nil
}
