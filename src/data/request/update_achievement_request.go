package request

// UpdateAchievementRequest represents the request structure for updating achievement data
// @Description Update achievement request structure
type UpdateAchievementRequest struct {
	Name        string `json:"name" validate:"required,min=5,max=255" example:"First blood updated" extensions:"x-order=0"`         // Achievement name
	Description string `json:"description" validate:"required,min=5,max=255" example:"Kill the first enemy" extensions:"x-order=1"` // Achievement description
}
