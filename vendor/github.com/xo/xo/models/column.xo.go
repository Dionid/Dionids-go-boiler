package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// Column is a column.
type Column struct {
	FieldOrdinal int            `json:"field_ordinal"`  // field_ordinal
	ColumnName   string         `json:"column_name"`    // column_name
	DataType     string         `json:"data_type"`      // data_type
	NotNull      bool           `json:"not_null"`       // not_null
	DefaultValue sql.NullString `json:"default_value"`  // default_value
	IsPrimaryKey bool           `json:"is_primary_key"` // is_primary_key
	Comment      sql.NullString `json:"comment"`        // comment
}

// PostgresTableColumns runs a custom query, returning results as [Column].
func PostgresTableColumns(ctx context.Context, db DB, schema, table string, sys bool) ([]*Column, error) {
	// query
	const sqlstr = `SELECT ` +
		`a.attnum, ` + // ::integer AS field_ordinal
		`a.attname, ` + // ::varchar AS column_name
		`format_type(a.atttypid, a.atttypmod), ` + // ::varchar AS data_type
		`a.attnotnull, ` + // ::boolean AS not_null
		`COALESCE(pg_get_expr(ad.adbin, ad.adrelid), ''), ` + // ::varchar AS default_value
		`COALESCE(ct.contype = 'p', false), ` + // ::boolean AS is_primary_key
		`d.description ` + // ::varchar as comment
		`FROM pg_attribute a ` +
		`JOIN ONLY pg_class c ON c.oid = a.attrelid ` +
		`JOIN ONLY pg_namespace n ON n.oid = c.relnamespace ` +
		`LEFT JOIN pg_constraint ct ON ct.conrelid = c.oid ` +
		`AND a.attnum = ANY(ct.conkey) ` +
		`AND ct.contype = 'p' ` +
		`LEFT JOIN pg_attrdef ad ON ad.adrelid = c.oid ` +
		`AND ad.adnum = a.attnum ` +
		`LEFT JOIN pg_description d on d.objoid = c.oid ` +
		`AND d.objsubid = a.attnum ` +
		`WHERE a.attisdropped = false ` +
		`AND n.nspname = $1 ` +
		`AND c.relname = $2 ` +
		`AND ($3 OR a.attnum > 0) ` +
		`ORDER BY a.attnum`
	// run
	logf(sqlstr, schema, table, sys)
	rows, err := db.QueryContext(ctx, sqlstr, schema, table, sys)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Column
	for rows.Next() {
		var c Column
		// scan
		if err := rows.Scan(&c.FieldOrdinal, &c.ColumnName, &c.DataType, &c.NotNull, &c.DefaultValue, &c.IsPrimaryKey, &c.Comment); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &c)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// MysqlTableColumns runs a custom query, returning results as [Column].
func MysqlTableColumns(ctx context.Context, db DB, schema, table string) ([]*Column, error) {
	// query
	const sqlstr = `SELECT ` +
		`ordinal_position AS field_ordinal, ` +
		`column_name, ` +
		`IF(data_type = 'enum', column_name, column_type) AS data_type, ` +
		`IF(is_nullable = 'YES', false, true) AS not_null, ` +
		`column_default AS default_value, ` +
		`IF(column_key = 'PRI', true, false) AS is_primary_key, ` +
		`column_comment AS comment ` +
		`FROM information_schema.columns ` +
		`WHERE table_schema = ? ` +
		`AND table_name = ? ` +
		`ORDER BY ordinal_position`
	// run
	logf(sqlstr, schema, table)
	rows, err := db.QueryContext(ctx, sqlstr, schema, table)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Column
	for rows.Next() {
		var c Column
		// scan
		if err := rows.Scan(&c.FieldOrdinal, &c.ColumnName, &c.DataType, &c.NotNull, &c.DefaultValue, &c.IsPrimaryKey, &c.Comment); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &c)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// Sqlite3TableColumns runs a custom query, returning results as [Column].
func Sqlite3TableColumns(ctx context.Context, db DB, schema, table string) ([]*Column, error) {
	// query
	sqlstr := `/* ` + schema + ` */ ` +
		`SELECT ` +
		`cid AS field_ordinal, ` +
		`name AS column_name, ` +
		`type AS data_type, ` +
		`"notnull" AS not_null, ` +
		`dflt_value AS default_value, ` +
		`CAST(pk <> 0 AS boolean) AS is_primary_key ` +
		`FROM pragma_table_info($1)`
	// run
	logf(sqlstr, table)
	rows, err := db.QueryContext(ctx, sqlstr, table)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Column
	for rows.Next() {
		var c Column
		// scan
		if err := rows.Scan(&c.FieldOrdinal, &c.ColumnName, &c.DataType, &c.NotNull, &c.DefaultValue, &c.IsPrimaryKey); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &c)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// SqlserverTableColumns runs a custom query, returning results as [Column].
func SqlserverTableColumns(ctx context.Context, db DB, schema, table string) ([]*Column, error) {
	// query
	const sqlstr = `SELECT ` +
		`c.colid AS field_ordinal, ` +
		`c.name AS column_name, ` +
		`TYPE_NAME(c.xtype)+IIF(c.prec > 0, '('+CAST(c.prec AS varchar)+IIF(c.scale > 0,','+CAST(c.scale AS varchar),'')+')', '') AS data_type, ` +
		`IIF(c.isnullable=1, 0, 1) AS not_null, ` +
		`x.text AS default_value, ` +
		`IIF(COALESCE(( ` +
		`SELECT COUNT(z.colid) ` +
		`FROM sysindexes i ` +
		`INNER JOIN sysindexkeys z ON i.id = z.id ` +
		`AND i.indid = z.indid ` +
		`AND z.colid = c.colid ` +
		`WHERE i.id = o.id ` +
		`AND i.name = k.name ` +
		`), 0) > 0, 1, 0) AS is_primary_key ` +
		`FROM syscolumns c ` +
		`JOIN sysobjects o ON o.id = c.id ` +
		`LEFT JOIN sysobjects k ON k.xtype = 'PK' ` +
		`AND k.parent_obj = o.id ` +
		`LEFT JOIN syscomments x ON x.id = c.cdefault ` +
		`WHERE o.type IN('U', 'V') ` +
		`AND SCHEMA_NAME(o.uid) = @p1 ` +
		`AND o.name = @p2 ` +
		`ORDER BY c.colid`
	// run
	logf(sqlstr, schema, table)
	rows, err := db.QueryContext(ctx, sqlstr, schema, table)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Column
	for rows.Next() {
		var c Column
		// scan
		if err := rows.Scan(&c.FieldOrdinal, &c.ColumnName, &c.DataType, &c.NotNull, &c.DefaultValue, &c.IsPrimaryKey); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &c)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// OracleTableColumns runs a custom query, returning results as [Column].
func OracleTableColumns(ctx context.Context, db DB, schema, table string) ([]*Column, error) {
	// query
	const sqlstr = `SELECT ` +
		`c.column_id AS field_ordinal, ` +
		`LOWER(c.column_name) AS column_name, ` +
		`LOWER(CASE c.data_type ` +
		`WHEN 'CHAR' THEN 'CHAR(' || c.char_length || ')' ` +
		`WHEN 'NCHAR' THEN 'NCHAR(' || c.char_length || ')' ` +
		`WHEN 'VARCHAR2' THEN 'VARCHAR2(' || c.char_length || ')' ` +
		`WHEN 'NVARCHAR2' THEN 'NVARCHAR2(' || c.char_length || ')' ` +
		`WHEN 'NUMBER' THEN 'NUMBER(' || NVL(c.data_precision, 0) || ',' || NVL(c.data_scale, 0) || ')' ` +
		`WHEN 'RAW' THEN 'RAW(' || c.data_length || ')' ` +
		`ELSE c.data_type END) AS data_type, ` +
		`CASE WHEN c.nullable = 'N' THEN '1' ELSE '0' END AS not_null, ` +
		`CASE WHEN p.column_id IS NOT NULL THEN '1' ELSE '0' END as is_primary_key ` +
		`FROM all_tab_columns c ` +
		`LEFT JOIN ( ` +
		`SELECT distinct c.column_id FROM all_tab_columns c ` +
		`JOIN all_cons_columns l ON l.owner = c.owner ` +
		`AND c.column_name = l.column_name ` +
		`JOIN all_constraints r ON r.owner = c.owner ` +
		`AND r.table_name = c.table_name ` +
		`AND r.constraint_name = l.constraint_name ` +
		`AND r.constraint_type = 'P' ` +
		`WHERE c.owner = UPPER(:1) ` +
		`AND c.table_name = UPPER(:2) ` +
		`) p on p.column_id = c.column_id ` +
		`WHERE c.owner = UPPER(:1) ` +
		`AND c.table_name = UPPER(:2) ` +
		`ORDER BY c.column_id`
	// run
	logf(sqlstr, schema, table)
	rows, err := db.QueryContext(ctx, sqlstr, schema, table)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Column
	for rows.Next() {
		var c Column
		// scan
		if err := rows.Scan(&c.FieldOrdinal, &c.ColumnName, &c.DataType, &c.NotNull, &c.IsPrimaryKey); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &c)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}