package domain

import "github.com/gofiber/fiber/v2"

/**----------------------
 * MESSAGE
 *------------------------**/
var (
	PROFILE_GET_SUCCESS = "Successfully get profile"
	PROFILE_GET_FAILED  = "Failed to get profile"
)

/**----------------------
 * ERROR
 *------------------------**/
var (
	ErrProfileNotFound = fiber.NewError(fiber.StatusNotFound, "Profile not found")
)

/**----------------------
 * REQUEST
 *------------------------**/

/**----------------------
 * RESPONSE
 *------------------------**/
type (
	ProfileResponse struct {
		ID        string `json:"id"`
		Vanity    string `json:"vanity"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Photo     string `json:"photo"`
		Headline  string `json:"headline"`
		About     string `json:"about"`
	}
)
