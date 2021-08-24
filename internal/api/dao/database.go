package dao

import (
	"github.com/GuillaumeDeconinck/todos-go/internal/api/configuration"
	"github.com/GuillaumeDeconinck/todos-go/pkg/models"
	"github.com/GuillaumeDeconinck/todos-go/pkg/tools"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(config *configuration.Configuration) {
	var err error
	db, err = gorm.Open(postgres.Open(config.GetDBUrl()), &gorm.Config{})
	if err != nil {
		tools.SugaredLogger.Fatal("failed to connect database")
	}

	db.AutoMigrate(&models.Todo{})
}
