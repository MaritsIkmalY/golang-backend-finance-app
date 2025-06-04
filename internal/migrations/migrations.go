package migrations

import (
	"log"

	"github.com/maritsikmaly/golang-finance-app/database"
	"github.com/maritsikmaly/golang-finance-app/internal/entities"
)

func MigrateDatabase(db database.Database) {
	if err := db.GetDb().AutoMigrate(
		&entities.User{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
