// Code generated by ent, DO NOT EDIT.

package exchangerate

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the exchangerate type in the database.
	Label = "exchange_rate"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldCoinTypeID holds the string denoting the coin_type_id field in the database.
	FieldCoinTypeID = "coin_type_id"
	// FieldMarketValue holds the string denoting the market_value field in the database.
	FieldMarketValue = "market_value"
	// FieldSettleValue holds the string denoting the settle_value field in the database.
	FieldSettleValue = "settle_value"
	// FieldSettlePercent holds the string denoting the settle_percent field in the database.
	FieldSettlePercent = "settle_percent"
	// FieldSetter holds the string denoting the setter field in the database.
	FieldSetter = "setter"
	// Table holds the table name of the exchangerate in the database.
	Table = "exchange_rates"
)

// Columns holds all SQL columns for exchangerate fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAppID,
	FieldCoinTypeID,
	FieldMarketValue,
	FieldSettleValue,
	FieldSettlePercent,
	FieldSetter,
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
	// DefaultAppID holds the default value on creation for the "app_id" field.
	DefaultAppID func() uuid.UUID
	// DefaultCoinTypeID holds the default value on creation for the "coin_type_id" field.
	DefaultCoinTypeID func() uuid.UUID
	// DefaultMarketValue holds the default value on creation for the "market_value" field.
	DefaultMarketValue decimal.Decimal
	// DefaultSettleValue holds the default value on creation for the "settle_value" field.
	DefaultSettleValue decimal.Decimal
	// DefaultSettlePercent holds the default value on creation for the "settle_percent" field.
	DefaultSettlePercent uint32
	// DefaultSetter holds the default value on creation for the "setter" field.
	DefaultSetter func() uuid.UUID
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
