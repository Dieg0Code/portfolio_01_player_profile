package response

type AchievementResponse struct {
	ID          uint   `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required,min=3,max=255"`
	Description string `json:"description" validate:"required,min=5,max=255"`
}
