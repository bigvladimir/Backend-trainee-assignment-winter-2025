package service

import (
	"context"

	"avito-shop-service/internal/app/model"
)

type storage interface {
	GetUserIDByUsername(ctx context.Context, username string) (int, error)
	GetUserAuthInfoByUsername(ctx context.Context, username string) (model.UserAuthInfo, error)

	GetUserCoinsByID(ctx context.Context, id int) (int, error)
	GetUserInventoryByID(ctx context.Context, id int) ([]model.InfoResponseInventory, error)
	GetUserCoinHistoryReceivedByID(ctx context.Context, id int) ([]model.InfoResponseCoinHistoryReceived, error)
	GetUserCoinHistorySentByID(ctx context.Context, id int) ([]model.InfoResponseCoinHistorySent, error)

	GetMerchIDbyName(ctx context.Context, name string) (int, error)
	GetMerchPricebyID(ctx context.Context, id int) (int, error)

	AddUser(ctx context.Context, input model.UserCreation) error
	AddCoinTransaction(ctx context.Context, input model.SaveTransactionRequest) error
	AddPurchase(ctx context.Context, input model.SavePurchaseRequest) error

	UpdateUserBalance(ctx context.Context, input model.UpdateBalanceRequest) error

	IsNotFound(err error) bool
}

type transactionManager interface {
	RunSerializable(ctx context.Context, f func(ctxTX context.Context) error) error
}

type tokenManager interface {
	CreateToken(userID int) (string, error)
}

type InputServiceSettings struct {
	userStartBalance int
}

type serviceSettings struct {
	userStartBalance int
}

type Service struct {
	stor      storage
	txManager transactionManager
	tm        tokenManager

	settings serviceSettings
}

func NewService(s storage, tx transactionManager, tm tokenManager, ss InputServiceSettings) *Service {
	return &Service{
		stor:      s,
		txManager: tx,
		tm:        tm,

		settings: serviceSettings{
			userStartBalance: ss.userStartBalance,
		},
	}
}
