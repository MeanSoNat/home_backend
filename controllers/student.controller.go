package controllers

import (
	"log"
	"net/http"

	"github.com/Makeyabe/Home_Backend/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentController struct {
	DB *gorm.DB
}

func NewStudentController(db *gorm.DB) *StudentController {
	return &StudentController{DB: db}
}

type StudentLoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (sc *StudentController) StudentLogin(c *gin.Context) {
	var input StudentLoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	var student model.Student // สมมุติว่ามี model ที่ชื่อว่า Student
	if err := sc.DB.Where("username = ?", input.Username).First(&student).Error; err != nil {
		c.JSON(400, gin.H{"message": "Student not found"})
		return
	}

	if student.Password != input.Password {
		c.JSON(400, gin.H{"message": "Incorrect password"})
		return
	}

	// หากต้องการให้ลบส่วนเกี่ยวกับ token ที่นี่ เช่น
	c.JSON(200, gin.H{"message": "Login successful"})
}

func (sc *StudentController) GetStudentData(c *gin.Context) {
	var students []model.Student

	// ดึงข้อมูลนักเรียนทั้งหมดจากฐานข้อมูล
	if err := sc.DB.Find(&students).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve students"})
		return
	}

	c.JSON(200, students) // ส่งข้อมูลนักเรียนกลับไปยัง client
}

func (sc *StudentController) GetStudentByID(c *gin.Context) {
	// ดึงค่า ID จากพารามิเตอร์ใน URL
	studentID := c.Param("id")

	var student model.Student

	// ดึงข้อมูลนักเรียนจากฐานข้อมูลโดยใช้ ID
	if err := sc.DB.Where("username = ?", studentID).First(&student).Error; err != nil {
		c.JSON(404, gin.H{"error": "Student not found"})
		return
	}

	// ส่งข้อมูลนักเรียนกลับไปยัง client
	c.JSON(200, student)
}

func (sc *StudentController) UpdateStudent(ctx *gin.Context) {
	var student model.Student // Variable to hold the updated data
	id := ctx.Param("id")     // Get the student ID from the URL
	log.Println("Student",id)  
	
	// Bind the JSON body to the student struct
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use GORM to update the student in the database
	if err := sc.DB.Model(&student).Where("username = ?", id).Updates(student).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Student updated successfully"})
}
