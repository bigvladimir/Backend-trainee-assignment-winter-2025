package postgresql

import (
	"context"
	"fmt"

	"avito-shop-service/internal/app/model"
	"avito-shop-service/internal/app/repository"
	"avito-shop-service/internal/app/repository/rep_errors"
)

func (s *ServiceStorage) AddUser(ctx context.Context, input model.UserCreation) error {
	var user repository.User
	user.MapFromCreationServiceModel(input)
	_, err := s.qe.GetQueryEngine(ctx).Exec(
		ctx, "INSERT INTO user(username, password_hash, balance) VALUES ($1, $2, $3)", user.Username, user.PasswordHash, user.Balance,
	)
	if err != nil {
		return fmt.Errorf("AddUser INSERT query: %w", err)
	}
	return nil
}

func (s *ServiceStorage) UpdateUserBalance(ctx context.Context, input model.UpdateBalanceRequest) error {
	var upd repository.UpdateBalance
	upd.MapFromUpdateBalanceServiceModel(input)
	commandTag, err := s.qe.GetQueryEngine(ctx).Exec(
		ctx, "UPDATE user SET balance = $1 WHERE id = $2;",
		upd.Amount, upd.UserID,
	)
	if err != nil {
		err = fmt.Errorf("UpdateUserBalance UPDATE query: %w", err)
		return err
	}
	rowsAffected := commandTag.RowsAffected()
	if rowsAffected == 0 {
		return rep_errors.ErrNotFound
	}
	return nil
}
