package request

// UpdatePlayerProfileRequest represents the request structure for updating player profile data
type UpdatePlayerProfileRequest struct {
	Nickname   string `json:"nickname" validate:"required,min=3,max=20"` // Player nickname
	Avatar     string `json:"avatar" validate:"required"`                // Player avatar
	Level      int    `json:"level" validate:"required"`                 // Player level
	Experience int    `json:"experience" validate:"required"`            // Player experience
	Points     int    `json:"points" validate:"required"`                // Player points
}
