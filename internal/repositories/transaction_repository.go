package repositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maritsikmaly/golang-finance-app/internal/entities"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(tx *fiber.Ctx, transaction *entities.Transaction) (*entities.Transaction, error)
	Update(tx *fiber.Ctx, transaction *entities.Transaction) error
	Delete(tx *fiber.Ctx, id string) error
	DeleteMultiple(tx *fiber.Ctx, ids []string) error
	Show(tx *fiber.Ctx, id string) (*entities.Transaction, error)
	GetByUserID(tx *fiber.Ctx, userID string) ([]*entities.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (t *transactionRepository) Create(tx *fiber.Ctx, transaction *entities.Transaction) (*entities.Transaction, error) {
	if err := t.db.Create(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t *transactionRepository) Update(tx *fiber.Ctx, transaction *entities.Transaction) error {
	if err := t.db.Save(transaction).Error; err != nil {
		return err
	}
	return nil
}

func (t *transactionRepository) Show(tx *fiber.Ctx, id string) (*entities.Transaction, error) {
	var transaction entities.Transaction

	if err := t.db.First(&transaction, id).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (t *transactionRepository) Delete(tx *fiber.Ctx, id string) error {
	if err := t.db.Delete(&entities.Transaction{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (t *transactionRepository) DeleteMultiple(tx *fiber.Ctx, ids []string) error {
	if err := t.db.Where("id IN ?", ids).Delete(&entities.Transaction{}).Error; err != nil {
		return err
	}
	return nil
}

func (t *transactionRepository) GetByUserID(tx *fiber.Ctx, userID string) ([]*entities.Transaction, error) {
	var transactions []*entities.Transaction
	if err := t.db.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
