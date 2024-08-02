package request

// CreateUserRequest represents the request structure for creating a new user
type CreateUserRequest struct {
	UserName string `json:"user_name" validate:"required,min=3,max=255"` // User name
	Password string `json:"password" validate:"required,min=8,max=255"`  // User password
	Email    string `json:"email" validate:"required,email"`             // User email
	Age      int    `json:"age" validate:"required,min=18"`              // User age
}
