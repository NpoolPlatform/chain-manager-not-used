package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-manager/pkg/db/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency"
)

// LegalCurrency holds the schema definition for the LegalCurrency entity.
type LegalCurrency struct {
	ent.Schema
}

func (LegalCurrency) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the LegalCurrency.
func (LegalCurrency) Fields() []ent.Field {
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
			String("feed_type").
			Optional().
			Default(npool.FeedType_DefaultFeedType.String()),
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

// Edges of the LegalCurrency.
func (LegalCurrency) Edges() []ent.Edge {
	return nil
}
