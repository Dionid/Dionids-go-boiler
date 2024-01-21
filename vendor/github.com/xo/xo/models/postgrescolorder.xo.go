package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// PostgresColOrder is a index column order.
type PostgresColOrder struct {
	Ord string `json:"ord"` // ord
}

// PostgresGetColOrder runs a custom query, returning results as [PostgresColOrder].
func PostgresGetColOrder(ctx context.Context, db DB, schema, index string) (*PostgresColOrder, error) {
	// query
	const sqlstr = `SELECT ` +
		`i.indkey ` + // ::varchar AS ord
		`FROM pg_index i ` +
		`JOIN ONLY pg_class c ON c.oid = i.indrelid ` +
		`JOIN ONLY pg_namespace n ON n.oid = c.relnamespace ` +
		`JOIN ONLY pg_class ic ON ic.oid = i.indexrelid ` +
		`WHERE n.nspname = $1 ` +
		`AND ic.relname = $2`
	// run
	logf(sqlstr, schema, index)
	var pco PostgresColOrder
	if err := db.QueryRowContext(ctx, sqlstr, schema, index).Scan(&pco.Ord); err != nil {
		return nil, logerror(err)
	}
	return &pco, nil
}
