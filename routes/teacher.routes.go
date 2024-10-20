package routes

import (
    "github.com/Makeyabe/Home_Backend/controllers"
    "github.com/gin-gonic/gin"
)

type TeacherRouteController struct {
    teacherController *controllers.TeacherController // ใช้ pointer
}

func NewTeacherRouteController(teacherController *controllers.TeacherController) TeacherRouteController {
    return TeacherRouteController{teacherController}
}

func (trc *TeacherRouteController) TeacherRoutes(router *gin.RouterGroup) {
    router.POST("/teacher/login", trc.teacherController.Login) // Endpoint สำหรับล็อกอิน
    router.GET("/teacher/students/:id", trc.teacherController.GetStudentsByClass) // Endpoint สำหรับดึงนักเรียน
}
 

