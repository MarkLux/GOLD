package main

import (
	"github.com/gin-gonic/gin"
	"github.com/MarkLux/GOLD/eventing/http/trigger"
	"net/http"
)


func main() {
	r := gin.Default()
	r.Use(Cors())
	r.POST("/restful/response", trigger.HandleInvoke)
	r.Run(":8080")
}

// handle cors middleware
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}