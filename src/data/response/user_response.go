package response

type UserResponse struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}
