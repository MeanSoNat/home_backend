package routes

import (
    "github.com/Makeyabe/Home_Backend/controllers"
    "github.com/gin-gonic/gin"
)

func StudentRoutes(router *gin.RouterGroup, studentController *controllers.StudentController) {
    router.POST("/student/login", studentController.StudentLogin)
}
