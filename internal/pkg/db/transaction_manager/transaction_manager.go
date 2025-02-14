package transaction_manager

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"go.uber.org/multierr"
)

var transactionKey = transactionKeyType{}

type transactionKeyType struct{}

type DbOps interface {
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type dbOpsService interface {
	DbOps
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

type TransactionManager struct {
	pool dbOpsService
}

func NewTransactionManager(pool dbOpsService) *TransactionManager {
	return &TransactionManager{
		pool: pool,
	}
}

// RunSerializable запускает Serializable транзакцию
func (t *TransactionManager) RunSerializable(ctx context.Context, f func(ctxTX context.Context) error) error {
	tx, err := t.pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.Serializable,
		AccessMode: pgx.ReadWrite,
	})
	if err != nil {
		return fmt.Errorf("pool.BeginTx: %w", err)
	}

	if err = f(context.WithValue(ctx, transactionKey, newTransaction(tx))); err != nil {
		return multierr.Combine(err, tx.Rollback(ctx))
	}

	if err := tx.Commit(ctx); err != nil {
		return multierr.Combine(err, tx.Rollback(ctx))
	}

	return nil
}

// GetQueryEngine повзволяет вызвать запросы к бд внутри транзакции,
// в противном случае позволяет обратиться к бд вне транзакции
func (t *TransactionManager) GetQueryEngine(ctx context.Context) DbOps {
	if tx, ok := ctx.Value(transactionKey).(DbOps); ok {
		return tx
	}
	return t.pool
}
