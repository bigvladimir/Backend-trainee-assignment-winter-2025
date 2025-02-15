package model

type InfoResponse struct {
	Coins       int                      `json:"coins,omitempty"`
	Inventory   []InfoResponseInventory  `json:"inventory,omitempty"`
	CoinHistory *InfoResponseCoinHistory `json:"coinHistory,omitempty"`
}

type InfoResponseInventory struct {
	// Тип предмета.
	Type_ string `json:"type,omitempty"`
	// Количество предметов.
	Quantity int `json:"quantity,omitempty"`
}

type InfoResponseCoinHistory struct {
	Received []InfoResponseCoinHistoryReceived `json:"received,omitempty"`
	Sent     []InfoResponseCoinHistorySent     `json:"sent,omitempty"`
}

type InfoResponseCoinHistoryReceived struct {
	// Имя пользователя, который отправил монеты.
	FromUser string `json:"fromUser,omitempty"`
	// Количество полученных монет.
	Amount int `json:"amount,omitempty"`
}

type InfoResponseCoinHistorySent struct {
	// Имя пользователя, которому отправлены монеты.
	ToUser string `json:"toUser,omitempty"`
	// Количество отправленных монет.
	Amount int `json:"amount,omitempty"`
}
