package response

// AchievementWithPlayers represents the response structure for achievement with players data
// @Description Achievement with players response structure
type AchievementWithPlayers struct {
	ID      uint           `json:"achievement_id" example:"1" extensions:"x-order=0"`             // Achievement ID
	Name    string         `json:"achievement_name" example:"First blood" extensions:"x-order=1"` // Achievement name
	Players []PlayerSumary `json:"players" extensions:"x-order=2"`                                // List of players who have the achievement
}

// PlayerSumary represents the response structure for player data used in AchievementWithPlayers
// @Description Player summary response structure
type PlayerSumary struct {
	ID       uint   `json:"plyer_id" example:"1" extensions:"x-order=0"`                // Player ID (primary key) in the database
	Nickname string `json:"player_nickname" example:"elPepe123" extensions:"x-order=1"` // Player nickname
}
