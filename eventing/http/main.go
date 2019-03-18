package main

import (
	"github.com/gin-gonic/gin"
	"github.com/MarkLux/GOLD/eventing/http/trigger"
)


func main() {
	r := gin.Default()
	r.POST("/service/response", trigger.HandleInvoke)
	r.Run(":8080")
}
