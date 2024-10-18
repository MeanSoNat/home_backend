package main

import (
	"log"
	"net/http"

	"github.com/Makeyabe/Home_Backend/controllers"
	"github.com/Makeyabe/Home_Backend/initializers"
	"github.com/Makeyabe/Home_Backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server                 *gin.Engine
	AuthController         controllers.AuthController
	AuthRouteController    routes.AuthRouteController //มาเป็นคู่
	AdminController        controllers.AdminController
	AdminRouteController   routes.AdminRouteController //มาเป็นคู่
	TeacherController      *controllers.TeacherController 
	TeacherRouteController routes.TeacherRouteController // มาเป็นคู่ใหม่สำหรับครู
	StudentController      *controllers.StudentController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	TeacherController = controllers.NewTeacherController(initializers.DB)         // กำหนดค่า TeacherController
	TeacherRouteController = routes.NewTeacherRouteController(TeacherController) // กำหนดค่า TeacherRouteController

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Access-Control-Allow-Origin", "*"}

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	AdminRouteController.AdminRoute(router)
	TeacherRouteController.TeacherRoutes(router) // เพิ่มเส้นทางล็อกอินสำหรับครู
	routes.StudentRoutes(router, StudentController)

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Route Not Found"})
	})

	log.Fatal(server.Run(":" + config.ServerPort))
}
