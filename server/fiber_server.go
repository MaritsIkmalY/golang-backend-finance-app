package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/maritsikmaly/golang-finance-app/config"
	"github.com/maritsikmaly/golang-finance-app/database"
	"github.com/maritsikmaly/golang-finance-app/internal/delivery/http"
	"github.com/maritsikmaly/golang-finance-app/internal/delivery/http/route"
	"github.com/maritsikmaly/golang-finance-app/internal/repositories"
	"github.com/maritsikmaly/golang-finance-app/internal/usecases"
)

type fiberServer struct {
	app    *fiber.App
	db     database.Database
	config *config.Config
}

func NewFiberServer(db database.Database, con *config.Config) Server {
	fiberApp := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	return &fiberServer{
		app:    fiberApp,
		db:     db,
		config: con,
	}
}

func (s *fiberServer) Start() {
	s.app.Use(logger.New())
	s.app.Use(cors.New())

	userRepository := repositories.NewUserRepository(s.db.GetDb())
	userUsecase := usecases.NewUserUseCase(userRepository)
	userController := http.NewUserController(userUsecase)

	routeConfig := &route.RouteConfig{
		App:            s.app,
		UserController: userController,
	}

	routeConfig.SetupRoutes()

	for _, routes := range s.app.Stack() {
		for _, route := range routes {
			fmt.Printf("[%s] %s\n", route.Method, route.Path)
		}
	}

	log.Printf("Server starting on port %s", s.config.Server.Port)
	log.Fatal(s.app.Listen(":" + s.config.Server.Port))
}
