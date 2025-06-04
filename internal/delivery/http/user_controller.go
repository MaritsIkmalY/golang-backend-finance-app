package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/maritsikmaly/golang-finance-app/internal/models"
	"github.com/maritsikmaly/golang-finance-app/internal/usecases"
)

type UserController struct {
	userUsecase usecases.UserUseCase
	validator *validator.Validate
}

func NewUserController(userUsecase usecases.UserUseCase) *UserController{
	return &UserController{
		userUsecase: userUsecase,
		validator:   validator.New(),
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	var req models.RegisterUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := c.validator.Struct(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response, err := c.userUsecase.Register(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"data":    response,
	})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var req models.LoginUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := c.validator.Struct(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response, err := c.userUsecase.Login(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User logged in successfully",
		"data":    response,
	})
}
