package response

type PlayerProfileResponse struct {
	ID         uint   `json:"id" validate:"required,gt=0"`
	Nickname   string `json:"nickname" validate:"required"`
	Avatar     string `json:"avatar" validate:"required"`
	Level      int    `json:"level" validate:"required"`
	Experience int    `json:"experience" validate:"required"`
	Points     int    `json:"points" validate:"required"`
	UserID     uint   `json:"user_id" validate:"required,gt=0"`
}
