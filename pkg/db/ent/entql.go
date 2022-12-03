// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/appcoin"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinbase"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/coindescription"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinextra"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/currency"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/exchangerate"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/setting"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/tran"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 8)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   appcoin.Table,
			Columns: appcoin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appcoin.FieldID,
			},
		},
		Type: "AppCoin",
		Fields: map[string]*sqlgraph.FieldSpec{
			appcoin.FieldCreatedAt:                {Type: field.TypeUint32, Column: appcoin.FieldCreatedAt},
			appcoin.FieldUpdatedAt:                {Type: field.TypeUint32, Column: appcoin.FieldUpdatedAt},
			appcoin.FieldDeletedAt:                {Type: field.TypeUint32, Column: appcoin.FieldDeletedAt},
			appcoin.FieldAppID:                    {Type: field.TypeUUID, Column: appcoin.FieldAppID},
			appcoin.FieldCoinTypeID:               {Type: field.TypeUUID, Column: appcoin.FieldCoinTypeID},
			appcoin.FieldName:                     {Type: field.TypeString, Column: appcoin.FieldName},
			appcoin.FieldLogo:                     {Type: field.TypeString, Column: appcoin.FieldLogo},
			appcoin.FieldForPay:                   {Type: field.TypeBool, Column: appcoin.FieldForPay},
			appcoin.FieldWithdrawAutoReviewAmount: {Type: field.TypeOther, Column: appcoin.FieldWithdrawAutoReviewAmount},
			appcoin.FieldProductPage:              {Type: field.TypeString, Column: appcoin.FieldProductPage},
			appcoin.FieldDisabled:                 {Type: field.TypeBool, Column: appcoin.FieldDisabled},
			appcoin.FieldDailyRewardAmount:        {Type: field.TypeOther, Column: appcoin.FieldDailyRewardAmount},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   coinbase.Table,
			Columns: coinbase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coinbase.FieldID,
			},
		},
		Type: "CoinBase",
		Fields: map[string]*sqlgraph.FieldSpec{
			coinbase.FieldCreatedAt:      {Type: field.TypeUint32, Column: coinbase.FieldCreatedAt},
			coinbase.FieldUpdatedAt:      {Type: field.TypeUint32, Column: coinbase.FieldUpdatedAt},
			coinbase.FieldDeletedAt:      {Type: field.TypeUint32, Column: coinbase.FieldDeletedAt},
			coinbase.FieldName:           {Type: field.TypeString, Column: coinbase.FieldName},
			coinbase.FieldLogo:           {Type: field.TypeString, Column: coinbase.FieldLogo},
			coinbase.FieldPresale:        {Type: field.TypeBool, Column: coinbase.FieldPresale},
			coinbase.FieldUnit:           {Type: field.TypeString, Column: coinbase.FieldUnit},
			coinbase.FieldEnv:            {Type: field.TypeString, Column: coinbase.FieldEnv},
			coinbase.FieldReservedAmount: {Type: field.TypeOther, Column: coinbase.FieldReservedAmount},
			coinbase.FieldForPay:         {Type: field.TypeBool, Column: coinbase.FieldForPay},
			coinbase.FieldDisabled:       {Type: field.TypeBool, Column: coinbase.FieldDisabled},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   coindescription.Table,
			Columns: coindescription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coindescription.FieldID,
			},
		},
		Type: "CoinDescription",
		Fields: map[string]*sqlgraph.FieldSpec{
			coindescription.FieldCreatedAt:  {Type: field.TypeUint32, Column: coindescription.FieldCreatedAt},
			coindescription.FieldUpdatedAt:  {Type: field.TypeUint32, Column: coindescription.FieldUpdatedAt},
			coindescription.FieldDeletedAt:  {Type: field.TypeUint32, Column: coindescription.FieldDeletedAt},
			coindescription.FieldAppID:      {Type: field.TypeUUID, Column: coindescription.FieldAppID},
			coindescription.FieldCoinTypeID: {Type: field.TypeUUID, Column: coindescription.FieldCoinTypeID},
			coindescription.FieldUsedFor:    {Type: field.TypeString, Column: coindescription.FieldUsedFor},
			coindescription.FieldTitle:      {Type: field.TypeString, Column: coindescription.FieldTitle},
			coindescription.FieldMessage:    {Type: field.TypeString, Column: coindescription.FieldMessage},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   coinextra.Table,
			Columns: coinextra.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coinextra.FieldID,
			},
		},
		Type: "CoinExtra",
		Fields: map[string]*sqlgraph.FieldSpec{
			coinextra.FieldCreatedAt:  {Type: field.TypeUint32, Column: coinextra.FieldCreatedAt},
			coinextra.FieldUpdatedAt:  {Type: field.TypeUint32, Column: coinextra.FieldUpdatedAt},
			coinextra.FieldDeletedAt:  {Type: field.TypeUint32, Column: coinextra.FieldDeletedAt},
			coinextra.FieldCoinTypeID: {Type: field.TypeUUID, Column: coinextra.FieldCoinTypeID},
			coinextra.FieldHomePage:   {Type: field.TypeString, Column: coinextra.FieldHomePage},
			coinextra.FieldSpecs:      {Type: field.TypeString, Column: coinextra.FieldSpecs},
			coinextra.FieldStableUsd:  {Type: field.TypeBool, Column: coinextra.FieldStableUsd},
		},
	}
	graph.Nodes[4] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   currency.Table,
			Columns: currency.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: currency.FieldID,
			},
		},
		Type: "Currency",
		Fields: map[string]*sqlgraph.FieldSpec{
			currency.FieldCreatedAt:       {Type: field.TypeUint32, Column: currency.FieldCreatedAt},
			currency.FieldUpdatedAt:       {Type: field.TypeUint32, Column: currency.FieldUpdatedAt},
			currency.FieldDeletedAt:       {Type: field.TypeUint32, Column: currency.FieldDeletedAt},
			currency.FieldCoinTypeID:      {Type: field.TypeUUID, Column: currency.FieldCoinTypeID},
			currency.FieldFeedType:        {Type: field.TypeString, Column: currency.FieldFeedType},
			currency.FieldMarketValueHigh: {Type: field.TypeOther, Column: currency.FieldMarketValueHigh},
			currency.FieldMarketValueLow:  {Type: field.TypeOther, Column: currency.FieldMarketValueLow},
		},
	}
	graph.Nodes[5] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   exchangerate.Table,
			Columns: exchangerate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: exchangerate.FieldID,
			},
		},
		Type: "ExchangeRate",
		Fields: map[string]*sqlgraph.FieldSpec{
			exchangerate.FieldCreatedAt:     {Type: field.TypeUint32, Column: exchangerate.FieldCreatedAt},
			exchangerate.FieldUpdatedAt:     {Type: field.TypeUint32, Column: exchangerate.FieldUpdatedAt},
			exchangerate.FieldDeletedAt:     {Type: field.TypeUint32, Column: exchangerate.FieldDeletedAt},
			exchangerate.FieldAppID:         {Type: field.TypeUUID, Column: exchangerate.FieldAppID},
			exchangerate.FieldCoinTypeID:    {Type: field.TypeUUID, Column: exchangerate.FieldCoinTypeID},
			exchangerate.FieldMarketValue:   {Type: field.TypeOther, Column: exchangerate.FieldMarketValue},
			exchangerate.FieldSettleValue:   {Type: field.TypeOther, Column: exchangerate.FieldSettleValue},
			exchangerate.FieldSettlePercent: {Type: field.TypeUint32, Column: exchangerate.FieldSettlePercent},
			exchangerate.FieldSetter:        {Type: field.TypeUUID, Column: exchangerate.FieldSetter},
		},
	}
	graph.Nodes[6] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   setting.Table,
			Columns: setting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: setting.FieldID,
			},
		},
		Type: "Setting",
		Fields: map[string]*sqlgraph.FieldSpec{
			setting.FieldCreatedAt:                   {Type: field.TypeUint32, Column: setting.FieldCreatedAt},
			setting.FieldUpdatedAt:                   {Type: field.TypeUint32, Column: setting.FieldUpdatedAt},
			setting.FieldDeletedAt:                   {Type: field.TypeUint32, Column: setting.FieldDeletedAt},
			setting.FieldCoinTypeID:                  {Type: field.TypeUUID, Column: setting.FieldCoinTypeID},
			setting.FieldFeeCoinTypeID:               {Type: field.TypeUUID, Column: setting.FieldFeeCoinTypeID},
			setting.FieldWithdrawFeeByStableUsd:      {Type: field.TypeBool, Column: setting.FieldWithdrawFeeByStableUsd},
			setting.FieldWithdrawFeeAmount:           {Type: field.TypeOther, Column: setting.FieldWithdrawFeeAmount},
			setting.FieldCollectFeeAmount:            {Type: field.TypeOther, Column: setting.FieldCollectFeeAmount},
			setting.FieldHotWalletFeeAmount:          {Type: field.TypeOther, Column: setting.FieldHotWalletFeeAmount},
			setting.FieldLowFeeAmount:                {Type: field.TypeOther, Column: setting.FieldLowFeeAmount},
			setting.FieldHotWalletAccountAmount:      {Type: field.TypeOther, Column: setting.FieldHotWalletAccountAmount},
			setting.FieldPaymentAccountCollectAmount: {Type: field.TypeOther, Column: setting.FieldPaymentAccountCollectAmount},
		},
	}
	graph.Nodes[7] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   tran.Table,
			Columns: tran.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tran.FieldID,
			},
		},
		Type: "Tran",
		Fields: map[string]*sqlgraph.FieldSpec{
			tran.FieldCreatedAt:     {Type: field.TypeUint32, Column: tran.FieldCreatedAt},
			tran.FieldUpdatedAt:     {Type: field.TypeUint32, Column: tran.FieldUpdatedAt},
			tran.FieldDeletedAt:     {Type: field.TypeUint32, Column: tran.FieldDeletedAt},
			tran.FieldCoinTypeID:    {Type: field.TypeUUID, Column: tran.FieldCoinTypeID},
			tran.FieldFromAccountID: {Type: field.TypeUUID, Column: tran.FieldFromAccountID},
			tran.FieldToAccountID:   {Type: field.TypeUUID, Column: tran.FieldToAccountID},
			tran.FieldAmount:        {Type: field.TypeOther, Column: tran.FieldAmount},
			tran.FieldFeeAmount:     {Type: field.TypeOther, Column: tran.FieldFeeAmount},
			tran.FieldChainTxID:     {Type: field.TypeString, Column: tran.FieldChainTxID},
			tran.FieldState:         {Type: field.TypeString, Column: tran.FieldState},
			tran.FieldExtra:         {Type: field.TypeString, Column: tran.FieldExtra},
			tran.FieldType:          {Type: field.TypeString, Column: tran.FieldType},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (acq *AppCoinQuery) addPredicate(pred func(s *sql.Selector)) {
	acq.predicates = append(acq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the AppCoinQuery builder.
func (acq *AppCoinQuery) Filter() *AppCoinFilter {
	return &AppCoinFilter{config: acq.config, predicateAdder: acq}
}

// addPredicate implements the predicateAdder interface.
func (m *AppCoinMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the AppCoinMutation builder.
func (m *AppCoinMutation) Filter() *AppCoinFilter {
	return &AppCoinFilter{config: m.config, predicateAdder: m}
}

// AppCoinFilter provides a generic filtering capability at runtime for AppCoinQuery.
type AppCoinFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *AppCoinFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *AppCoinFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(appcoin.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *AppCoinFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(appcoin.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *AppCoinFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(appcoin.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *AppCoinFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(appcoin.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *AppCoinFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(appcoin.FieldAppID))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *AppCoinFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(appcoin.FieldCoinTypeID))
}

// WhereName applies the entql string predicate on the name field.
func (f *AppCoinFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(appcoin.FieldName))
}

// WhereLogo applies the entql string predicate on the logo field.
func (f *AppCoinFilter) WhereLogo(p entql.StringP) {
	f.Where(p.Field(appcoin.FieldLogo))
}

// WhereForPay applies the entql bool predicate on the for_pay field.
func (f *AppCoinFilter) WhereForPay(p entql.BoolP) {
	f.Where(p.Field(appcoin.FieldForPay))
}

// WhereWithdrawAutoReviewAmount applies the entql other predicate on the withdraw_auto_review_amount field.
func (f *AppCoinFilter) WhereWithdrawAutoReviewAmount(p entql.OtherP) {
	f.Where(p.Field(appcoin.FieldWithdrawAutoReviewAmount))
}

// WhereProductPage applies the entql string predicate on the product_page field.
func (f *AppCoinFilter) WhereProductPage(p entql.StringP) {
	f.Where(p.Field(appcoin.FieldProductPage))
}

// WhereDisabled applies the entql bool predicate on the disabled field.
func (f *AppCoinFilter) WhereDisabled(p entql.BoolP) {
	f.Where(p.Field(appcoin.FieldDisabled))
}

// WhereDailyRewardAmount applies the entql other predicate on the daily_reward_amount field.
func (f *AppCoinFilter) WhereDailyRewardAmount(p entql.OtherP) {
	f.Where(p.Field(appcoin.FieldDailyRewardAmount))
}

// addPredicate implements the predicateAdder interface.
func (cbq *CoinBaseQuery) addPredicate(pred func(s *sql.Selector)) {
	cbq.predicates = append(cbq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CoinBaseQuery builder.
func (cbq *CoinBaseQuery) Filter() *CoinBaseFilter {
	return &CoinBaseFilter{config: cbq.config, predicateAdder: cbq}
}

// addPredicate implements the predicateAdder interface.
func (m *CoinBaseMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CoinBaseMutation builder.
func (m *CoinBaseMutation) Filter() *CoinBaseFilter {
	return &CoinBaseFilter{config: m.config, predicateAdder: m}
}

// CoinBaseFilter provides a generic filtering capability at runtime for CoinBaseQuery.
type CoinBaseFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *CoinBaseFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *CoinBaseFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(coinbase.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *CoinBaseFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(coinbase.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *CoinBaseFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(coinbase.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *CoinBaseFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(coinbase.FieldDeletedAt))
}

// WhereName applies the entql string predicate on the name field.
func (f *CoinBaseFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(coinbase.FieldName))
}

// WhereLogo applies the entql string predicate on the logo field.
func (f *CoinBaseFilter) WhereLogo(p entql.StringP) {
	f.Where(p.Field(coinbase.FieldLogo))
}

// WherePresale applies the entql bool predicate on the presale field.
func (f *CoinBaseFilter) WherePresale(p entql.BoolP) {
	f.Where(p.Field(coinbase.FieldPresale))
}

// WhereUnit applies the entql string predicate on the unit field.
func (f *CoinBaseFilter) WhereUnit(p entql.StringP) {
	f.Where(p.Field(coinbase.FieldUnit))
}

// WhereEnv applies the entql string predicate on the env field.
func (f *CoinBaseFilter) WhereEnv(p entql.StringP) {
	f.Where(p.Field(coinbase.FieldEnv))
}

// WhereReservedAmount applies the entql other predicate on the reserved_amount field.
func (f *CoinBaseFilter) WhereReservedAmount(p entql.OtherP) {
	f.Where(p.Field(coinbase.FieldReservedAmount))
}

// WhereForPay applies the entql bool predicate on the for_pay field.
func (f *CoinBaseFilter) WhereForPay(p entql.BoolP) {
	f.Where(p.Field(coinbase.FieldForPay))
}

// WhereDisabled applies the entql bool predicate on the disabled field.
func (f *CoinBaseFilter) WhereDisabled(p entql.BoolP) {
	f.Where(p.Field(coinbase.FieldDisabled))
}

// addPredicate implements the predicateAdder interface.
func (cdq *CoinDescriptionQuery) addPredicate(pred func(s *sql.Selector)) {
	cdq.predicates = append(cdq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CoinDescriptionQuery builder.
func (cdq *CoinDescriptionQuery) Filter() *CoinDescriptionFilter {
	return &CoinDescriptionFilter{config: cdq.config, predicateAdder: cdq}
}

// addPredicate implements the predicateAdder interface.
func (m *CoinDescriptionMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CoinDescriptionMutation builder.
func (m *CoinDescriptionMutation) Filter() *CoinDescriptionFilter {
	return &CoinDescriptionFilter{config: m.config, predicateAdder: m}
}

// CoinDescriptionFilter provides a generic filtering capability at runtime for CoinDescriptionQuery.
type CoinDescriptionFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *CoinDescriptionFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *CoinDescriptionFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(coindescription.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *CoinDescriptionFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(coindescription.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *CoinDescriptionFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(coindescription.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *CoinDescriptionFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(coindescription.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *CoinDescriptionFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(coindescription.FieldAppID))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *CoinDescriptionFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(coindescription.FieldCoinTypeID))
}

// WhereUsedFor applies the entql string predicate on the used_for field.
func (f *CoinDescriptionFilter) WhereUsedFor(p entql.StringP) {
	f.Where(p.Field(coindescription.FieldUsedFor))
}

// WhereTitle applies the entql string predicate on the title field.
func (f *CoinDescriptionFilter) WhereTitle(p entql.StringP) {
	f.Where(p.Field(coindescription.FieldTitle))
}

// WhereMessage applies the entql string predicate on the message field.
func (f *CoinDescriptionFilter) WhereMessage(p entql.StringP) {
	f.Where(p.Field(coindescription.FieldMessage))
}

// addPredicate implements the predicateAdder interface.
func (ceq *CoinExtraQuery) addPredicate(pred func(s *sql.Selector)) {
	ceq.predicates = append(ceq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CoinExtraQuery builder.
func (ceq *CoinExtraQuery) Filter() *CoinExtraFilter {
	return &CoinExtraFilter{config: ceq.config, predicateAdder: ceq}
}

// addPredicate implements the predicateAdder interface.
func (m *CoinExtraMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CoinExtraMutation builder.
func (m *CoinExtraMutation) Filter() *CoinExtraFilter {
	return &CoinExtraFilter{config: m.config, predicateAdder: m}
}

// CoinExtraFilter provides a generic filtering capability at runtime for CoinExtraQuery.
type CoinExtraFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *CoinExtraFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *CoinExtraFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(coinextra.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *CoinExtraFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(coinextra.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *CoinExtraFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(coinextra.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *CoinExtraFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(coinextra.FieldDeletedAt))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *CoinExtraFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(coinextra.FieldCoinTypeID))
}

// WhereHomePage applies the entql string predicate on the home_page field.
func (f *CoinExtraFilter) WhereHomePage(p entql.StringP) {
	f.Where(p.Field(coinextra.FieldHomePage))
}

// WhereSpecs applies the entql string predicate on the specs field.
func (f *CoinExtraFilter) WhereSpecs(p entql.StringP) {
	f.Where(p.Field(coinextra.FieldSpecs))
}

// WhereStableUsd applies the entql bool predicate on the stable_usd field.
func (f *CoinExtraFilter) WhereStableUsd(p entql.BoolP) {
	f.Where(p.Field(coinextra.FieldStableUsd))
}

// addPredicate implements the predicateAdder interface.
func (cq *CurrencyQuery) addPredicate(pred func(s *sql.Selector)) {
	cq.predicates = append(cq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CurrencyQuery builder.
func (cq *CurrencyQuery) Filter() *CurrencyFilter {
	return &CurrencyFilter{config: cq.config, predicateAdder: cq}
}

// addPredicate implements the predicateAdder interface.
func (m *CurrencyMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CurrencyMutation builder.
func (m *CurrencyMutation) Filter() *CurrencyFilter {
	return &CurrencyFilter{config: m.config, predicateAdder: m}
}

// CurrencyFilter provides a generic filtering capability at runtime for CurrencyQuery.
type CurrencyFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *CurrencyFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[4].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *CurrencyFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(currency.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *CurrencyFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(currency.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *CurrencyFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(currency.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *CurrencyFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(currency.FieldDeletedAt))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *CurrencyFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(currency.FieldCoinTypeID))
}

// WhereFeedType applies the entql string predicate on the feed_type field.
func (f *CurrencyFilter) WhereFeedType(p entql.StringP) {
	f.Where(p.Field(currency.FieldFeedType))
}

// WhereMarketValueHigh applies the entql other predicate on the market_value_high field.
func (f *CurrencyFilter) WhereMarketValueHigh(p entql.OtherP) {
	f.Where(p.Field(currency.FieldMarketValueHigh))
}

// WhereMarketValueLow applies the entql other predicate on the market_value_low field.
func (f *CurrencyFilter) WhereMarketValueLow(p entql.OtherP) {
	f.Where(p.Field(currency.FieldMarketValueLow))
}

// addPredicate implements the predicateAdder interface.
func (erq *ExchangeRateQuery) addPredicate(pred func(s *sql.Selector)) {
	erq.predicates = append(erq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ExchangeRateQuery builder.
func (erq *ExchangeRateQuery) Filter() *ExchangeRateFilter {
	return &ExchangeRateFilter{config: erq.config, predicateAdder: erq}
}

// addPredicate implements the predicateAdder interface.
func (m *ExchangeRateMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ExchangeRateMutation builder.
func (m *ExchangeRateMutation) Filter() *ExchangeRateFilter {
	return &ExchangeRateFilter{config: m.config, predicateAdder: m}
}

// ExchangeRateFilter provides a generic filtering capability at runtime for ExchangeRateQuery.
type ExchangeRateFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *ExchangeRateFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[5].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *ExchangeRateFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(exchangerate.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *ExchangeRateFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(exchangerate.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *ExchangeRateFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(exchangerate.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *ExchangeRateFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(exchangerate.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *ExchangeRateFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(exchangerate.FieldAppID))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *ExchangeRateFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(exchangerate.FieldCoinTypeID))
}

// WhereMarketValue applies the entql other predicate on the market_value field.
func (f *ExchangeRateFilter) WhereMarketValue(p entql.OtherP) {
	f.Where(p.Field(exchangerate.FieldMarketValue))
}

// WhereSettleValue applies the entql other predicate on the settle_value field.
func (f *ExchangeRateFilter) WhereSettleValue(p entql.OtherP) {
	f.Where(p.Field(exchangerate.FieldSettleValue))
}

// WhereSettlePercent applies the entql uint32 predicate on the settle_percent field.
func (f *ExchangeRateFilter) WhereSettlePercent(p entql.Uint32P) {
	f.Where(p.Field(exchangerate.FieldSettlePercent))
}

// WhereSetter applies the entql [16]byte predicate on the setter field.
func (f *ExchangeRateFilter) WhereSetter(p entql.ValueP) {
	f.Where(p.Field(exchangerate.FieldSetter))
}

// addPredicate implements the predicateAdder interface.
func (sq *SettingQuery) addPredicate(pred func(s *sql.Selector)) {
	sq.predicates = append(sq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the SettingQuery builder.
func (sq *SettingQuery) Filter() *SettingFilter {
	return &SettingFilter{config: sq.config, predicateAdder: sq}
}

// addPredicate implements the predicateAdder interface.
func (m *SettingMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the SettingMutation builder.
func (m *SettingMutation) Filter() *SettingFilter {
	return &SettingFilter{config: m.config, predicateAdder: m}
}

// SettingFilter provides a generic filtering capability at runtime for SettingQuery.
type SettingFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *SettingFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[6].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *SettingFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(setting.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *SettingFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(setting.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *SettingFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(setting.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *SettingFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(setting.FieldDeletedAt))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *SettingFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(setting.FieldCoinTypeID))
}

// WhereFeeCoinTypeID applies the entql [16]byte predicate on the fee_coin_type_id field.
func (f *SettingFilter) WhereFeeCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(setting.FieldFeeCoinTypeID))
}

// WhereWithdrawFeeByStableUsd applies the entql bool predicate on the withdraw_fee_by_stable_usd field.
func (f *SettingFilter) WhereWithdrawFeeByStableUsd(p entql.BoolP) {
	f.Where(p.Field(setting.FieldWithdrawFeeByStableUsd))
}

// WhereWithdrawFeeAmount applies the entql other predicate on the withdraw_fee_amount field.
func (f *SettingFilter) WhereWithdrawFeeAmount(p entql.OtherP) {
	f.Where(p.Field(setting.FieldWithdrawFeeAmount))
}

// WhereCollectFeeAmount applies the entql other predicate on the collect_fee_amount field.
func (f *SettingFilter) WhereCollectFeeAmount(p entql.OtherP) {
	f.Where(p.Field(setting.FieldCollectFeeAmount))
}

// WhereHotWalletFeeAmount applies the entql other predicate on the hot_wallet_fee_amount field.
func (f *SettingFilter) WhereHotWalletFeeAmount(p entql.OtherP) {
	f.Where(p.Field(setting.FieldHotWalletFeeAmount))
}

// WhereLowFeeAmount applies the entql other predicate on the low_fee_amount field.
func (f *SettingFilter) WhereLowFeeAmount(p entql.OtherP) {
	f.Where(p.Field(setting.FieldLowFeeAmount))
}

// WhereHotWalletAccountAmount applies the entql other predicate on the hot_wallet_account_amount field.
func (f *SettingFilter) WhereHotWalletAccountAmount(p entql.OtherP) {
	f.Where(p.Field(setting.FieldHotWalletAccountAmount))
}

// WherePaymentAccountCollectAmount applies the entql other predicate on the payment_account_collect_amount field.
func (f *SettingFilter) WherePaymentAccountCollectAmount(p entql.OtherP) {
	f.Where(p.Field(setting.FieldPaymentAccountCollectAmount))
}

// addPredicate implements the predicateAdder interface.
func (tq *TranQuery) addPredicate(pred func(s *sql.Selector)) {
	tq.predicates = append(tq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TranQuery builder.
func (tq *TranQuery) Filter() *TranFilter {
	return &TranFilter{config: tq.config, predicateAdder: tq}
}

// addPredicate implements the predicateAdder interface.
func (m *TranMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the TranMutation builder.
func (m *TranMutation) Filter() *TranFilter {
	return &TranFilter{config: m.config, predicateAdder: m}
}

// TranFilter provides a generic filtering capability at runtime for TranQuery.
type TranFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *TranFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[7].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *TranFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(tran.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *TranFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(tran.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *TranFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(tran.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *TranFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(tran.FieldDeletedAt))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *TranFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(tran.FieldCoinTypeID))
}

// WhereFromAccountID applies the entql [16]byte predicate on the from_account_id field.
func (f *TranFilter) WhereFromAccountID(p entql.ValueP) {
	f.Where(p.Field(tran.FieldFromAccountID))
}

// WhereToAccountID applies the entql [16]byte predicate on the to_account_id field.
func (f *TranFilter) WhereToAccountID(p entql.ValueP) {
	f.Where(p.Field(tran.FieldToAccountID))
}

// WhereAmount applies the entql other predicate on the amount field.
func (f *TranFilter) WhereAmount(p entql.OtherP) {
	f.Where(p.Field(tran.FieldAmount))
}

// WhereFeeAmount applies the entql other predicate on the fee_amount field.
func (f *TranFilter) WhereFeeAmount(p entql.OtherP) {
	f.Where(p.Field(tran.FieldFeeAmount))
}

// WhereChainTxID applies the entql string predicate on the chain_tx_id field.
func (f *TranFilter) WhereChainTxID(p entql.StringP) {
	f.Where(p.Field(tran.FieldChainTxID))
}

// WhereState applies the entql string predicate on the state field.
func (f *TranFilter) WhereState(p entql.StringP) {
	f.Where(p.Field(tran.FieldState))
}

// WhereExtra applies the entql string predicate on the extra field.
func (f *TranFilter) WhereExtra(p entql.StringP) {
	f.Where(p.Field(tran.FieldExtra))
}

// WhereType applies the entql string predicate on the type field.
func (f *TranFilter) WhereType(p entql.StringP) {
	f.Where(p.Field(tran.FieldType))
}
