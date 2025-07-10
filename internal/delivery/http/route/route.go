package route

import (
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/maritsikmaly/golang-finance-app/internal/delivery/http"
	"github.com/maritsikmaly/golang-finance-app/internal/delivery/http/middleware"
)

type RouteConfig struct {
	App                   *fiber.App
	UserController        *http.UserController
	TransactionController *http.TransactionController
}

func (c *RouteConfig) SetupRoutes() {
	api := c.App.Group("/api/v1/finance-app")
	guestRoutes(c, api)

	auth := api.Group("/", middleware.AuthMiddleware(os.Getenv("JWT_SECRET")))

	authRoutes(c, auth)
}

func guestRoutes(c *RouteConfig, auth fiber.Router) {
	auth.Post("/register", c.UserController.Register)
	auth.Post("/login", c.UserController.Login)
}

func authRoutes(c *RouteConfig, auth fiber.Router) {
	auth.Post("/transactions", c.TransactionController.Create)
	auth.Get("/transactions", c.TransactionController.GetByUserID)
	auth.Get("/transactions/:id", c.TransactionController.Show)
	auth.Put("/transactions/:id", c.TransactionController.Update)
	auth.Delete("/transactions/:id", c.TransactionController.Delete)
	auth.Delete("/transactions", c.TransactionController.DeleteMultiple)
}
