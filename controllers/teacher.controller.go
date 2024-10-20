package controllers

import (
	"net/http"

	"github.com/Makeyabe/Home_Backend/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TeacherController struct {
	DB *gorm.DB
}

func NewTeacherController(db *gorm.DB) *TeacherController {
	return &TeacherController{DB: db} // คืนค่า pointer
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (tc *TeacherController) Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	var teacher model.Teacher
	if err := tc.DB.Where("username = ?", input.Username).First(&teacher).Error; err != nil {
		c.JSON(400, gin.H{"error": "Teacher not found"})
		return
	}

	if teacher.Password != input.Password {
		c.JSON(400, gin.H{"error": "Incorrect password"})
		return
	}

	c.JSON(200, gin.H{"message": "Login successful"})
}

func (tc *TeacherController) GetStudentsByClass(c *gin.Context) {
    // ดึง ID ของคุณครูจาก URL parameter
    teacherID := c.Param("id")

    var teacher model.Teacher
    var students []model.Student

    // ค้นหาข้อมูลของคุณครูจาก teacherID
    if err := tc.DB.Where("username = ?", teacherID).First(&teacher).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve teacher"})
        return
    }

    // ดึงค่า teacherClass จากข้อมูลคุณครู
    teacherClass := teacher.StuClass

    // ค้นหานักเรียนที่มี stuClass ตรงกับ teacherClass ของคุณครู
    if err := tc.DB.Where("stu_class = ?", teacherClass).Find(&students).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve students"})
        return
    }

    // ส่งข้อมูลนักเรียนที่ตรงกับ stuClass กลับไปยัง client
    c.JSON(http.StatusOK, students)
}
