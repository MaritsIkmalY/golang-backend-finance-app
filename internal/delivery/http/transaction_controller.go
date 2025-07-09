package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/maritsikmaly/golang-finance-app/internal/models"
	"github.com/maritsikmaly/golang-finance-app/internal/usecases"
)

type TransactionController struct {
	transactionUsecase usecases.TransactionUseCase
	validator          *validator.Validate
}

func NewTransactionController(transactionUsecase usecases.TransactionUseCase, validator *validator.Validate) *TransactionController {
	return &TransactionController{
		transactionUsecase: transactionUsecase,
		validator:          validator,
	}
}

func (c *TransactionController) Create(ctx *fiber.Ctx) error {
	var req models.TransactionRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.validator.Struct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := c.transactionUsecase.Create(&req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(res)
}

func (c *TransactionController) Update(ctx *fiber.Ctx) error {
	var req models.TransactionRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.validator.Struct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := c.transactionUsecase.Update(&req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (c *TransactionController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Transaction ID is required"})
	}

	err := c.transactionUsecase.Delete(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

func (c *TransactionController) GetByUserID(ctx *fiber.Ctx) error {
	userID := ctx.Params("user_id")
	if userID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User ID is required"})
	}

	transactions, err := c.transactionUsecase.GetByUserID(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(transactions)
}
