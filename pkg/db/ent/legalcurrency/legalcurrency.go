// Code generated by ent, DO NOT EDIT.

package legalcurrency

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the legalcurrency type in the database.
	Label = "legal_currency"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldCoinTypeID holds the string denoting the coin_type_id field in the database.
	FieldCoinTypeID = "coin_type_id"
	// FieldFeedType holds the string denoting the feed_type field in the database.
	FieldFeedType = "feed_type"
	// FieldMarketValueHigh holds the string denoting the market_value_high field in the database.
	FieldMarketValueHigh = "market_value_high"
	// FieldMarketValueLow holds the string denoting the market_value_low field in the database.
	FieldMarketValueLow = "market_value_low"
	// Table holds the table name of the legalcurrency in the database.
	Table = "legal_currencies"
)

// Columns holds all SQL columns for legalcurrency fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldCoinTypeID,
	FieldFeedType,
	FieldMarketValueHigh,
	FieldMarketValueLow,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/chain-manager/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultCoinTypeID holds the default value on creation for the "coin_type_id" field.
	DefaultCoinTypeID func() uuid.UUID
	// DefaultFeedType holds the default value on creation for the "feed_type" field.
	DefaultFeedType string
	// DefaultMarketValueHigh holds the default value on creation for the "market_value_high" field.
	DefaultMarketValueHigh decimal.Decimal
	// DefaultMarketValueLow holds the default value on creation for the "market_value_low" field.
	DefaultMarketValueLow decimal.Decimal
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
