package domain

var AUTH_USER = "AUTH_USER"

/**----------------------
 * MESSAGE
 *------------------------**/
var (
	AUTH_REGISTER_SUCCESS = "Register user success"
	AUTH_REGISTER_FAILED  = "Register user failed"
)

/**----------------------
 * REQUEST
 *------------------------**/
type (
	AuthRegisterRequest struct {
		FirstName string `json:"first_name" validate:"required"`
		LastName  string `json:"last_name" validate:"required"`
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required,min=8"`
		Phone     string `json:"phone" validate:"required,min=10"`
	}

	AuthLoginRequest struct {
		Credential string `json:"credential" validate:"required"`
		Password   string `json:"password" validate:"required,min=8"`
	}
)

/**----------------------
 * RESPONSE
 *------------------------**/
type (
	AuthLoginResponse struct {
		Token string `json:"token"`
	}
)
