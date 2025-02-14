package transaction_manager

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type transaction struct {
	tx pgx.Tx
}

func newTransaction(tx pgx.Tx) *transaction {
	return &transaction{tx: tx}
}

func (t transaction) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, t.tx, dest, query, args...)
}

func (t transaction) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, t.tx, dest, query, args...)
}

func (t transaction) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return t.tx.Exec(ctx, query, args...)
}

func (t transaction) ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return t.tx.QueryRow(ctx, query, args...)
}
