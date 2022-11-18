package server

import (
	"github.com/gofiber/fiber/v2"
)

func ToHandler(endPoint EndPoint, decorators ...EndPointDecorator) func(*fiber.Ctx) error {
	decoratedEndPoint := endPoint
	for _, decorator := range decorators {
		decoratedEndPoint = decorator(decoratedEndPoint)
	}
	return func(context *fiber.Ctx) error {
		responder, err := decoratedEndPoint(context)
		if err != nil {
			return err
		}
		return responder(context)
	}
}
