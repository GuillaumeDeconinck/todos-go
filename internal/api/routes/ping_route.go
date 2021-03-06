package routes

import (
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func AddPingRoutesHandlers(r *gin.Engine) {
	r.GET("/ping", ping)
}
