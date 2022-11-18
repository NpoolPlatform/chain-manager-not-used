// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/currencyfeed"
	"github.com/google/uuid"
)

// CurrencyFeed is the model entity for the CurrencyFeed schema.
type CurrencyFeed struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// FeedSource holds the value of the "feed_source" field.
	FeedSource string `json:"feed_source,omitempty"`
	// FeedType holds the value of the "feed_type" field.
	FeedType string `json:"feed_type,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CurrencyFeed) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case currencyfeed.FieldCreatedAt, currencyfeed.FieldUpdatedAt, currencyfeed.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case currencyfeed.FieldFeedSource, currencyfeed.FieldFeedType:
			values[i] = new(sql.NullString)
		case currencyfeed.FieldID, currencyfeed.FieldCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CurrencyFeed", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CurrencyFeed fields.
func (cf *CurrencyFeed) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case currencyfeed.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				cf.ID = *value
			}
		case currencyfeed.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cf.CreatedAt = uint32(value.Int64)
			}
		case currencyfeed.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cf.UpdatedAt = uint32(value.Int64)
			}
		case currencyfeed.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				cf.DeletedAt = uint32(value.Int64)
			}
		case currencyfeed.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				cf.CoinTypeID = *value
			}
		case currencyfeed.FieldFeedSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field feed_source", values[i])
			} else if value.Valid {
				cf.FeedSource = value.String
			}
		case currencyfeed.FieldFeedType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field feed_type", values[i])
			} else if value.Valid {
				cf.FeedType = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CurrencyFeed.
// Note that you need to call CurrencyFeed.Unwrap() before calling this method if this CurrencyFeed
// was returned from a transaction, and the transaction was committed or rolled back.
func (cf *CurrencyFeed) Update() *CurrencyFeedUpdateOne {
	return (&CurrencyFeedClient{config: cf.config}).UpdateOne(cf)
}

// Unwrap unwraps the CurrencyFeed entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cf *CurrencyFeed) Unwrap() *CurrencyFeed {
	_tx, ok := cf.config.driver.(*txDriver)
	if !ok {
		panic("ent: CurrencyFeed is not a transactional entity")
	}
	cf.config.driver = _tx.drv
	return cf
}

// String implements the fmt.Stringer.
func (cf *CurrencyFeed) String() string {
	var builder strings.Builder
	builder.WriteString("CurrencyFeed(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cf.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", cf.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", cf.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", cf.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", cf.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("feed_source=")
	builder.WriteString(cf.FeedSource)
	builder.WriteString(", ")
	builder.WriteString("feed_type=")
	builder.WriteString(cf.FeedType)
	builder.WriteByte(')')
	return builder.String()
}

// CurrencyFeeds is a parsable slice of CurrencyFeed.
type CurrencyFeeds []*CurrencyFeed

func (cf CurrencyFeeds) config(cfg config) {
	for _i := range cf {
		cf[_i].config = cfg
	}
}
