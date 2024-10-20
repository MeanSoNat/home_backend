package model

type Student struct {
	ID            int    `gorm:"type:autoIncrement;primaryKey"`
	Username      string `gorm:"type:varchar(255)"`
	Password      string `gorm:"type:varchar(255)"` // รหัส��่าน
	Name          string `gorm:"type:varchar(255)"` // ชื่อ
	Nickname      string `gorm:"type:varchar(255)"` // ชื่อเล่น
	Idcard        string `gorm:"type:varchar(255)"` //
	StuId         int
	StuPhone      string `gorm:"type:varchar(100)"` //
	StuClass      string `gorm:"type:varchar(10)"`  //
	StuBirthDate  string `gorm:"type:varchar(255)"` //
	Address       string `gorm:"type:varchar(255)"` // ที่อยู่
	Distance      string `gorm:"type:varchar(100)"` // ระยะทาง
	Transport     string `gorm:"type:varchar(255)"` // การเดินทา
	Skills        string `gorm:"type:varchar(255)"` //ทักษะ
	FatherName    string `gorm:"type:varchar(255)"` // ชื่อบิดา
	FatherJob     string `gorm:"type:varchar(255)"` // อาชืพบิดา
	FatherPhone   string `gorm:"type:varchar(255)"` // เบอร์บิดา
	FatherSalary  int    // เงินเดือนบิดา
	FatherEdu     string `gorm:"type:varchar(255)"` //การศึกษาบิดา
	MotherName    string `gorm:"type:varchar(255)"` // ชื่อมารดา
	MotherJob     string `gorm:"type:varchar(255)"` // อา��ีพมารดา
	MotherPhone   string `gorm:"type:varchar(255)"` // เบอร์มารดา
	MotherSalary  int    // เงินเดือนมารดา
	MotherEdu     string `gorm:"type:varchar(255)"` //การ����ก��ามารดา
	ParentName    string `gorm:"type:varchar(255)"` // ชื่อ��ู้ปกครอง
	Relation      string `gorm:"type:varchar(255)"` //
	ParentPhone   string `gorm:"type:varchar(255)"` // เบอร์�ู้ปกครอง
	ParentAddress string `gorm:"type:varchar(255)"` // ที่อยู่�ู้ปกครอง\
	PStatus       string `gorm:"type:varchar(255)"` // สถานะ�ู้ปกครอง
	LivesWith     string `gorm:"type:varchar(255)"` // อา��ัยอยู่กับ
	FamCount      string `gorm:"type:varchar(255)"` //
	SibStudy      string `gorm:"type:varchar(255)"`
	EmpCount      string `gorm:"type:varchar(255)"`
	UnempCount    string `gorm:"type:varchar(255)"`
}

type SignInStudent struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
