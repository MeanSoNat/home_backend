package model

import "time"

type Booking struct {
	ID          int    `gorm:"type:autoIncrement;primaryKey"`
	StuName     string `gorm:"type:varchar(30)"`
	StuId       int
	BookingDate time.Time
}
