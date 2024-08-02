package middlewares

import (
	"time"

	"github.com/santhureinoo_agnos_backend/models"
	"github.com/santhureinoo_agnos_backend/repositories"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		// Log request and response to the database
		logEntry := models.LogEntry{
			Method:     c.Request.Method,
			Path:       c.Request.URL.Path,
			StatusCode: c.Writer.Status(),
			Duration:   duration,
		}
		go repositories.SaveLog(logEntry) // Save asynchronously
	}
}
