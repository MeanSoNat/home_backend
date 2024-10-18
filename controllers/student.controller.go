package controllers

import (
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
