package usecases

import (
	"log"
	"time"

	"github.com/maritsikmaly/golang-finance-app/internal/entities"
	"github.com/maritsikmaly/golang-finance-app/internal/models"
	"github.com/maritsikmaly/golang-finance-app/internal/repositories"
)

type TransactionUseCase interface {
	Create(req *models.TransactionRequest) (*models.TransactionResponse, error)
	Update(req *models.TransactionRequest) (*models.TransactionResponse, error)
	Delete(id string) error
	GetByUserID(userID string) ([]*models.TransactionResponse, error)
}

type transactionUseCase struct {
	transactionRepo repositories.TransactionRepository
}

func NewTransactionUseCase(tr repositories.TransactionRepository) TransactionUseCase {
	return &transactionUseCase{
		transactionRepo: tr,
	}
}

func (t *transactionUseCase) Create(req *models.TransactionRequest) (*models.TransactionResponse, error) {
	dateString := req.Date
	dateFormat := "2006-01-02"

	parsedDate, err := time.Parse(dateFormat, dateString)

	if err != nil {
		return nil, err
	}

	log.Println("Flag 1")

	transaction := &entities.Transaction{
		UserID:      req.UserID,
		Description: req.Description,
		Amount:      req.Amount,
		Category:    req.Category,
		Date:        parsedDate,
	}

	log.Println("Transaction:", transaction)

	createdTransaction, err := t.transactionRepo.Create(transaction)



	if err != nil {
		return nil, err
	}

	return &models.TransactionResponse{
		ID:          createdTransaction.ID,
		UserID:      createdTransaction.UserID,
		Description: createdTransaction.Description,
		Amount:      createdTransaction.Amount,
		Category:    createdTransaction.Category,
		Date:        createdTransaction.Date,
		CreatedAt:   createdTransaction.CreatedAt,
		UpdatedAt:   createdTransaction.UpdatedAt,
	}, nil
}

func (t *transactionUseCase) Update(req *models.TransactionRequest) (*models.TransactionResponse, error) {
	dateString := req.Date
	dateFormat := "2006-01-02"

	parsedDate, err := time.Parse(dateFormat, dateString)

	if err != nil {
		return nil, err
	}

	transaction := &entities.Transaction{
		ID:          req.ID,
		UserID:      req.UserID,
		Description: req.Description,
		Amount:      req.Amount,
		Category:    req.Category,
		Date:        parsedDate,
	}

	log.Println("Transaction:", transaction)

	if err := t.transactionRepo.Update(transaction); err != nil {
		return nil, err
	}

	return &models.TransactionResponse{
		ID:          transaction.ID,
		UserID:      transaction.UserID,
		Description: transaction.Description,
		Amount:      transaction.Amount,
		Category:    transaction.Category,
		Date:        transaction.Date,
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
	}, nil
}

func (t *transactionUseCase) Delete(id string) error {
	if err := t.transactionRepo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (t *transactionUseCase) GetByUserID(userID string) ([]*models.TransactionResponse, error) {
	transactions, err := t.transactionRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	var response []*models.TransactionResponse
	for _, transaction := range transactions {
		response = append(response, &models.TransactionResponse{
			ID:          transaction.ID,
			UserID:      transaction.UserID,
			Description: transaction.Description,
			Amount:      transaction.Amount,
			Category:    transaction.Category,
			Date:        transaction.Date,
			CreatedAt:   transaction.CreatedAt,
			UpdatedAt:   transaction.UpdatedAt,
		})
	}

	return response, nil
}
