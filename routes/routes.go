package routes

import (
	"github.com/santhureinoo_agnos_backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/api/strong_password_steps", controllers.PasswordStrengthHandler)
}
