package converter

import (
	"github.com/maritsikmaly/golang-finance-app/internal/entities"
	"github.com/maritsikmaly/golang-finance-app/internal/models"
)

const (
	dateFormat = "2006-01-02"
)

func TransactionToResponse(transaction *entities.Transaction) *models.TransactionResponse {
	if transaction == nil {
		return nil
	}

	return &models.TransactionResponse{
		ID:          transaction.ID,
		UserID:      transaction.UserID,
		Description: transaction.Description,
		Amount:      transaction.Amount,
		Category:    transaction.Category,
		Date:        transaction.Date.Format(dateFormat),
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
	}
}