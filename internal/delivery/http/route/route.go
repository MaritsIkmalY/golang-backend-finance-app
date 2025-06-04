package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/maritsikmaly/golang-finance-app/internal/delivery/http"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *http.UserController
}

func (c *RouteConfig) SetupRoutes() {
	auth := c.App.Group("/api/v1/finance-app")
	guestRoutes(c, auth)
}

func guestRoutes(c *RouteConfig, auth fiber.Router) {
	auth.Post("/register", c.UserController.Register)
	auth.Post("/login", c.UserController.Login)
}
