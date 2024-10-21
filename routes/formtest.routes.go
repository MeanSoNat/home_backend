package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Makeyabe/Home_Backend/controllers"
)

// FormRoutes กำหนดเส้นทางสำหรับฟอร์ม
func FormRoutes(router *gin.RouterGroup, formController *controllers.FormController) {
	formRoute := router.Group("/forms")
	{
		formRoute.GET("/", formController.GetForms)          // ดึงข้อมูลฟอร์มทั้งหมด
		formRoute.POST("/", formController.CreateForm)       // สร้างฟอร์มใหม่
	}
}
