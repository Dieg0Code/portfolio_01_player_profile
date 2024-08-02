package response

// PlayerProfileResponse represents the response structure for player profile data
type PlayerProfileResponse struct {
	ID         uint   `json:"id" validate:"required,gt=0"`      // Player profile ID (primary key) in the database
	Nickname   string `json:"nickname" validate:"required"`     // Player nickname
	Avatar     string `json:"avatar" validate:"required"`       // Player avatar
	Level      int    `json:"level" validate:"required"`        // Player level
	Experience int    `json:"experience" validate:"required"`   // Player experience points
	Points     int    `json:"points" validate:"required"`       // Player points
	UserID     uint   `json:"user_id" validate:"required,gt=0"` // User ID (foreign key) in the database
}
