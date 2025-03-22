package domain

import "github.com/gofiber/fiber/v2"

/**----------------------
 * MESSAGE
 *------------------------**/
var (
	PROFILE_GET_SUCCESS = "Successfully get profile"
	PROFILE_GET_FAILED  = "Failed to get profile"

	PROFILE_UPDATE_SUCCESS = "Successfully update profile"
	PROFILE_UPDATE_FAILED  = "Failed to update profile"
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
type (
	ProfileEditRequest struct {
		FirstName string `json:"first_name" form:"first_name" validate:"required"`
		LastName  string `json:"last_name" form:"last_name" validate:"required"`
		Headline  string `json:"headline" form:"headline" validate:"max=100"`
		About     string `json:"about" form:"about" validate:"max=500"`
	}
)

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
