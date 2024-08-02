package response

// UserResponse represents the response structure for user data
type UserResponse struct {
	ID       uint   `json:"id" validate:"required"`                      // User ID (primary key) in the database
	UserName string `json:"user_name" validate:"required,min=3,max=255"` // User name
	Email    string `json:"email" validate:"required,email"`             // User email
	Age      int    `json:"age" validate:"required,gte=18"`              // User age
}
