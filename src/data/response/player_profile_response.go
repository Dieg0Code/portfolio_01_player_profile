package response

// PlayerProfileResponse represents the response structure for player profile data
// @Description Player profile response structure
type PlayerProfileResponse struct {
	ID         uint   `json:"id" validate:"required,gt=0" example:"1" extensions:"x-order=0"`                             // Player ID (primary key) in the database
	Nickname   string `json:"nickname" validate:"required" example:"elPepe123" extensions:"x-order=1"`                    // Player nickname
	Avatar     string `json:"avatar" validate:"required" example:"https://example.com/avatar.png" extensions:"x-order=2"` // Player avatar URL
	Level      int    `json:"level" validate:"required" example:"1" extensions:"x-order=3"`                               // Player level
	Experience int    `json:"experience" validate:"required" example:"100" extensions:"x-order=4"`                        // Player experience
	Points     int    `json:"points" validate:"required" example:"100" extensions:"x-order=5"`                            // Player points
	UserID     uint   `json:"user_id" validate:"required,gt=0" example:"1" extensions:"x-order=6"`                        // User ID (foreign key) in the database
}
