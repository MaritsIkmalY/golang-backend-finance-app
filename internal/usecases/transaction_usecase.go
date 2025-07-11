package usecases

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/maritsikmaly/golang-finance-app/internal/entities"
	"github.com/maritsikmaly/golang-finance-app/internal/models"
	"github.com/maritsikmaly/golang-finance-app/internal/models/converter"
	"github.com/maritsikmaly/golang-finance-app/internal/repositories"
)

const (
	dateFormat = "2006-01-02"
)

type TransactionUseCase interface {
	Create(ctx *fiber.Ctx, req *models.TransactionRequest) (*models.TransactionResponse, error)
	Update(ctx *fiber.Ctx, req *models.TransactionRequest, id string) (*models.TransactionResponse, error)
	Delete(ctx *fiber.Ctx, id string) error
	DeleteMultiple(ctx *fiber.Ctx, ids []uint) error
	Show(ctx *fiber.Ctx, id string) (*models.TransactionResponse, error)
	GetByUserID(ctx *fiber.Ctx, userID string) ([]*models.TransactionResponse, error)
}

type transactionUseCase struct {
	transactionRepo repositories.TransactionRepository
}

func NewTransactionUseCase(tr repositories.TransactionRepository) TransactionUseCase {
	return &transactionUseCase{
		transactionRepo: tr,
	}
}

func (t *transactionUseCase) Create(ctx *fiber.Ctx, req *models.TransactionRequest) (*models.TransactionResponse, error) {
	parsedDate, err := parseDate(req.Date)

	if err != nil {
		return nil, err
	}

	userID, err := getUserID(ctx)

	if err != nil {
		return nil, err
	}

	transaction := &entities.Transaction{
		UserID:      userID,
		Description: req.Description,
		Amount:      req.Amount,
		Category:    req.Category,
		Date:        parsedDate,
	}

	createdTransaction, err := t.transactionRepo.Create(transaction)

	if err != nil {
		return nil, err
	}

	return converter.TransactionToResponse(createdTransaction), nil
}

func (t *transactionUseCase) Update(ctx *fiber.Ctx, req *models.TransactionRequest, id string) (*models.TransactionResponse, error) {
	parsedDate, err := parseDate(req.Date)

	if err != nil {
		return nil, err
	}

	userID, err := getUserID(ctx)

	if err != nil {
		return nil, err
	}

	transaction, err := t.transactionRepo.Show(id)

	if err != nil {
		return nil, err
	}

	err = checkOwnership(userID, transaction)

	if err != nil {
		return nil, err
	}

	transaction.Description = req.Description
	transaction.Amount = req.Amount
	transaction.Category = req.Category
	transaction.Date = parsedDate

	if err := t.transactionRepo.Update(transaction); err != nil {
		return nil, err
	}

	return converter.TransactionToResponse(transaction), nil
}

func (t *transactionUseCase) Show(ctx *fiber.Ctx, id string) (*models.TransactionResponse, error) {
	transaction, err := t.transactionRepo.Show(id)

	if err != nil {
		return nil, err
	}

	userID, err := getUserID(ctx)

	if err != nil {
		return nil, err
	}

	err = checkOwnership(userID, transaction)

	if err != nil {
		return nil, err
	}

	return converter.TransactionToResponse(transaction), nil
}

func (t *transactionUseCase) Delete(ctx *fiber.Ctx, id string) error {
	userID, err := getUserID(ctx)

	if err != nil {
		return err
	}

	transaction, err := t.transactionRepo.Show(id)

	if err != nil {
		return err
	}

	err = checkOwnership(userID, transaction)

	if err != nil {
		return err
	}

	if err := t.transactionRepo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (t *transactionUseCase) GetByUserID(ctx *fiber.Ctx, userID string) ([]*models.TransactionResponse, error) {
	transactions, err := t.transactionRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	var response []*models.TransactionResponse
	for _, transaction := range transactions {
		response = append(response, converter.TransactionToResponse(transaction))
	}

	return response, nil
}

func (t *transactionUseCase) DeleteMultiple(ctx *fiber.Ctx, ids []uint) error {
	userID, err := getUserID(ctx)
	if err != nil {
		return err
	}

	idStrings := make([]string, len(ids))
	for i, id := range ids {
		idStrings[i] = strconv.FormatUint(uint64(id), 10)
	}

	transactions, err := t.transactionRepo.GetByIDs(idStrings)
	if err != nil {
		return err
	}

	if len(transactions) != len(idStrings) {
		return fiber.NewError(fiber.StatusNotFound, "some transactions not found")
	}

	for _, tx := range transactions {
		err = checkOwnership(userID, tx)
		if err != nil {
			return err
		}
	}

	if err := t.transactionRepo.DeleteMultiple(idStrings); err != nil {
		return err
	}

	return nil
}

func parseDate(dateString string) (time.Time, error) {
	parsedDate, err := time.Parse(dateFormat, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return parsedDate, nil
}

func getUserID(ctx *fiber.Ctx) (uint, error) {
	userID := ctx.Locals("user_id")
	if userID == nil {
		return 0, fiber.NewError(fiber.StatusUnauthorized, "user_id not found in context")
	}

	parsedUserID, ok := userID.(float64)
	if !ok {
		return 0, fiber.NewError(fiber.StatusInternalServerError, "invalid user_id type")
	}

	return uint(parsedUserID), nil
}

func checkOwnership(userID uint, transaction *entities.Transaction) error {
	if transaction.UserID != userID {
		return fiber.NewError(fiber.StatusForbidden, "you do not have permission to access this transaction")
	}

	return nil
}
