package model

type Admin struct {
	ID       int
	Username string `gorm:"type:varchar(30)"`
	Password string `gorm:"type:varchar"`
	Role     string `gorm:"type:varchar(20)"` // เพิ่มฟิลด์ Role
}

type SignInAdmin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Status   string `json:"status"`
	Role     string `json:"role"` // เพิ่ม Role ใน response หากต้องการแสดงข้อมูลนี้ใน response ด้วย
}
