package database

import (
	"fmt"
	"sync"

	"github.com/maritsikmaly/golang-finance-app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	DB *gorm.DB
}

var (
  once       sync.Once
  dbInstance *postgresDatabase
)

func NewPostgresDatabase(conf *config.Config) Database {
  once.Do(func() {
    dsn := fmt.Sprintf(
      "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
      conf.DB.Host,
      conf.DB.User,
      conf.DB.Password,
      conf.DB.DBName,
      conf.DB.Port,
      conf.DB.SSLMode,
      conf.DB.TimeZone,
    )
    
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
      panic("failed to connect database")
    }
    
    dbInstance = &postgresDatabase{DB: db}
  })

  return dbInstance
}

func (p *postgresDatabase) GetDb() *gorm.DB {
  return dbInstance.DB
}