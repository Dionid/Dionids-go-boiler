package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// ProcParam is a stored procedure param.
type ProcParam struct {
	ParamName string `json:"param_name"` // param_name
	ParamType string `json:"param_type"` // param_type
}

// PostgresProcParams runs a custom query, returning results as [ProcParam].
func PostgresProcParams(ctx context.Context, db DB, schema, id string) ([]*ProcParam, error) {
	// query
	const sqlstr = `SELECT ` +
		`COALESCE(pp.param_name, ''), ` + // ::varchar AS param_name
		`pp.param_type ` + // ::varchar AS param_type
		`FROM pg_proc p ` +
		`JOIN ONLY pg_namespace n ON p.pronamespace = n.oid ` +
		`JOIN ( ` +
		`SELECT ` +
		`p.oid, ` +
		`UNNEST(p.proargnames) AS param_name, ` +
		`format_type(UNNEST(p.proargtypes), NULL) AS param_type ` +
		`FROM pg_proc p ` +
		`) AS pp ON p.oid = pp.oid ` +
		`WHERE n.nspname = $1 ` +
		`AND p.oid::varchar = $2 ` +
		`AND pp.param_type IS NOT NULL`
	// run
	logf(sqlstr, schema, id)
	rows, err := db.QueryContext(ctx, sqlstr, schema, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*ProcParam
	for rows.Next() {
		var pp ProcParam
		// scan
		if err := rows.Scan(&pp.ParamName, &pp.ParamType); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &pp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// MysqlProcParams runs a custom query, returning results as [ProcParam].
func MysqlProcParams(ctx context.Context, db DB, schema, id string) ([]*ProcParam, error) {
	// query
	const sqlstr = `SELECT ` +
		`p.parameter_name AS param_name, ` +
		`p.dtd_identifier AS param_type ` +
		`FROM information_schema.routines r ` +
		`INNER JOIN information_schema.parameters p ON p.specific_schema = r.routine_schema ` +
		`AND p.specific_name = r.routine_name ` +
		`AND p.parameter_mode = 'IN' ` +
		`WHERE r.routine_schema = ? ` +
		`AND r.routine_name = ? ` +
		`ORDER BY p.ordinal_position`
	// run
	logf(sqlstr, schema, id)
	rows, err := db.QueryContext(ctx, sqlstr, schema, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*ProcParam
	for rows.Next() {
		var pp ProcParam
		// scan
		if err := rows.Scan(&pp.ParamName, &pp.ParamType); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &pp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// SqlserverProcParams runs a custom query, returning results as [ProcParam].
func SqlserverProcParams(ctx context.Context, db DB, schema, id string) ([]*ProcParam, error) {
	// query
	const sqlstr = `SELECT ` +
		`SUBSTRING(p.name, 2, LEN(p.name)-1) AS param_name, ` +
		`TYPE_NAME(p.user_type_id) AS param_type ` +
		`FROM sys.objects o ` +
		`INNER JOIN sys.parameters p ON o.object_id = p.object_id ` +
		`WHERE SCHEMA_NAME(schema_id) = @p1 ` +
		`AND STR(o.object_id) = @p2 ` +
		`AND p.is_output = 'false' ` +
		`ORDER BY p.parameter_id`
	// run
	logf(sqlstr, schema, id)
	rows, err := db.QueryContext(ctx, sqlstr, schema, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*ProcParam
	for rows.Next() {
		var pp ProcParam
		// scan
		if err := rows.Scan(&pp.ParamName, &pp.ParamType); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &pp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// OracleProcParams runs a custom query, returning results as [ProcParam].
func OracleProcParams(ctx context.Context, db DB, schema, id string) ([]*ProcParam, error) {
	// query
	const sqlstr = `SELECT ` +
		`LOWER(a.argument_name) AS param_name, ` +
		`LOWER(CASE a.data_type ` +
		`WHEN 'CHAR' THEN 'CHAR(' || a.data_length || ')' ` +
		`WHEN 'VARCHAR2' THEN 'VARCHAR2(' || a.data_length || ')' ` +
		`WHEN 'NUMBER' THEN 'NUMBER(' || NVL(a.data_precision, 0) || ',' || NVL(a.data_scale, 0) || ')' ` +
		`ELSE a.data_type END) AS param_type ` +
		`FROM all_objects o ` +
		`JOIN sys.all_arguments a ON a.object_id = o.object_id ` +
		`AND a.in_out = 'IN' ` +
		`WHERE o.object_type IN ('FUNCTION','PROCEDURE') ` +
		`AND o.owner = UPPER(:1) ` +
		`AND CAST(o.object_id AS NVARCHAR2(255)) = UPPER(:2) ` +
		`ORDER BY a.position`
	// run
	logf(sqlstr, schema, id)
	rows, err := db.QueryContext(ctx, sqlstr, schema, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*ProcParam
	for rows.Next() {
		var pp ProcParam
		// scan
		if err := rows.Scan(&pp.ParamName, &pp.ParamType); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &pp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}