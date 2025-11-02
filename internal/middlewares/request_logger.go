package middlewares

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func coloredBackground(backgroundColor int, text string) string {
	return fmt.Sprintf("\x1b[37;%vm %v \x1b[0m", backgroundColor, text)
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestTime := time.Now()
		c.Next()
		log.Printf("[CUSTOM-LOGGER] | %s | %s | %s | %s  --  %s", coloredBackground(45, c.Request.Method), time.Since(requestTime),
			c.Request.URL.Host, coloredBackground(42, fmt.Sprint(c.Writer.Status())), c.Request.URL.Path)
	}
}
