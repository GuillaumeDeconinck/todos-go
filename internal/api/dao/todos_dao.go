package dao

import (
	"github.com/GuillaumeDeconinck/todos-go/pkg/models"
)

func ListTodos(owner_uuid *string) ([]models.Todo, error) {
	var todos []models.Todo
	result := db.Find(&todos)

	return todos, result.Error
}

func GetTodo(uuidToGet *string) (models.Todo, error) {
	var todo models.Todo
	result := db.Where("uuid = ?", *uuidToGet).First(&todo)

	return todo, result.Error
}

func CreateTodo(todo *models.Todo) error {
	result := db.Create(&todo)

	return result.Error
}

func UpdateTodo(todo *models.Todo) error {
	result := db.Where("uuid = ?", todo.Uuid).Save(&todo)

	return result.Error
}

func DeleteTodo(uuidToDelete *string) error {
	result := db.Where("uuid = ?", *uuidToDelete).Delete(&models.Todo{})
	return result.Error
}
