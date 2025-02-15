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

func (s *ServiceStorage) GetMerchIDbyName(ctx context.Context, name string) (int, error) {
	var id int
	err := s.qe.GetQueryEngine(ctx).Get(ctx, &id, "SELECT id FROM merch WHERE name=$1", name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, rep_errors.ErrNotFound
		}
		return 0, fmt.Errorf("GetMerchIDbyName SELECT query: %w", err)
	}
	return id, nil
}

func (s *ServiceStorage) GetMerchPricebyID(ctx context.Context, id int) (int, error) {
	var price int
	err := s.qe.GetQueryEngine(ctx).Get(ctx, &price, "SELECT price FROM merch WHERE id=$1", id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, rep_errors.ErrNotFound
		}
		return 0, fmt.Errorf("GetMerchPricebyID SELECT query: %w", err)
	}
	return price, nil
}

func (s *ServiceStorage) AddPurchase(ctx context.Context, input model.SavePurchaseRequest) error {
	var pur repository.Purchase
	pur.MapFromPurchaseServiceModel(input)
	_, err := s.qe.GetQueryEngine(ctx).Exec(ctx, "INSERT INTO purchase(user_id, merch_id) VALUES ($1, $2)", pur.UserID, pur.MerchID)
	if err != nil {
		return fmt.Errorf("AddPurchase INSERT query: %w", err)
	}
	return nil
}

func (s *ServiceStorage) AddCoinTransaction(ctx context.Context, input model.SaveTransactionRequest) error {
	var tr repository.SaveTransaction
	tr.MapFromSaveTransactionServiceModel(input)
	_, err := s.qe.GetQueryEngine(ctx).Exec(
		ctx, "INSERT INTO coins_transaction(sender_id, receiver_id, amount) VALUES ($1, $2, $3)", tr.SenderID, tr.ReceiverID, tr.Amount,
	)
	if err != nil {
		return fmt.Errorf("AddUser INSERT query: %w", err)
	}
	return nil
}
