package request

// UpdateAchievementRequest represents the request structure for updating achievement data
type UpdateAchievementRequest struct {
	Name        string `json:"name" validate:"required,min=5,max=255"`        // Achievement name
	Description string `json:"description" validate:"required,min=5,max=255"` // Achievement description
}
