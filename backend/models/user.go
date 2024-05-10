package models

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(256);not null"`
}

type Users []User
