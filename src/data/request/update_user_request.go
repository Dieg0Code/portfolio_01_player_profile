package request

type UpdateUserRequest struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name,omitempty"` // omitempty para permitir actualizaciones parciales
	PassWord string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Age      int    `json:"age,omitempty"`
}
