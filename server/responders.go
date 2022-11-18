package server

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func JsonResponder[T any](response JsonResponse[T]) Responder {
	return func(context *fiber.Ctx) error {
		b, err := json.Marshal(response)
		context.Response().Header.Set("Content-Type", "application/json")
		context.SendStatus(response.Status)
		context.SendString(string(b))
		return err
	}
}
