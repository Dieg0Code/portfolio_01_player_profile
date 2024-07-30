package request

type CreatePlayerProfileAchievementRequest struct {
	PlayerProfileID uint `json:"player_profile_id" validate:"required,gt=0"`
	AchievementID   uint `json:"achievement_id" validate:"required,gt=0"`
}
