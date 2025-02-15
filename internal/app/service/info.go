package service

import (
	"context"

	"avito-shop-service/internal/app/model"
	"avito-shop-service/internal/app/service/service_errors"
)

func (s *Service) Info(ctx context.Context, id int) (model.InfoResponse, error) {

	if id <= 0 {
		return model.InfoResponse{}, service_errors.ErrInvalidReq
	}

	var err error
	var info model.InfoResponse
	if info.Coins, err = s.stor.GetUserCoinsByID(ctx, id); err != nil {
		if s.stor.IsNotFound(err) {
			return model.InfoResponse{}, service_errors.ErrInvalidReq
		}
		return model.InfoResponse{}, err
	}
	if info.Inventory, err = s.stor.GetUserInventoryByID(ctx, id); err != nil {
		return model.InfoResponse{}, err
	}
	var received []model.InfoResponseCoinHistoryReceived
	if received, err = s.stor.GetUserCoinHistoryReceivedByID(ctx, id); err != nil {
		return model.InfoResponse{}, err
	}
	var sent []model.InfoResponseCoinHistorySent
	if sent, err = s.stor.GetUserCoinHistorySentByID(ctx, id); err != nil {
		return model.InfoResponse{}, err
	}
	info.CoinHistory = &model.InfoResponseCoinHistory{Received: received, Sent: sent}

	return info, nil
}
