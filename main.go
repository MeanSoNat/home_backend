package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Makeyabe/Home_Backend/controllers"
	"github.com/Makeyabe/Home_Backend/initializers"
	"github.com/Makeyabe/Home_Backend/model"
	"github.com/Makeyabe/Home_Backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server                 *gin.Engine
	AuthController         controllers.AuthController
	AuthRouteController    routes.AuthRouteController
	AdminController        controllers.AdminController
	AdminRouteController   routes.AdminRouteController
	TeacherController      *controllers.TeacherController
	TeacherRouteController routes.TeacherRouteController
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

	AdminController = controllers.NewAdminController(initializers.DB)      // กำหนดค่า AdminController
	AdminRouteController = routes.NewAdminRouteController(AdminController) // กำหนดค่า AdminRouteController

	TeacherController = controllers.NewTeacherController(initializers.DB)        // กำหนดค่า TeacherController
	TeacherRouteController = routes.NewTeacherRouteController(TeacherController) // กำหนดค่า TeacherRouteController

	StudentController = controllers.NewStudentController(initializers.DB) // กำหนดค่า StudentController

	err = initializers.DB.AutoMigrate(
		&model.Student{},
		&model.Teacher{},
		&model.Booking{},
		&model.Summary{},
		&model.Formvisit{},
		&model.Formcheck{},
		&model.SectionForm{},
		&model.Subsection{},
		&model.Field{},
		&model.Option{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	fmt.Println("? Migration complete")

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
	AdminRouteController.AdminRoute(router)         // เส้นทาง Admin
	TeacherRouteController.TeacherRoutes(router)    // เส้นทาง Teacher
	routes.StudentRoutes(router, StudentController) // เส้นทาง Student

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Route Not Found"})
	})

	log.Fatal(server.Run(":" + config.ServerPort))

}
