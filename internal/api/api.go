package api

import (
	"github.com/GuillaumeDeconinck/todos-go/internal/api/configuration"
	"github.com/GuillaumeDeconinck/todos-go/internal/api/dao"
	"github.com/GuillaumeDeconinck/todos-go/internal/api/handlers"
	"github.com/GuillaumeDeconinck/todos-go/pkg/tools"
	"github.com/gin-gonic/gin"
)

func SetupApi() *gin.Engine {
	var configuration, err = configuration.LoadConfig(".")
	if err != nil {
		tools.SugaredLogger.Warnf("Couldn't get the configuration: %s\n", err)
	}

	dao.InitDB(&configuration)

	r := handlers.SetupRouter(configuration)

	return r
}