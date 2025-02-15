package repository

import (
	"avito-shop-service/internal/app/model"
)

type CoinHistoryReceived struct {
	// Имя пользователя, который отправил монеты.
	FromUser string `db:"sender_id"`
	// Количество полученных монет.
	Amount int `db:"amount"`
}

func (c *CoinHistoryReceived) MapToCoinHistoryReceivedServiceModel() model.InfoResponseCoinHistoryReceived {
	return model.InfoResponseCoinHistoryReceived{
		FromUser: c.FromUser,
		Amount:   c.Amount,
	}
}

type CoinHistorySent struct {
	// Имя пользователя, которому отправлены монеты.
	ToUser string `db:"receiver_id"`
	// Количество отправленных монет.
	Amount int `db:"amount"`
}

func (c *CoinHistorySent) MapToCoinHistorySentServiceModel() model.InfoResponseCoinHistorySent {
	return model.InfoResponseCoinHistorySent{
		ToUser: c.ToUser,
		Amount: c.Amount,
	}
}

type Purchase struct {
	UserID  int
	MerchID int
}

func (p *Purchase) MapFromPurchaseServiceModel(u model.SavePurchaseRequest) {
	p.UserID = u.UserID
	p.MerchID = u.MerchID
}

type SaveTransaction struct {
	SenderID   int
	ReceiverID int
	Amount     int
}

func (s *SaveTransaction) MapFromSaveTransactionServiceModel(d model.SaveTransactionRequest) {
	s.SenderID = d.SenderID
	s.ReceiverID = d.ReceiverID
	s.Amount = d.Amount
}
