package model

import "time"

type Booking struct {
	ID          int       `gorm:"type:autoIncrement;primaryKey"`
	StuName     string    `gorm:"type:varchar(30)"`
	StuId       int       `gorm:"type:integer"`
	BookingDate time.Time `gorm:"type:timestamp"`
	StuClass    string    `gorm:"type:varchar(10)"`
}
