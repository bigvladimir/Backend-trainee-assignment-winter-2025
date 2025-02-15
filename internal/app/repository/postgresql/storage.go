package postgresql

import (
	"context"

	"avito-shop-service/internal/pkg/db/transaction_manager"
)

type queryEngineManager interface {
	GetQueryEngine(ctx context.Context) transaction_manager.DbOps
}

type ServiceStorage struct {
	qe queryEngineManager
}

func NewServiceStorage(queryEngine queryEngineManager) *ServiceStorage {
	return &ServiceStorage{qe: queryEngine}
}
