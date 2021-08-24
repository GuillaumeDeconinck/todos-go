package dao

import (
	"github.com/GuillaumeDeconinck/todos-go/pkg/models"
	"github.com/GuillaumeDeconinck/todos-go/pkg/tools"
)

func ListTodos(owner_uuid *string) ([]models.Todo, error) {
	var todos []models.Todo
	result := db.Find(&todos)

	// var todo models.Todo
	// result := db.First(&todo)

	// if result.Error != nil {
	// 	log.Fatalln("Error while retrieving the todos")
	// }

	// var todos []models.Todo
	// todos = [todo]

	return todos, result.Error
}

func GetTodo() {

}

func CreateTodo(todo *models.Todo) (*string, error) {
	tools.SugaredLogger.Infof("Create todo")
	result := db.Create(&todo)

	return todo.Uuid, result.Error
}

func UpdateTodo() {

}

func DeleteTodo(uuidToDelete *string) error {
	result := db.Where("uuid = ?", *uuidToDelete).Delete(&models.Todo{})
	return result.Error
}
