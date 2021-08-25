package routeshandlers

import (
	"github.com/GuillaumeDeconinck/todos-go/internal/api/configuration"
	"github.com/gin-gonic/gin"
)

func SetupRouter(c configuration.Configuration) *gin.Engine {
	// Set mode
	gin.SetMode(gin.DebugMode)

	// New router
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Setup routes
	AddPingRoutesHandlers(r)
	AddTodosRoutesHandlers(r)

	return r
}
