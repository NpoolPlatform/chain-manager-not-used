package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-manager/pkg/db/mixin"
	"github.com/google/uuid"
)

// LegalType holds the schema definition for the LegalType entity.
type LegalType struct {
	ent.Schema
}

func (LegalType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the LegalType.
func (LegalType) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("name", uuid.UUID{}).
			Optional().
			Default(uuid.New),
	}
}

// Edges of the LegalType.
func (LegalType) Edges() []ent.Edge {
	return nil
}
