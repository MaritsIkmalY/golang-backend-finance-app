package repositories

import (
	"github.com/maritsikmaly/golang-finance-app/internal/entities"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *entities.Transaction) (*entities.Transaction, error)
	Update(transaction *entities.Transaction) error
	Delete(id string) error
	GetByUserID(userID string) ([]*entities.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}


func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (t *transactionRepository) Create(transaction *entities.Transaction) (*entities.Transaction, error) {
	if err := t.db.Create(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t *transactionRepository) Delete(id string) error {
	if err := t.db.Delete(&entities.Transaction{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (t *transactionRepository) GetByUserID(userID string) ([]*entities.Transaction, error) {
	var transactions []*entities.Transaction
	if err := t.db.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (t *transactionRepository) Update(transaction *entities.Transaction) error {
	if err := t.db.Save(transaction).Error; err != nil {
		return err
	}
	return nil
}