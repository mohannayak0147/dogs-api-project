package routes

import (
	"dogs-api/v1/http"

	"github.com/gofiber/fiber/v2"
)

// ErrorResp represents the structure of a standard error response.
type ErrorResp struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMsg"`
}

// StandardErrorHandler generates a standardized error message
func StandardErrorHandler(ctx *fiber.Ctx, err error) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	switch err := err.(type) {
	case http.HTTPError:
		return ctx.Status(err.StatusCode()).JSON(ErrorResp{Status: "error", ErrMsg: err.Error()})
	default:
		return ctx.Status(200).JSON(ErrorResp{Status: "error", ErrMsg: err.Error()})
	}
}
