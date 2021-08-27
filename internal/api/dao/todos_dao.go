package dao

import (
	daoerror "github.com/GuillaumeDeconinck/todos-go/internal/api/dao/daoError"
	"github.com/GuillaumeDeconinck/todos-go/pkg/models"
	"gorm.io/gorm"
)

func ListTodos(ownerUuid *string) ([]models.Todo, error) {
	var todos []models.Todo

	var result *gorm.DB
	if ownerUuid != nil && *ownerUuid != "" {
		result = db.Where("owner_uuid = ?", *ownerUuid).Find(&todos)
	} else {
		result = db.Find(&todos)
	}

	if result.Error != nil {
		return nil, daoerror.ConvertToDaoError(result.Error)
	}

	return todos, result.Error
}

func GetTodo(uuidToGet *string) (*models.Todo, error) {
	var todos []models.Todo
	result := db.Where("uuid = ?", *uuidToGet).Find(&todos)

	if len(todos) == 0 {
		return nil, daoerror.New(daoerror.NOT_FOUND_CODE)
	}

	if result.Error != nil {
		return nil, daoerror.ConvertToDaoError(result.Error)
	}

	return &todos[0], nil
}

func CreateTodo(todo *models.Todo) error {
	result := db.Create(&todo)

	if result.Error != nil {
		return daoerror.ConvertToDaoError(result.Error)
	}

	return nil
}

func UpdateTodo(todo *models.Todo) error {
	result := db.Model(todo).Where("uuid = ?", todo.Uuid).Updates(&todo)

	if result.Error != nil {
		return daoerror.ConvertToDaoError(result.Error)
	}

	if result.RowsAffected == 0 {
		return daoerror.New(daoerror.NOT_FOUND_CODE)
	}

	return nil
}

func DeleteTodo(uuidToDelete *string) error {
	result := db.Where("uuid = ?", *uuidToDelete).Delete(&models.Todo{})

	if result.Error != nil {
		return daoerror.ConvertToDaoError(result.Error)
	}

	return nil
}
