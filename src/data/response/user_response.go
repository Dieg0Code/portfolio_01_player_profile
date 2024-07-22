package response

type UserResponse struct {
	ID       uint   `json:"id" validate:"required"`
	UserName string `json:"user_name" validate:"required,min=3,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Age      int    `json:"age" validate:"required,gte=18"`
}
