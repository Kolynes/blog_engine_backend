package server

import "github.com/gofiber/fiber/v2"

type (
	Responder         func(*fiber.Ctx) error
	EndPoint          func(*fiber.Ctx, ...interface{}) (Responder, error)
	EndPointDecorator func(EndPoint) EndPoint

	IController interface {
		RegisterRoutes(fiber.Router)
	}

	IMiddleware interface {
		Handler(*fiber.Ctx) error
	}

	ControllerConfig struct {
		Route      string
		Controller IController
		Middleware []IMiddleware
	}

	ServerConfig struct {
		Host          string
		ErrorHandlers []IMiddleware
		Middleware    []IMiddleware
		Controllers   []ControllerConfig
	}
)
