package request

// LoginRequest represents the request structure for user login
// @Description Login request structure
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"example@example.com" extensions:"x-order=0"` // User email
	Password string `json:"password" validate:"required" example:"012345678" extensions:"x-order=1"`              // User password
}
