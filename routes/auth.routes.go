package routes

import (
	"github.com/Makeyabe/Home_Backend/controllers"
	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")
	router.POST("/login", rc.authController.SignInAdmin)
	router.GET("/refresh", rc.authController.RefreshAccessToken)
	// router.GET("/logout", middleware.DeserializeUser(), rc.authController.LogoutUser)
}
