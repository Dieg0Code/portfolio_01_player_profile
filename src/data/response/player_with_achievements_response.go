package response

// PlayerWithAchievements represents the response structure for player with achievements data
// @Description Player with achievements response structure
type PlayerWithAchievements struct {
	ID           uint                 `json:"player_id" example:"1" extensions:"x-order=0"`               // Player ID (primary key) in the database
	Nickname     string               `json:"player_nickname" example:"elPepe123" extensions:"x-order=1"` // Player nickname
	Achievements []AchievementsSumary `json:"achievements" extensions:"x-order=2"`                        // List of player achievements
}

// AchievementsSumary represents the response structure for achievements data used in PlayerWithAchievements
// @Description Achievements summary response structure
type AchievementsSumary struct {
	ID   uint   `json:"achievement_id" example:"1" extensions:"x-order=0"`             // Achievement ID
	Name string `json:"achievement_name" example:"First blood" extensions:"x-order=1"` // Achievement name
}
