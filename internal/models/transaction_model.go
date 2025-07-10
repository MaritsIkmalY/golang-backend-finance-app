package models

import "time"

type TransactionResponse struct {
	ID          uint      `json:"id"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Date        string    `json:"date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      uint      `json:"user_id"`
}

type TransactionRequest struct {
	Amount      float64 `json:"amount" validate:"required"`
	Description string  `json:"description"`
	Category    string  `json:"category" validate:"required,category_valid"`
	Date        string  `json:"date" validate:"required"`
}

type DeleteMultipleRequest struct {
	IDs []uint `json:"ids" validate:"required"`
}
