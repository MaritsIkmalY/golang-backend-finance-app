package main

import (
	"github.com/maritsikmaly/golang-finance-app/config"
	"github.com/maritsikmaly/golang-finance-app/database"
	"github.com/maritsikmaly/golang-finance-app/internal/migrations"
	"github.com/maritsikmaly/golang-finance-app/server"
)

func main() {
	conf := config.GetConfig()

	db := database.NewPostgresDatabase(conf)

	migrations.MigrateDatabase(db)

	validator := config.NewValidator()

	server.NewFiberServer(db, conf, validator).Start()
}