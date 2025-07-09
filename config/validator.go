package config

import (
	"github.com/go-playground/validator/v10"
)

const (
    CategoryIncome  = "income"
    CategoryExpense = "expenses"
)

func NewValidator() *validator.Validate {
	validator := validator.New()
	
	validator.RegisterValidation("category_valid", categoryValidator)

	return validator
}

func categoryValidator(fl validator.FieldLevel) bool {
	category := fl.Field().String()
	return category == CategoryIncome || category == CategoryExpense
}