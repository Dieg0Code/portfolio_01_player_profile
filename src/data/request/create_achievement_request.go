package request

// CreateAchievementRequest represents the request structure for creating a new achievement
// @Description Create achievement request structure
type CreateAchievementRequest struct {
	Name        string `json:"name" validate:"required,min=5,max=255" example:"First blood"`                 // Achievement name
	Description string `json:"description" validate:"required,min=5,max=255" example:"Kill the first enemy"` // Achievement description
}
