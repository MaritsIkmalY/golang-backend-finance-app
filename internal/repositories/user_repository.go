package repositories

import (
	"github.com/maritsikmaly/golang-finance-app/internal/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entities.User) (error)
	GetByEmail(email string) (*entities.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *entities.User) (error) {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByEmail(email string) (*entities.User, error) {
	var user entities.User

	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
