package request

// UpdateUserRequest represents the request structure for updating user data
// @Description Update user request structure
type UpdateUserRequest struct {
	UserName string `json:"user_name,omitempty" validate:"required,min=3,max=20" example:"Pepe new" extensions:"x-order=0"` // User name
	Email    string `json:"email,omitempty" validate:"required,email" example:"new@example.com" extensions:"x-order=1"`     // User email
	Age      int    `json:"age,omitempty" validate:"required,gte=18" example:"25" extensions:"x-order=2"`                   // User age
}
