package helpers

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	response := &Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	return c.Status(status).JSON(response)
}

func ErrorResponse(c *fiber.Ctx, status int, message string) error {
	response := &Response{
		Status:  "error",
		Message: message,
	}
	return c.Status(status).JSON(response)
}
