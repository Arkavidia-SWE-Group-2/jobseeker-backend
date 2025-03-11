package domain

/**----------------------
 * MESSAGE
 *------------------------**/
var (
	EDUCATION_CREATE_SUCCESS = "Create education success"
	EDUCATION_CREATE_FAILED  = "Create education failed"
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
