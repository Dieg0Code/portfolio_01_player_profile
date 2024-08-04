package response

// AchievementResponse represents the response structure for achievement data
// @Description Achievement response structure
type AchievementResponse struct {
	ID          uint   `json:"id" validate:"required" example:"1" extensions:"x-order=0"`                                           // Achievement ID
	Name        string `json:"name" validate:"required,min=3,max=255" exampl:"First blood" extensions:"x-order=1"`                  // Achievement name
	Description string `json:"description" validate:"required,min=5,max=255" example:"Kill the first enemy" extensions:"x-order=2"` // Achievement description
}
