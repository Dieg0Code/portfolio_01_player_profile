package models

type PlayerProfile struct {
	PlayerProfileID int           `gorm:"type:int;primaryKey;autoIncrement"`
	Nickname        string        `gorm:"type:varchar(255);unique;not null"`
	Avatar          string        `gorm:"type:varchar(255);not null"`
	Level           int           `gorm:"type:int;not null"`
	Experience      int           `gorm:"type:int;not null"`
	Points          int           `gorm:"type:int;not null"`
	UserID          int           `gorm:"type:int;not null"`          // Clave foránea
	User            User          `gorm:"foreignKey:UserID"`          // Relación con User
	Achievements    []Achievement `gorm:"foreignKey:PlayerProfileID"` // Relación uno a muchos con Achievement
}
