package response

// UserResponse represents the response structure for user data
// @Description User response structure
type UserResponse struct {
	ID       uint   `json:"id" validate:"required" example:"1" extensions:"x-order=0"`                         // User ID
	UserName string `json:"user_name" validate:"required,min=3,max=255" example:"Pepe" extensions:"x-order=1"` // User name
	Email    string `json:"email" validate:"required,email" example:"examaple@example" extensions:"x-order=2"` // User email
	Age      int    `json:"age" validate:"required,gte=18" example:"25" extensions:"x-order=3"`                // User age
}
