package model

type Class struct {
	ID int64
	StuClass string `gorm:"type:varchar(10)"`
}