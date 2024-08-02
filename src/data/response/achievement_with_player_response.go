package response

// AchievementWithPlayers represents the response structure for achievement with players data
type AchievementWithPlayers struct {
	ID      uint           `json:"achievement_id"`   // Achievement ID (primary key) in the database
	Name    string         `json:"achievement_name"` // Achievement name
	Players []PlayerSumary `json:"players"`          // Players who have the achievement
}

// PlayerSumary represents the response structure for player data used in AchievementWithPlayers
type PlayerSumary struct {
	ID       uint   `json:"plyer_id"`        // Player ID (primary key) in the database
	Nickname string `json:"player_nickname"` // Player nickname
}
