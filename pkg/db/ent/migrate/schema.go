// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CoinBasesColumns holds the columns for the "coin_bases" table.
	CoinBasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "logo", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "presale", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "unit", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "env", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "reserved_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "for_pay", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// CoinBasesTable holds the schema information for the "coin_bases" table.
	CoinBasesTable = &schema.Table{
		Name:       "coin_bases",
		Columns:    CoinBasesColumns,
		PrimaryKey: []*schema.Column{CoinBasesColumns[0]},
	}
	// CoinExtrasColumns holds the columns for the "coin_extras" table.
	CoinExtrasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "home_page", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// CoinExtrasTable holds the schema information for the "coin_extras" table.
	CoinExtrasTable = &schema.Table{
		Name:       "coin_extras",
		Columns:    CoinExtrasColumns,
		PrimaryKey: []*schema.Column{CoinExtrasColumns[0]},
	}
	// TransColumns holds the columns for the "trans" table.
	TransColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "from_account_id", Type: field.TypeUUID, Nullable: true},
		{Name: "to_account_id", Type: field.TypeUUID, Nullable: true},
		{Name: "amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "fee_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "chain_tx_id", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "state", Type: field.TypeString, Nullable: true, Default: "DefaultTxState"},
		{Name: "extra", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// TransTable holds the schema information for the "trans" table.
	TransTable = &schema.Table{
		Name:       "trans",
		Columns:    TransColumns,
		PrimaryKey: []*schema.Column{TransColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CoinBasesTable,
		CoinExtrasTable,
		TransTable,
	}
)

func init() {
}
