package entities

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID          uint    `gorm:"primaryKey"`
	Amount      float64 `gorm:"not null"`
	Description string
	Category    string          `gorm:"not null"`
	Date        time.Time       `gorm:"not null;type:date"`
	CreatedAt   time.Time       `gorm:"autoCreateTime"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime"`
	UserID      uint            `gorm:"not null;index"`
	User        User            `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DeletedAt   *gorm.DeletedAt `json:"-" gorm:"index"`
}
