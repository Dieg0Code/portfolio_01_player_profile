package request

// UpdateUserRequest represents the request structure for updating user data
type UpdateUserRequest struct {
	UserName string `json:"user_name,omitempty" validate:"required,min=3,max=20"` // User name
	Email    string `json:"email,omitempty" validate:"required,email"`            // User email
	Age      int    `json:"age,omitempty" validate:"required,gte=18"`             // User age
}
