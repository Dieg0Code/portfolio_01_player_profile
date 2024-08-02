package request

// CreateAchievementRequest represents the request structure for creating a new achievement
type CreateAchievementRequest struct {
	Name        string `json:"name" validate:"required,min=5,max=255"`        // Achievement name
	Description string `json:"description" validate:"required,min=5,max=255"` // Achievement description
}
