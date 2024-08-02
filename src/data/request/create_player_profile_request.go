package request

// CreatePlayerProfileRequest represents the request structure for creating a new player profile
type CreatePlayerProfileRequest struct {
	Nickname   string `json:"nickname" validate:"required,min=3,max=20"` // Player nickname
	Avatar     string `json:"avatar" validate:"required"`                // Player avatar
	Level      int    `json:"level" validate:"required"`                 // Player level
	Experience int    `json:"experience" validate:"required"`            // Player experience
	Points     int    `json:"point" validate:"required"`                 // Player points
	UserID     uint   `json:"user_id" validate:"required,gt=0"`          // User ID (foreign key) in the database
}
