package models

import "time"

type TransactionResponse struct {
	ID          uint      `json:"id"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      uint      `json:"user_id"`
}

type TransactionRequest struct {
	ID          uint    `json:"id,omitempty"`
	Amount      float64 `json:"amount" validate:"required"`
	Description string  `json:"description"`
	Category    string  `json:"category" validate:"required,category_valid"`
	UserID      uint    `json:"user_id" validate:"required"`
	Date        string  `json:"date" validate:"required"`
}
