package request

type CreateUserRequest struct {
	UserName string `json:"user_name" validate:"required,min=3,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Age      int    `json:"age" validate:"required,min=18"`
}
