package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"math"
	"time"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		log.Info().Fields(map[string]interface{}{
			"method":     c.Request.Method,
			"path":       path,
			"statusCode": statusCode,
			"cost":       latency,
			"clientIP":   clientIP,
			"referer":    referer,
			"dataLength": dataLength,
			"userAgent":  clientUserAgent,
		}).Send()
	}
}
