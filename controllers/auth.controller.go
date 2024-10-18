package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Makeyabe/Home_Backend/initializers"
	"github.com/Makeyabe/Home_Backend/model"
	"github.com/Makeyabe/Home_Backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

func (ac *AuthController) SignInAdmin(ctx *gin.Context) {
	var payload model.SignInAdmin // ไม่จำเป็นต้องใช้ pointer ที่นี่

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid input"})
		return
	}

	var admin model.Admin
	result := ac.DB.First(&admin, "username = ?", strings.ToLower(payload.Username))
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Incorrect username or password"})
		return
	}

	if err := utils.VerifyPassword(admin.Password, payload.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Incorrect username or password"})
		return
	}

	config, err := initializers.LoadConfig(".")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Could not load configuration"})
		return
	}

	// Generate Tokens
	accessToken, err := utils.CreateToken(config.AccessTokenExpiresIn, admin.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Could not create access token"})
		return
	}

	// ถ้าต้องการ refresh token ให้เพิ่มการสร้างที่นี่

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
}

func (ac *AuthController) RefreshAccessToken(ctx *gin.Context) {
	message := "Could not refresh access token"

	cookie, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": message})
		return
	}

	config, err := initializers.LoadConfig(".")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Could not load configuration"})
		return
	}

	sub, err := utils.ValidateToken(cookie, config.RefreshTokenPublicKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var admin model.Admin
	result := ac.DB.First(&admin, "id = ?", fmt.Sprint(sub))
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "User associated with this token no longer exists"})
		return
	}

	accessToken, err := utils.CreateToken(config.AccessTokenExpiresIn, admin.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", accessToken, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
}
