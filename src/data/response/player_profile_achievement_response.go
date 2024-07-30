package response

type PlayerProfileAchievementResponse struct {
	ID              uint   `json:"id" validate:"required,gt=0"`
	PlayerProfileID uint   `json:"player_profile_id" validate:"required,gt=0"`
	AchievementID   uint   `json:"achievement_id" validate:"required,gt=0"`
	CreatedAt       string `json:"created_at" validate:"required"`
}
