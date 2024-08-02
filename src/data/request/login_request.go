package request

// LoginRequest represents the request structure for user login
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"` // User email
	Password string `json:"password" validate:"required"`    // User password
}
