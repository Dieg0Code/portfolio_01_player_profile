package models

type Achievement struct {
	AchievementID   int           `gorm:"type:int;primaryKey;autoIncrement"`
	Name            string        `gorm:"type:varchar(255);not null"`
	Description     string        `gorm:"type:varchar(255);not null"`
	PlayerProfileID int           `gorm:"type:int;not null"`          // Clave foránea
	PlayerProfile   PlayerProfile `gorm:"foreignKey:PlayerProfileID"` // Relación con PlayerProfile
}
