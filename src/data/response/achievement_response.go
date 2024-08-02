package response

// AchievementResponse represents the response structure for achievement data
type AchievementResponse struct {
	ID          uint   `json:"id" validate:"required"`                        // Achievement ID (primary key) in the database
	Name        string `json:"name" validate:"required,min=3,max=255"`        // Achievement name
	Description string `json:"description" validate:"required,min=5,max=255"` // Achievement description
}
