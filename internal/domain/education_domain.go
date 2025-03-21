package domain

import "github.com/gofiber/fiber/v2"

/**----------------------
 * ERROR
 *------------------------**/
var (
	ErrEducationNotFound = fiber.NewError(fiber.StatusNotFound, "Education not found")
)

/**----------------------
 * MESSAGE
 *------------------------**/
var (
	EDUCATION_CREATE_SUCCESS = "Create education success"
	EDUCATION_CREATE_FAILED  = "Create education failed"

	EDUCATION_DETAIL_SUCCESS = "Get education success"
	EDUCATION_DETAIL_FAILED  = "Get education failed"

	EDUCATION_UPDATE_SUCCESS = "Update education success"
	EDUCATION_UPDATE_FAILED  = "Update education failed"

	EDUCATION_DELETE_SUCCESS = "Delete education success"
	EDUCATION_DELETE_FAILED  = "Delete education failed"

	EDUCATION_GET_BY_USER_SUCCESS = "Get education by user success"
	EDUCATION_GET_BY_USER_FAILED  = "Get education by user failed"
)

/**----------------------
 * REQUEST
 *------------------------**/
type (
	EducationCreateRequest struct {
		School      string `json:"school" validate:"required,max=100"`
		Degree      string `json:"degree" validate:"max=100"`
		Description string `json:"description" validate:"max=1000"`
		StartDate   string `json:"start_date" validate:"required"`
		EndDate     string `json:"end_date" validate:"required"`
	}

	EducationUpdateRequest = EducationCreateRequest
)

/**----------------------
 * RESPONSE
 *------------------------**/
type (
	EducationDetailResponse struct {
		ID          string `json:"id"`
		School      string `json:"school"`
		Degree      string `json:"degree"`
		Description string `json:"description"`
		StartDate   string `json:"start_date"`
		EndDate     string `json:"end_date"`
	}
)
