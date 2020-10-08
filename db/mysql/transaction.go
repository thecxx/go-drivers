package mysql

import (
	"context"
	"database/sql"
)

type Transaction struct {
	tx *sql.Tx
}

// Query executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (t *Transaction) Query(ctx context.Context, query string, args ...interface{}) (Result, error) {
	rows, err := t.tx.QueryContext(ctx, query, args...)
	return Result{rows: rows}, err
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (t *Transaction) Exec(ctx context.Context, query string, args ...interface{}) (Result, error) {
	result, err := t.tx.ExecContext(ctx, query, args...)
	return Result{result: result}, err
}

// Rollback aborts the transaction.
func (t *Transaction) Rollback() error {
	return t.tx.Rollback()
}

// Commit commits the transaction.
func (t *Transaction) Commit() error {
	return t.tx.Commit()
}
