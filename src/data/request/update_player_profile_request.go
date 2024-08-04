package request

// UpdatePlayerProfileRequest represents the request structure for updating player profile data
// @Description Update player profile data
type UpdatePlayerProfileRequest struct {
	Nickname   string `json:"nickname" validate:"required,min=3,max=20" example:"NoobMaster69" extensions:"x-order=0"`        // Player nickname
	Avatar     string `json:"avatar" validate:"required" example:"https://example.com/avatar-new.png" extensions:"x-order=1"` // Player avatar URL
	Level      int    `json:"level" validate:"required" example:"2" extensions:"x-order=2"`                                   // Player level
	Experience int    `json:"experience" validate:"required" example:"200" extensions:"x-order=3"`                            // Player experience
	Points     int    `json:"points" validate:"required" example:"200" extensions:"x-order=4"`                                // Player points
}
