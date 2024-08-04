package request

// CreateUserRequest represents the request structure for creating a new user
// @Description Create user request structure
type CreateUserRequest struct {
	UserName string `json:"user_name" validate:"required,min=3,max=255" example:"Pepe" extensions:"x-order=0"`     // User name
	Email    string `json:"email" validate:"required,email" example:"example@example.com" extensions:"x-order=1"`  // User email
	Password string `json:"password" validate:"required,min=8,max=255" example:"012345678" extensions:"x-order=2"` // User password
	Age      int    `json:"age" validate:"required,min=18" example:"25" extensions:"x-order=3"`                    // User age
}
