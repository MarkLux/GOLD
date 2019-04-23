package main

import (
	"github.com/MarkLux/GOLD/api/restful/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(Cors())

	// user controller
	userController := controller.NewUserController()

	r.POST("/user/register", userController.RegisterUser)
	r.POST("/user/login", userController.LoginUser)
	r.GET("/user/current", userController.GetLoginUser)

	// function controller
	functionController := controller.NewFunctionServiceController()
	r.POST("/function/service", functionController.CreateFunctionService)

	r.Run(":8090")
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