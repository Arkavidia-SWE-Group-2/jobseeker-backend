package domain

import "github.com/gofiber/fiber/v2"

/**----------------------
 * ERROR
 *------------------------**/
var (
	ErrInvalidParameter = fiber.NewError(fiber.StatusBadRequest, "Invalid parameter")
)
