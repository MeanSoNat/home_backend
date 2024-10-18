package model

type Teacher struct {
	ID        int
	Classid   int
	StuClass string `gorm:"type:varchar(20)"`
	Username string `gorm:"type:varchar(20)"`
	Password string `gorm:"type:varchar(255)"`
	Formid    int
	Bookingid int
	TeacherName string `gorm:"type:varchar(50)"`
	PhoneNumber string `gorm:"type:varchar(10)"`
}

type SignInTeacher struct {
    Username string `json:"username"`
    Password string `json:"password"`
}