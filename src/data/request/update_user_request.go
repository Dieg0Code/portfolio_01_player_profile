package request

type UpdateUserRequest struct {
	UserName string `json:"user_name,omitempty" validate:"required, min=3, max=20"`
	Email    string `json:"email,omitempty" validate:"required, email"`
	Age      int    `json:"age,omitempty" validate:"required, gte=18"`
}
