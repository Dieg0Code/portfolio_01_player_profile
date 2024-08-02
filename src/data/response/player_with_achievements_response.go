package response

// PlayerWithAchievements represents the response structure for player with achievements data
type PlayerWithAchievements struct {
	ID           uint                 `json:"player_id"`       // Player ID (primary key) in the database
	Nickname     string               `json:"player_nickname"` // Player nickname
	Achievements []AchievementsSumary `json:"achievements"`    // Player achievements
}

// AchievementsSumary represents the response structure for achievements data used in PlayerWithAchievements
type AchievementsSumary struct {
	ID   uint   `json:"achievement_id"`   // Achievement ID (primary key) in the database
	Name string `json:"achievement_name"` // Achievement name
}
