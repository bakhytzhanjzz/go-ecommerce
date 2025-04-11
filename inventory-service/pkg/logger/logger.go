package logger

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{})
	return log
}

func GinLogger(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		latency := time.Since(start)

		log.WithFields(logrus.Fields{
			"status":  c.Writer.Status(),
			"method":  c.Request.Method,
			"path":    path,
			"ip":      c.ClientIP(),
			"latency": latency,
		}).Info("request")
	}
}
