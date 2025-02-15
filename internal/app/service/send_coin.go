package service

import (
	"context"

	"avito-shop-service/internal/app/model"
	"avito-shop-service/internal/app/service/service_errors"

	"go.uber.org/multierr"
)

func (s *Service) SendCoin(ctx context.Context, input model.SendCoinRequest) error {

	if input.SenderID <= 0 || input.ToUser == "" || input.Amount <= 0 {
		return service_errors.ErrInvalidReq
	}

	if err := s.txManager.RunSerializable(ctx,
		func(ctxTX context.Context) error {
			var tr model.SaveTransactionRequest
			var err error
			if tr.ReceiverID, err = s.stor.GetUserIDByUsername(ctx, input.ToUser); err != nil {
				return err
			}

			var senderCoins int
			if senderCoins, err = s.stor.GetUserCoinsByID(ctx, input.SenderID); err != nil {
				return err
			}
			if senderCoins < input.Amount {
				return service_errors.ErrInvalidReq
			}

			var receiverCoins int
			if receiverCoins, err = s.stor.GetUserCoinsByID(ctx, tr.ReceiverID); err != nil {
				return err
			}
			senderCoins -= input.Amount
			receiverCoins += input.Amount
			if err = s.stor.UpdateUserBalance(
				ctx, model.UpdateBalanceRequest{UserID: input.SenderID, Amount: senderCoins},
			); err != nil {
				return err
			}
			if err = s.stor.UpdateUserBalance(
				ctx, model.UpdateBalanceRequest{UserID: tr.ReceiverID, Amount: receiverCoins},
			); err != nil {
				return err
			}

			tr.SenderID = input.SenderID
			tr.Amount = input.Amount
			if err = s.stor.AddCoinTransaction(ctx, tr); err != nil {
				return err
			}

			return nil
		},
	); err != nil {
		for _, e := range multierr.Errors(err) {
			if s.stor.IsNotFound(e) || s.IsInvalidReq(e) {
				return service_errors.ErrInvalidReq
			}
		}
		return err
	}

	return nil
}
