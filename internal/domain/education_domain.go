package domain

/**----------------------
 * MESSAGE
 *------------------------**/
var (
	EDUCATION_CREATE_SUCCESS = "Create education success"
	EDUCATION_CREATE_FAILED  = "Create education failed"

	EDUCATION_DETAIL_SUCCESS = "Get education success"
	EDUCATION_DETAIL_FAILED  = "Get education failed"
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
