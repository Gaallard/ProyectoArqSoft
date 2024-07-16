package models

type Course struct {
	Id_course   int    `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:varchar(100);not null"`
}

type Courses []Course
