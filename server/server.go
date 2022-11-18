package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func StartServer(serverConfig ServerConfig) error {
	app := fiber.New()
	app.Use(recover.New())
	for _, m := range serverConfig.Middleware {
		app.Use(m.Handler)
	}
	for _, c := range serverConfig.Controllers {
		app.Route(c.Route, func(router fiber.Router) {
			for _, m := range c.Middleware {
				router.Use(m.Handler)
			}
			c.Controller.RegisterRoutes(router)
		})
	}
	return app.Listen(serverConfig.Host)
}
