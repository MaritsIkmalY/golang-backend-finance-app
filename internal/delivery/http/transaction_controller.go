package http

import (
	"fmt"

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

	res, err := c.transactionUsecase.Create(ctx, &req)
	
	if err != nil {
		if fe, ok := err.(*fiber.Error); ok {
			return ctx.Status(fe.Code).JSON(fiber.Map{
				"error": fe.Message,
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(res)
}

func (c *TransactionController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var req models.TransactionRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.validator.Struct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := c.transactionUsecase.Update(ctx, &req, id)

	if err != nil {
		if fe, ok := err.(*fiber.Error); ok {
			return ctx.Status(fe.Code).JSON(fiber.Map{
				"error": fe.Message,
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (c *TransactionController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Transaction ID is required"})
	}

	err := c.transactionUsecase.Delete(ctx, id)

	if err != nil {
		if fe, ok := err.(*fiber.Error); ok {
			return ctx.Status(fe.Code).JSON(fiber.Map{
				"error": fe.Message,
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

func (c *TransactionController) GetByUserID(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id")

	transactions, err := c.transactionUsecase.GetByUserID(ctx, fmt.Sprintf("%.0f", userID))
	
	if err != nil {
		if fe, ok := err.(*fiber.Error); ok {
			return ctx.Status(fe.Code).JSON(fiber.Map{
				"error": fe.Message,
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(transactions)
}

func (c *TransactionController) Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Transaction ID is required"})
	}

	transaction, err := c.transactionUsecase.Show(ctx, id)
	
	if err != nil {
		if fe, ok := err.(*fiber.Error); ok {
			return ctx.Status(fe.Code).JSON(fiber.Map{
				"error": fe.Message,
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(transaction)
}

func (c *TransactionController) DeleteMultiple(ctx *fiber.Ctx) error {
	var req models.DeleteMultipleRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if len(req.IDs) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "IDs are required"})
	}

	err := c.transactionUsecase.DeleteMultiple(ctx, req.IDs)
	
	if err != nil {
		if fe, ok := err.(*fiber.Error); ok {
			return ctx.Status(fe.Code).JSON(fiber.Map{
				"error": fe.Message,
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
