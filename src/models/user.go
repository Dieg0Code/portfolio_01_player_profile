package models

type User struct {
	UserID   int             `gorm:"type:int;primaryKey;autoIncrement"`
	UserName string          `gorm:"type:varchar(255);unique;not null"`
	PassWord string          `gorm:"type:varchar(255);not null"`
	Email    string          `gorm:"type:varchar(255);unique;not null"`
	Age      int             `gorm:"type:int;not null"`
	Profiles []PlayerProfile `gorm:"foreignKey:UserID"` // Relación uno a muchos con PlayerProfile
}
