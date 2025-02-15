package postgresql

import (
	"context"
	"errors"
	"fmt"

	"avito-shop-service/internal/app/model"
	"avito-shop-service/internal/app/repository"
	"avito-shop-service/internal/app/repository/rep_errors"

	"github.com/jackc/pgx/v4"
)

func (s *ServiceStorage) GetUserAuthInfoByUsername(ctx context.Context, username string) (model.UserAuthInfo, error) {
	var user repository.User
	err := s.qe.GetQueryEngine(ctx).Get(ctx, &user, "SELECT password_hash FROM user WHERE username=$1", username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.UserAuthInfo{}, rep_errors.ErrNotFound
		}
		return model.UserAuthInfo{}, fmt.Errorf("GetUserAuthInfoByUsername SELECT query: %w", err)
	}
	return user.MapToAuthInfoServiceModel(), nil
}

func (s *ServiceStorage) GetUserIDByUsername(ctx context.Context, username string) (int, error) {
	var id int
	err := s.qe.GetQueryEngine(ctx).Get(ctx, &id, "SELECT id FROM user WHERE username=$1", username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, rep_errors.ErrNotFound
		}
		return 0, fmt.Errorf("GetUserIDByUsername SELECT query: %w", err)
	}
	return id, nil
}

func (s *ServiceStorage) GetUserCoinsByID(ctx context.Context, id int) (int, error) {
	var coins int
	err := s.qe.GetQueryEngine(ctx).Get(ctx, &coins, "SELECT balance FROM user WHERE id=$1", id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, rep_errors.ErrNotFound
		}
		return 0, fmt.Errorf("GetUserCoinsByID SELECT query: %w", err)
	}
	return coins, nil
}

func (s *ServiceStorage) GetUserInventoryByID(ctx context.Context, id int) ([]model.InfoResponseInventory, error) {
	var inv []repository.InfoInventory
	err := s.qe.GetQueryEngine(ctx).Select(ctx, &inv,
		`SELECT merch.name, COUNT(*) AS quantity
		FROM purchase
		JOIN merch ON purchase.merch_id = merch.id
		WHERE purchase.user_id=$1
		GROUP BY merch.name`, id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("GetUserInventoryByID SELECT query: %w", err)
	}
	resInv := make([]model.InfoResponseInventory, len(inv))
	for i, v := range inv {
		resInv[i] = v.MapToInfoInventoryServiceModel()
	}
	return resInv, nil
}

func (s *ServiceStorage) GetUserCoinHistoryReceivedByID(ctx context.Context, id int) ([]model.InfoResponseCoinHistoryReceived, error) {
	var rec []repository.CoinHistoryReceived
	err := s.qe.GetQueryEngine(ctx).Select(ctx, &rec, "SELECT sender_id, amount FROM coins_transaction WHERE receiver_id=$1", id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("GetUserCoinHistoryReceivedByID SELECT query: %w", err)
	}
	resRec := make([]model.InfoResponseCoinHistoryReceived, len(rec))
	for i, v := range rec {
		resRec[i] = v.MapToCoinHistoryReceivedServiceModel()
	}
	return resRec, nil
}

func (s *ServiceStorage) GetUserCoinHistorySentByID(ctx context.Context, id int) ([]model.InfoResponseCoinHistorySent, error) {
	var sent []repository.CoinHistorySent
	err := s.qe.GetQueryEngine(ctx).Select(ctx, &sent, "SELECT receiver_id, amount FROM coins_transaction WHERE sender_id=$1", id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("GetUserCoinHistorySentByID SELECT query: %w", err)
	}
	resSent := make([]model.InfoResponseCoinHistorySent, len(sent))
	for i, v := range sent {
		resSent[i] = v.MapToCoinHistorySentServiceModel()
	}
	return resSent, nil
}
