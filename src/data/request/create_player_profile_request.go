package request

// CreatePlayerProfileRequest represents the request structure for creating a new player profile
// @Description Create player profile request structure
type CreatePlayerProfileRequest struct {
	Nickname   string `json:"nickname" validate:"required,min=3,max=20" example:"NoobMaster69" extensions:"x-order=0"`    // Player nickname
	Avatar     string `json:"avatar" validate:"required" example:"https://example.com/avatar.png" extensions:"x-order=1"` // Player avatar URL
	Level      int    `json:"level" validate:"required" example:"1" extensions:"x-order=2"`                               // Player level
	Experience int    `json:"experience" validate:"required" example:"100" extensions:"x-order=3"`                        // Player experience
	Points     int    `json:"point" validate:"required" example:"100" extensions:"x-order=4"`                             // Player points
	UserID     uint   `json:"user_id" validate:"required,gt=0" example:"1" extensions:"x-order=5"`                        // User ID (foreign key) in the database
}
