package controllers

import (
	"net/http"

	"github.com/santhureinoo_agnos_backend/services"

	"github.com/gin-gonic/gin"
)

type PasswordRequest struct {
	InitPassword string `json:"init_password"`
}

func PasswordStrengthHandler(c *gin.Context) {

	// Parse request body
	var req PasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	numOfSteps := services.CalculateSteps(req.InitPassword)

	// Log request and response
	// Optionally call logging function here

	c.JSON(http.StatusOK, gin.H{"num_of_steps": numOfSteps})
}
