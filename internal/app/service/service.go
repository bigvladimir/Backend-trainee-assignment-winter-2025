package service

import (
	"context"

	"avito-shop-service/internal/app/model"
)

type storage interface {
	GetUserIDByUsername(ctx context.Context, username string) (int, error)
	GetUserCoinsByID(ctx context.Context, id int) (int, error)
	GetUserInventoryByID(ctx context.Context, id int) ([]model.InfoResponseInventory, error)
	GetUserCoinHistoryReceivedByID(ctx context.Context, id int) ([]model.InfoResponseCoinHistoryReceived, error)
	GetUserCoinHistorySentByID(ctx context.Context, id int) ([]model.InfoResponseCoinHistorySent, error)

	GetMerchIDbyName(ctx context.Context, name string) (int, error)
	GetMerchPricebyID(ctx context.Context, id int) (int, error)

	AddCoinTransaction(ctx context.Context, input model.SaveTransactionRequest) error
	AddPurchase(ctx context.Context, input model.SavePurchaseRequest) error

	UpdateUserBalance(ctx context.Context, input model.UpdateBalanceRequest) error

	IsNotFound(err error) bool
}

type transactionManager interface {
	RunSerializable(ctx context.Context, f func(ctxTX context.Context) error) error
}

type Service struct {
	stor      storage
	txManager transactionManager
}

func NewService(s storage, tx transactionManager) *Service {
	return &Service{
		stor:      s,
		txManager: tx,
	}
}
