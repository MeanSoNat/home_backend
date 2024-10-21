package model

// Formcheck เป็นโครงสร้างหลักสำหรับฟอร์ม
type Formcheck struct {
	ID           uint          `gorm:"primaryKey"`
	SectionForms []SectionForm `json:"sectionforms" gorm:"foreignKey:FormcheckID;references:ID"`
}

// SectionForm เป็นโครงสร้างสำหรับส่วนต่างๆ ของฟอร์ม
type SectionForm struct {
	ID           uint         `gorm:"primaryKey"`
	FormcheckID  uint         `json:"formcheck_id"`
	SectionLabel string       `json:"sectionLabel"`
	Subsections  []Subsection `json:"subsections" gorm:"foreignKey:SectionFormID;references:ID"`
}

// Subsection เป็นโครงสร้างสำหรับส่วนย่อยของฟอร์ม
type Subsection struct {
	ID            uint    `gorm:"primaryKey"`
	SectionFormID uint    `json:"section_form_id"`
	Label         string  `json:"label"`
	Fields        []Field `json:"fields" gorm:"foreignKey:SubsectionID;references:ID"`
}

// Field เป็นโครงสร้างสำหรับฟิลด์ในฟอร์ม
type Field struct {
	ID           uint     `gorm:"primaryKey"`
	Label        string   `json:"label"`
	InputType    string   `json:"inputType"`
	Options      []Option `json:"options" gorm:"foreignKey:FieldID;references:ID"`
	SubsectionID uint     `json:"subsection_id"` // เพิ่มฟิลด์นี้เพื่อเก็บ foreign key
}

// Option เป็นโครงสร้างสำหรับตัวเลือกในฟิลด์
type Option struct {
	ID      uint   `gorm:"primaryKey"`
	FieldID uint   `json:"field_id"`
	Label   string `json:"label"`
	Checked bool   `json:"checked"`
}
