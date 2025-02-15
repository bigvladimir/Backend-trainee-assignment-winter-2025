package model

type SendCoinRequest struct {
	SenderID int
	// Имя пользователя, которому нужно отправить монеты.
	ToUser string `json:"toUser"`
	// Количество монет, которые необходимо отправить.
	Amount int `json:"amount"`
}

type SaveTransactionRequest struct {
	SenderID   int
	ReceiverID int
	Amount     int
}

type PurchaseRequest struct {
	UserID int
	Type_  string
}

type SavePurchaseRequest struct {
	UserID  int
	MerchID int
}

type UpdateBalanceRequest struct {
	UserID int
	Amount int
}
