package service

import (
	"context"

	"avito-shop-service/internal/app/model"
)

type storage interface {
	GetUserAuthInfoByUsername(ctx context.Context, username string) (model.UserAuthInfo, error)

	AddUser(ctx context.Context, input model.UserCreation) error

	IsNotFound(err error) bool
}

type transactionManager interface {
	RunSerializable(ctx context.Context, f func(ctxTX context.Context) error) error
}

type AuthManager struct {
	stor      storage
	txManager transactionManager
}

func NewAuthManager(s storage, tx transactionManager) *AuthManager {
	return &AuthManager{
		stor:      s,
		txManager: tx,
	}
}

// решить чо делать с мидлвэер
// наверное в эндпоинте создание токена и регистрация, а в сидлваре просто проверка токена где
// будет доставаться секретный ключ из конфига
