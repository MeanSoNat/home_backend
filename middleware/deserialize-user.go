package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Makeyabe/Home_Backend/initializers"
	"github.com/Makeyabe/Home_Backend/model"
	"github.com/Makeyabe/Home_Backend/utils"
	"github.com/gin-gonic/gin"
)

// DeserializeUser ใช้สำหรับดึงข้อมูลผู้ใช้จาก token
func DeserializeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var access_token string
		cookie, err := ctx.Cookie("access_token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		if access_token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config, _ := initializers.LoadConfig(".")
		sub, err := utils.ValidateToken(access_token, config.AccessTokenPublicKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		var admin model.Admin
		result := initializers.DB.First(&admin, "id = ?", fmt.Sprint(sub))
		if result.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "The user belonging to this token no longer exists"})
			return
		}

		ctx.Set("currentUser", admin)
		ctx.Next()
	}
}

// MiddlewareAdmin ใช้ตรวจสอบว่าผู้ใช้มีสิทธิ์เป็น admin
func MiddlewareAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// เรียกใช้ DeserializeUser เพื่อตรวจสอบว่าเข้าสู่ระบบหรือไม่
		DeserializeUser()(ctx)

		// ตรวจสอบว่า context มีค่า currentUser หรือไม่
		user, exists := ctx.Get("currentUser")
		if !exists {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Not logged in"})
			return
		}

		// แปลงเป็น model.Admin เพื่อเข้าถึง role
		admin, ok := user.(model.Admin)
		if !ok || admin.Role != "admin" { // ตรวจสอบว่า role เป็น admin หรือไม่
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You do not have permission to access this resource"})
			return
		}

		ctx.Next()
	}
}
