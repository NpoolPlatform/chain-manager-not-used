package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-manager/pkg/db/mixin"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/feed"
	"github.com/google/uuid"
)

// CurrencyFeed holds the schema definition for the CurrencyFeed entity.
type CurrencyFeed struct {
	ent.Schema
}

func (CurrencyFeed) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the CurrencyFeed.
func (CurrencyFeed) Fields() []ent.Field {
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
			String("feed_source").
			Optional().
			Default(""),
		field.
			String("feed_type").
			Optional().
			Default(npool.FeedType_DefaultFeedType.String()),
	}
}

// Edges of the CurrencyFeed.
func (CurrencyFeed) Edges() []ent.Edge {
	return nil
}
