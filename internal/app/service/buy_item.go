package service

import (
	"context"

	"avito-shop-service/internal/app/model"
	"avito-shop-service/internal/app/service/service_errors"

	"go.uber.org/multierr"
)

func (s *Service) BuyItem(ctx context.Context, input model.PurchaseRequest) error {

	if input.UserID <= 0 {
		return service_errors.ErrInvalidReq
	}

	if err := s.txManager.RunSerializable(ctx,
		func(ctxTX context.Context) error {
			var err error
			var merchID int
			if merchID, err = s.stor.GetMerchIDbyName(ctxTX, input.Type_); err != nil {
				return err
			}
			var coins int
			if coins, err = s.stor.GetUserCoinsByID(ctxTX, input.UserID); err != nil {
				return err
			}
			var price int
			if price, err = s.stor.GetMerchPricebyID(ctxTX, merchID); err != nil {
				return err
			}
			if coins < price {
				return service_errors.ErrInvalidReq
			}

			if err = s.stor.UpdateUserBalance(
				ctxTX, model.UpdateBalanceRequest{UserID: input.UserID, Amount: (coins - price)},
			); err != nil {
				return err
			}

			if err = s.stor.AddPurchase(
				ctxTX, model.SavePurchaseRequest{UserID: input.UserID, MerchID: merchID},
			); err != nil {
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
