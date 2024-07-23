package request

type UpdateAchievementRequest struct {
	Name        string `json:"name" validate:"required,min=5,max=255"`
	Description string `json:"description" validate:"required,min=5,max=255"`
}
