package models

import "time"

type Expense struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	GroupID     uint           `json:"group_id" gorm:"not null;index"`
	Description string         `json:"description" gorm:"not null"`
	Amount      float64        `json:"amount" gorm:"not null"`
	Currency    string         `json:"currency" gorm:"default:TRY"`
	PaidBy      uint           `json:"paid_by" gorm:"not null"`
	SplitType   string         `json:"split_type" gorm:"default:equal"`
	Date        time.Time      `json:"date"`
	CreatedAt   time.Time      `json:"created_at"`
	Splits      []ExpenseSplit `json:"splits" gorm:"foreignKey:ExpenseID"`
}

type ExpenseSplit struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	ExpenseID uint    `json:"expense_id" gorm:"not null;index"`
	UserID    uint    `json:"user_id" gorm:"not null"`
	Amount    float64 `json:"amount" gorm:"not null"`
}
