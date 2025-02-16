package repository

import (
	"avito-shop-service/internal/app/model"
)

type User struct {
	ID           int    `db:"id"`
	Username     string `db:"username"`
	PasswordHash []byte `db:"password_hash"`
	Balance      int    `db:"balance"`
}

func (u *User) MapFromCreationServiceModel(d model.UserCreation) {
	u.Username = d.Username
	u.PasswordHash = d.PasswordHash
	u.Balance = d.Balance
}

func (u *User) MapToAuthInfoServiceModel() model.UserAuthInfo {
	return model.UserAuthInfo{
		UserID:       u.ID,
		PasswordHash: u.PasswordHash,
	}
}

type InfoInventory struct {
	// Тип предмета.
	Type_ string `db:"name"`
	// Количество предметов.
	Quantity int `db:"quantity"`
}

func (i *InfoInventory) MapToInfoInventoryServiceModel() model.InfoResponseInventory {
	return model.InfoResponseInventory{
		Type_:    i.Type_,
		Quantity: i.Quantity,
	}
}

type UpdateBalance struct {
	UserID int
	Amount int
}

func (u *UpdateBalance) MapFromUpdateBalanceServiceModel(b model.UpdateBalanceRequest) {
	u.UserID = b.UserID
	u.Amount = b.Amount
}
