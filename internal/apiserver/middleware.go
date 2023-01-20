package apiserver

import (
	"github.com/gin-gonic/gin"
)

func headerCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("X-PING")
		if header != "ping"{
			return
		}
		c.Writer.Header().Set("X-PONG", "pong")
		c.Next()
	}
}