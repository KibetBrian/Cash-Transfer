package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id   uuid.UUID `json:"userId" gorm:"primaryKey; not null;"`
	Name  string    `json:"userName" gorm:"not null;"`
	Email string    `json:"userEmail" gorm:"unique; not null" `
	Password  string    `json:"password" gorm:"not null;"`
	CreatedAt time.Time `json:"createdAt" sql:"type:timestamp" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
	AccountId uuid.UUID `json:"AccountId"`
}

type Transaction struct {
	gorm.Model
	Id uuid.UUID       `json:"transactionId"`
	Sender        uuid.UUID       `json:"senderId"`
	Receiver      uuid.UUID       `json:"receiverId"`
	Amount        decimal.Decimal `json:"amount"`
	Message       string          `json:"message"`
}

type Account struct {
	gorm.Model
	Id    uuid.UUID       `json:"accountId" gorm:"primaryKey; unique; not null" `
	Balance      decimal.Decimal `json:"accountBalance" gorm:"not null;"`
	HolderId     uuid.UUID       `json:"holderId" gorm:"not null;"`
	Transactions []*Transaction  `json:"accountTransactions" gorm:"-"`
	Password     string			 `json:"password" gorm:"not null;"`
}

type TransactionRequest struct {
	SenderEmail   string          `json:"senderEmail"`
	ReceiverEmail string          `json:"receiverEmail"`
	Amount        decimal.Decimal `json:"amount"`
	Message       string          `json:"message"`
}