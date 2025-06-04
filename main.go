package main

import (
	"github.com/maritsikmaly/golang-finance-app/config"
	"github.com/maritsikmaly/golang-finance-app/database"
	"github.com/maritsikmaly/golang-finance-app/internal/migrations"
	"github.com/maritsikmaly/golang-finance-app/server"
)

func main() {
	config := config.GetConfig()

	db := database.NewPostgresDatabase(config)

	migrations.MigrateDatabase(db)

	server.NewFiberServer(db, config).Start()
}