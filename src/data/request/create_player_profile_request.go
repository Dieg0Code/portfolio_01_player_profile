package request

type CreatePlayerProfileRequest struct {
	Nickname   string `json:"nickname" validate:"required,min=3,max=20"`
	Avatar     string `json:"avatar" validate:"required"`
	Level      int    `json:"level" validate:"required"`
	Experience int    `json:"experience" validate:"required"`
	Points     int    `json:"point" validate:"required"`
	UserID     uint   `json:"user_id" validate:"required,gt=0"`
}
